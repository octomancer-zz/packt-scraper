// Copyright © 2018 Richard Gration <richard.gration@skybettingandgaming.com>
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License
// along with this program. If not, see <http://www.gnu.org/licenses/>.

package cmd

import (
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/books"
)

var all, regex bool

// searchCmd represents the search command
var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search book titles for search terms",
	Long: `
Search book titles for search terms

Example:

packt-scraper search -r '(?i)python'

This searches case-insensitively for the string 'python'.
`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.V(0).Infof("search called")
		bl := &books.EntitlementBooks{}
		searchResults := bl.Search(args)

		for _, book := range searchResults {
			fmt.Printf("%s: %s\n", book.ProductId, book.ProductName)
		}
	},
}

func init() {
	rootCmd.AddCommand(searchCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// searchCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// searchCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	searchCmd.PersistentFlags().BoolVarP(&all, "all", "a", false, "Match titles with all search terms, rather than any search terms (default).")
	searchCmd.PersistentFlags().BoolVarP(&regex, "regex", "r", false, "Treat argument as a regex to compile")
	viper.BindPFlag("all", searchCmd.PersistentFlags().Lookup("all"))
	viper.BindPFlag("regex", searchCmd.PersistentFlags().Lookup("regex"))
}
