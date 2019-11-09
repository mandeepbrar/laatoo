// Copyright © 2019 NAME HERE <EMAIL ADDRESS>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"fmt"
	"laatoo/cli/commands"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// appCreateCmd represents the appCreate command
var appCreateCmd = &cobra.Command{
	Use:   "create",
	Short: "Creates laatoo application skeleton inside a folder",
	Long: `Creates basic application skeleton inside a folder

	Usage: laatoo app create appname
	`,
	Run: func(cmd *cobra.Command, args []string) {
		apptempDir := viper.GetString("apptemplate")
		if apptempDir != "" {
			if !filepath.IsAbs(apptempDir) {
				apptempDir = filepath.Join(configHome, apptempDir)
			}
			if verbose {
				fmt.Println("Using app template directory ", apptempDir)
			}
			destdir, _ := cmd.Flags().GetString("destination")

			var err error
			var appname string
			if len(args) == 0 {
				appname = filepath.Base(destdir)
			} else {
				appname = args[0]
			}
			if verbose {
				fmt.Println("Creating app ", appname)
			}
			err = commands.Createapp(appname, destdir, apptempDir)
			if err != nil {
				exitWithError(err)
			}
		}
	},
}

func init() {
	appCmd.AddCommand(appCreateCmd)
	destdir, err := os.Getwd()
	if err != nil {
		exitWithError(err)
	}
	appCreateCmd.Flags().StringP("destination", "d", destdir, "Help message for toggle")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// appCreateCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// appCreateCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
