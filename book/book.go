package book

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"

	"github.com/golang/glog"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/client"
	"gitlab.com/octomancer/packt-scraper/locations"
)

// Checks map of bookfiles to see if we already have this one
func (b *EntitlementBook) BookFileExists(extension string) (bool, string) {
	bf := locations.GetBookFiles()

	key := b.ProductId + extension
	if name, ok := (*bf)[key]; ok {
		return true, name
	}

	return false, ""
}

// Print book title and flip boolean
func (b *EntitlementBook) PrintTitle(idx int) {
	if b.TitlePrinted {
		return
	}
	fmt.Printf("%03d : %s (%5s)\n", Brown(idx).Bold(), Brown(b.ProductName).Bold(), b.ProductId)
	b.TitlePrinted = true
}

// Download the link for one extension for a book
func (b *EntitlementBook) Download(extension string) (bool, error) {
	glog.V(0).Infof("Processing %s", extension)

	url := fmt.Sprintf("%sproducts-v1/products/%s/files/%s", locations.ServicesURL, b.ProductId, extension)

	glog.V(1).Infof("Using URL %s", url)

	// Get HTTP REST client
	api := client.GetClient()
	api.Login()

	ext := extension
	if ext == "code" {
		ext = "zip"
	}
	re := regexp.MustCompile(fmt.Sprintf(".*cloudfront.net/.*/(%s.*%s)", b.ProductId, ext))

	retry := true
	retries := 5
	for retry && retries > 0 {
		// logic is much easier if we assume success as there is only one
		// condition that triggers a retry
		retry = false
		retries--

		glog.V(0).Infof("Fetching book file")

		rd := &ResponseData{}
		api.GetPageJSON(url, rd)

		url := rd.Data
		if url != "" {
			glog.V(1).Infof("Got book url %s\n", url)
		} else {
			dummyFile := fmt.Sprintf("%s%s-dummy.%s", viper.GetString("bookdir"), b.ProductId, ext)
			os.Create(dummyFile)
			return false, fmt.Errorf("%s", rd.Message)
		}
		m := re.FindAllStringSubmatch(url, -1)
		if m == nil {
			return false, fmt.Errorf("Couldn't parse book filename from %s\n", url)
		}
		bookFile := fmt.Sprintf("%s%s", viper.GetString("bookdir"), m[0][1])
		glog.V(1).Infof("Got book filename %s\n", bookFile)

		res, err := api.Do(url)

		f, err := os.Create(bookFile)
		if err != nil {
			return true, fmt.Errorf("Error opening %s for writing: %s", bookFile, err)
		}
		defer f.Close()
		w, err := io.Copy(f, res.Body)
		if err != nil {
			return true, fmt.Errorf("Error writing to %s: %s", bookFile, err)
		}
		glog.V(0).Infof("Wrote %d bytes to %s", w, bookFile)
		if w == 110 {
			glog.V(0).Infof("Bytes written is 110. Examining file.")
			f2, err := os.Open(bookFile)
			if err != nil {
				return true, fmt.Errorf("Error opening %s for reading: %s", bookFile, err)
			}
			defer f2.Close()
			re, _ := regexp.Compile("<Error><Code>AccessDenied</Code><Message>Access denied</Message></Error>")
			s := bufio.NewScanner(f2)
			for s.Scan() {
				if m := re.MatchString(s.Text()); m {
					retry = true
					if retries > 0 {
						glog.Errorf("Found Access denied. Trying again.")
					} else {
						glog.Errorf("Found Access denied. Giving up.")
						f2.Close()
						os.Remove(bookFile)
					}
				}
			}
		}
	}

	return true, nil
}
