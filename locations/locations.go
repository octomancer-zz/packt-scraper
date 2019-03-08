package locations

import (
	"fmt"
	"os"
	"regexp"

	"github.com/golang/glog"
	"github.com/spf13/viper"
)

var bookFiles *map[string]string

// Generate the filename for a page
func GetMyEbooksPageFilename(page int) string {
	return fmt.Sprintf("%smy-ebooks%02d.html", viper.GetString("htmldir"), page)
}

// Generate the filename for a page
func GetFlbsFilename(productId string) string {
	return fmt.Sprintf("%s%s.json", viper.GetString("infodir"), productId)
}

// Initialise bookFiles if necessary and return pointer
func GetBookFiles() *map[string]string {
	if bookFiles == nil {
		FillBookFiles()
	}
	return bookFiles
}

// Create the map of productId + extension -> filenames
func FillBookFiles() {
	bf := make(map[string]string)
	bookFiles = &bf

	dir, err := os.Open(viper.GetString("bookdir"))
	if err != nil {
		glog.Fatalf("Error reading book dir: %s", err)
	}

	re := regexp.MustCompile("^(\\d+)[\\._-].*(pdf|epub|mobi|zip)")
	files, err := dir.Readdir(0)
	for _, f := range files {
		if f.IsDir() {
			continue
		}
		m := re.FindAllStringSubmatch(f.Name(), -1)
		if m == nil {
			continue
		}
		extension := m[0][2]
		if extension == "zip" {
			extension = "code"
		}
		key := m[0][1] + extension
		(*bookFiles)[key] = f.Name()
	}
}
