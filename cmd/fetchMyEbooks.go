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
	"github.com/golang/glog"
	"github.com/spf13/cobra"
	"gitlab.com/octomancer/packt-scraper/books"
)

// fetchMyEbooksCmd represents the fetch command
var fetchMyEbooksCmd = &cobra.Command{
	Use:   "fetchMyEbooks",
	Short: "Download the my-ebooks JSON and write to disk",
	Long: `
Download the my-ebooks JSON and writes it into htmldir.
`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.V(0).Infof("fetchMyEbooks called")
		bl := books.EntitlementBooks{}
		bl.FetchMyEbooks()
	},
}

func init() {
	rootCmd.AddCommand(fetchMyEbooksCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// fetchMyEbooksCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// fetchMyEbooksCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
