package books

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
	"time"

	"github.com/golang/glog"
	. "github.com/logrusorgru/aurora"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/book"
	"gitlab.com/octomancer/packt-scraper/client"
	"gitlab.com/octomancer/packt-scraper/locations"
)

const pauseSecs time.Duration = 5 * time.Second
const pageSize int = 25

// Search book titles for search term(s)
func (eb *EntitlementBooks) Search(search []string) []*book.EntitlementBook {
	// Make sure we have parsed my-ebooks*.html
	eb.ParseMyEbooks()

	var searchResults []*book.EntitlementBook
	if flag := viper.GetBool("all"); flag {
		// Match titles that have all search terms
		// Array to hold regexes
		var re []*regexp.Regexp
		for _, s := range search {
			r, _ := regexp.Compile(s)
			re = append(re, r)
		}

		// Loop over books
		idx := 0
		for _, b := range eb.Books {
			matched := true
			for _, r := range re {
				if !r.MatchString(b.ProductName) {
					matched = false
					break
				}
			}
			if matched {
				searchResults = append(searchResults, b)
			}
			idx++
		}
		return searchResults
	}
	var re *regexp.Regexp

	if flag := viper.GetBool("regex"); flag {
		re = regexp.MustCompile(search[0])
	} else {
		re = regexp.MustCompile(strings.Join(search, "|"))
	}

	// Loop over books
	for _, b := range eb.Books {
		if re.MatchString(b.ProductName) {
			searchResults = append(searchResults, b)
		}
	}
	return searchResults
}

// Just show the list of books in the data structure we created from my-ebooks.html
func (eb *EntitlementBooks) Show() {
	// Make sure we have parsed my-ebooks*.html
	eb.ParseMyEbooks()

	// Pick up start and end from command line or set defaults
	startIdx, endIdx := viper.GetInt("start"), viper.GetInt("end")
	if endIdx == 0 || endIdx > eb.Count {
		endIdx = eb.Count - 1
	}
	glog.V(0).Infof("Using start_index = %d, end_index = %d", startIdx, endIdx)

	extensions := []string{"pdf", "epub", "mobi", "code"}
	for idx := startIdx; idx <= endIdx; idx++ {
		b := eb.Books[idx]
		b.PrintTitle(idx)
		fmt.Printf("%-14s: %s\n", Blue("Product Id").Bold(), b.ProductId)
		fmt.Printf("%-14s: %s\n", Blue("Name").Bold(), b.ProductName)
		if flag := viper.GetBool("show-filenames"); flag {
			fmt.Printf("%-20s\n", Blue("local filenames").Bold())
			for _, extension := range extensions {
				if b.LocalFilenames[extension] != "" {
					fmt.Printf("%13s : %s\n", Cyan(extension), b.LocalFilenames[extension])
				}
			}
		}
	}
}

// Iterate over our data structure and download ALL THE THINGS
func (eb *EntitlementBooks) Download() error {
	// Make sure we have parsed my-ebooks*.html
	eb.ParseMyEbooks()

	// Pick up start and end from command line or set defaults
	startIdx, endIdx := viper.GetInt("start"), viper.GetInt("end")
	if endIdx == 0 || endIdx > eb.Count {
		endIdx = eb.Count - 1
	}
	glog.V(0).Infof("Using start_index = %d, end_index = %d", startIdx, endIdx)

	fetchFlag := viper.GetBool("fetch")
	extensions := []string{"pdf", "epub", "mobi", "code"}
	for idx := startIdx; idx <= endIdx; idx++ {
		b := eb.Books[idx]

		if glog.V(1) {
			b.PrintTitle(idx)
		}

		for _, extension := range extensions {
			fe, fn := b.BookFileExists(extension)
			if fe {
				fn = viper.GetString("bookdir") + fn
				glog.V(1).Infof("Have file for %-4s: \"%s\"", extension, fn)
				continue
			}
			b.PrintTitle(idx)
			glog.V(0).Infof("No file for \"%s\"", extension)
			if fetchFlag {
				glog.V(0).Infof("Fetch flag is true, trying to fetch")
			} else {
				glog.V(0).Infof("Fetch flag is false, not fetching")
				continue
			}

			fetched, err := b.Download(extension)
			if err != nil {
				glog.Errorf("Error doing book: %s", err)
			}
			if fetched {
				glog.V(1).Infof("Pausing for %v", pauseSecs)
				time.Sleep(pauseSecs)
			}
		}
	}
	return nil
}

// Fetch and save all my-ebooks pages
func (eb *EntitlementBooks) FetchMyEbooks() {
	// Get first page to set count properly
	eb.fetchMyEbooksPage(0)

	// Build up list of books in eb.Books
	pageNo := 1
	for len(eb.Books) < eb.Count {
		eb.fetchMyEbooksPage(pageNo)
		pageNo++
	}

	// Get the first my-ebooks filename
	fn := locations.GetMyEbooksPageFilename(0)

	f, err := os.Create(fn)
	if err != nil {
		glog.Fatalf("Error opening \"%s\": %s\n", fn, err)
	}
	defer f.Close()
	b, err := json.Marshal(eb)
	if err != nil {
		glog.Fatalf("Error marshalling books: %s\n", err)
	}
	f.Write(b)
}

func (eb *EntitlementBooks) fetchMyEbooksPage(pageNo int) {
	// Get HTTP REST client
	api := client.GetClient()
	api.Login()

	// Build URL
	url := fmt.Sprintf("%sentitlements-v1/users/me/products?sort=createdAt:DESC&limit=%d", locations.ServicesURL, pageSize)
	if pageNo > 0 {
		url = fmt.Sprintf("%s&offset=%d", url, pageNo*pageSize)
	}

	// Get my-ebooks page
	ebs := &EntitlementBooks{}
	api.GetPageJSON(url, ebs)

	eb.Count = ebs.Count
	eb.Books = append(eb.Books, ebs.Books...)
}

func (eb *EntitlementBooks) Diagnostics() {
	// Make sure we have parsed my-ebooks*.html
	eb.ParseMyEbooks()

	// Compile a map of book filenames
	bf := make(map[string]bool)
	for _, b := range eb.Books {
		for _, fn := range b.LocalFilenames {
			bf[fn] = true
		}
	}

	fmt.Printf("Finding aliens in bookdir\n")
	f, _ := os.Open(viper.GetString("bookdir"))
	fi, _ := f.Readdir(0)
	f.Close()
	for i, v := range fi {
		if v.IsDir() {
			continue
		}
		if bf[v.Name()] {
			continue
		}
		fmt.Printf("%d - %s\n", i, v.Name())
	}

	// Pick up start and end from command line or set defaults
	startIdx, endIdx := viper.GetInt("start"), viper.GetInt("end")
	if endIdx == 0 || endIdx > eb.Count {
		endIdx = eb.Count - 1
	}
	glog.V(0).Infof("Using start_index = %d, end_index = %d", startIdx, endIdx)
	fmt.Printf("Looking for missing filenames\n")

	//extensions := []string{"pdf","epub","mobi","code"}
	for idx := startIdx; idx <= endIdx; idx++ {
		b := eb.Books[idx]

		if len(b.LocalFilenames) == 0 {
			b.PrintTitle(idx)
			fmt.Printf("No localFilenames for %s\n", b.ProductName)
		}
	}
}

// sorts any array of *EntitlementBook by productId
func (eb *EntitlementBooks) SortBookArray() {
	// Sorts by book id
	sort.Slice(eb.Books, func(i, j int) bool {
		return eb.Books[i].ProductId < eb.Books[j].ProductId
	})
}

func (eb *EntitlementBooks) FindBookWithProductId(pid string) *book.EntitlementBook {
	eb.ParseMyEbooks()

	for _, b := range eb.Books {
		if b.ProductId == pid {
			return b
		}
	}
	return nil
}

// Read the downloaded my-ebooks page and unmarshall into global variable
// Also scan book directory to see what we have already downloaded
func (eb *EntitlementBooks) ParseMyEbooks() {
	if eb.Count > 0 {
		return
	}
	fn := locations.GetMyEbooksPageFilename(0)
	f, err := os.Open(fn)
	if err != nil {
		glog.Fatalf("Error opening \"%s\" for reading: %s", fn, err)
	}
	defer f.Close()

	err = json.NewDecoder(f).Decode(&eb)
	if err != nil {
		glog.Fatalf("Error decoding my-ebooks file into struct: %s", err)
	}

	eb.SortBookArray()

	for _, b := range eb.Books {
		if b.LocalFilenames == nil {
			b.LocalFilenames = make(map[string]string)
		}
		for _, ext := range []string{"pdf", "epub", "mobi", "code"} {
			exists, fn := b.BookFileExists(ext)
			if exists {
				b.LocalFilenames[ext] = fn
			}
		}
	}
	glog.V(0).Infof("Found %d books in my-ebooks file", eb.Count)
}

// Renames existing book files
func (eb *EntitlementBooks) Rename() {
	// Make sure we have parsed my-ebooks*.html
	eb.ParseMyEbooks()

	// Pick up start and end from command line or set defaults
	startIdx, endIdx := viper.GetInt("start"), viper.GetInt("end")
	if endIdx == 0 || endIdx > eb.Count {
		endIdx = eb.Count - 1
	}
	glog.V(0).Infof("Using start_index = %d, end_index = %d", startIdx, endIdx)

	extensions := []string{"pdf", "epub", "mobi", "code"}
	for idx := startIdx; idx <= endIdx; idx++ {
		b := eb.Books[idx]

		b.PrintTitle(idx)
		for _, extension := range extensions {
			fe, fn := b.BookFileExists(extension)
			if fe {
				fn = viper.GetString("bookdir") + fn
				fmt.Printf("%13s : %s\n", Cyan(extension), fn)
				lfn := viper.GetString("bookdir") + "newnameforbook"
				if fn != lfn {
					if flag := viper.GetBool("do-rename"); flag {
						os.Rename(fn, lfn)
						fmt.Printf("%20s : %s\n", Red("Renamed to"), Red(lfn).Bold())
					} else {
						fmt.Printf("%20s : %s\n", Red("Would rename to"), Red(lfn).Bold())
					}
				}
			} else {
				fmt.Printf("%13s : %s", Cyan(extension), "missing")
			}
		}
	}
}
