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
	"flag"

	"github.com/golang/glog"
	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var cfgFile string
var startIdx, endIdx, logLevel int
var fetch bool

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "packt-scraper",
	Short: "A utility to claim, download and organise packt free e-books",
	Long: `
Packt provide one free ebook per day at https://www.packtpub.com/packt/offers/free-learning. Each e-book has to be claimed on the day and it is added to your account. Download links are provided for a variety of formats eg. pdf, epub. Some ebooks also have an associated code file which will contain any runnable code referenced in the ebook itself.

This is a utility to manage your ebook empire.
`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		glog.Infof("%s", err)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// To keep glog happy
	rootCmd.PersistentFlags().AddGoFlagSet(flag.CommandLine)
	rootCmd.PersistentFlags().Lookup("logtostderr").Value.Set("true")
	rootCmd.PersistentFlags().Lookup("logtostderr").DefValue = "true"
	flag.CommandLine.Parse(nil) // glog detects if flags get parsed, so parse nil

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.
	rootCmd.PersistentFlags().StringVarP(&cfgFile, "config", "", "", "config file (default is $HOME/.packt-scraper.yaml)")
	rootCmd.PersistentFlags().IntVarP(&startIdx, "start", "s", 0, "Index at which to start in book list. Default 0. It is an error to specify an index greater than the number of books.")
	rootCmd.PersistentFlags().IntVarP(&endIdx, "end", "e", 0, "Index at which to start in book list. Default 0, which means process all books. If the end index is greater than the number of books it will be set to the number of books.")
	rootCmd.PersistentFlags().BoolVarP(&fetch, "fetch", "f", false, "Fetch flag. If there are missing packt files or book files, set this to true to download them.")
	viper.BindPFlag("start", rootCmd.PersistentFlags().Lookup("start"))
	viper.BindPFlag("end", rootCmd.PersistentFlags().Lookup("end"))
	viper.BindPFlag("fetch", rootCmd.PersistentFlags().Lookup("fetch"))

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			glog.Fatalf("%s", err)
		}

		// Search config in home directory with name ".packt-scraper" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".packt-scraper")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		glog.V(0).Infof("Using config file: %s", viper.ConfigFileUsed())
	}
}
