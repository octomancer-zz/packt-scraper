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

// diagnosticsCmd represents the diagnostics command
var diagnosticsCmd = &cobra.Command{
	Use:   "diagnostics",
	Short: "Report on unknown book files and undownloaded books",
	Long: `
Shows a list of files in bookdir that aren't in the list of claimed books and shows titles that have no downloaded files.
`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.V(0).Infof("diagnostics called")
		bl := books.EntitlementBooks{}
		bl.Diagnostics()
	},
}

func init() {
	rootCmd.AddCommand(diagnosticsCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// diagnosticsCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// diagnosticsCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
