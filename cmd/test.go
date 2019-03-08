// Copyright © 2018 Richard Gration <richgration@gmail.com>
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
	"encoding/json"
	"fmt"

	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.com/octomancer/packt-scraper/books"
)

// testCmd represents the test command
var testCmd = &cobra.Command{
	Use:   "test",
	Short: "Tests experimental code",
	Long: `
Runs the test code in the Test function.
`,
	Run: func(cmd *cobra.Command, args []string) {
		if glog.V(2) {
			glog.Info("test called")
		}
		eb := books.EntitlementBooks{}
		// This will trigger parseMyEbooks
		// _ = eb.Search([]string{"notabooktitle"})

		b, err := json.MarshalIndent(eb, "", "\t")
		if err != nil {
			glog.Fatalf("Error marshalling books: %s\n", err)
		}
		fmt.Println(string(b))

		for idx, book := range eb.Books {
			fmt.Printf("idx %d bookId %s title %s\n", idx, book.ProductId, book.ProductName)
		}
	},
}

func init() {
	rootCmd.AddCommand(testCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// testCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// testCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
