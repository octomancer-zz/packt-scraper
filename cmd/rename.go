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
	"github.com/spf13/viper"
	"gitlab.com/octomancer/packt-scraper/books"
)

// renameCmd represents the rename command
var renameCmd = &cobra.Command{
	Use:   "rename",
	Short: "Report on filenames and optionally rename",
	Long: `
Searches for a variety of filenames for each book file and renames them to the canonical (ie. Cloudfront) name
`,
	Run: func(cmd *cobra.Command, args []string) {
		glog.V(0).Infof("rename called")
		bl := books.EntitlementBooks{}
		bl.Rename()
	},
}

func init() {
	rootCmd.AddCommand(renameCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// renameCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// renameCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	renameCmd.Flags().BoolP("do-rename", "r", false, "Actually rename the files. Without this just print what we *would* do.")
	viper.BindPFlag("do-rename", renameCmd.Flags().Lookup("do-rename"))
}
