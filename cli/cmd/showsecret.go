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
	"crypto/aes"
	"crypto/cipher"
	"fmt"
	"laatoo/sdk/utils"
	"os"

	"github.com/spf13/cobra"
)

// showsecretCmd represents the showsecret command
var showsecretCmd = &cobra.Command{
	Use:   "showsecret",
	Short: "Shows the value of secret in a keystore",
	Long: `Shows secret key in a keys base for use by laatoo server. 
	
	Usage: laatoo showsecret [flags] keyname
	Flags:
	 -d, --directory: Directory where keys base is located`,
	Run: showSecretCmdFunc,
}

func init() {
	rootCmd.AddCommand(showsecretCmd)
	showsecretCmd.PersistentFlags().StringP("directory", "d", "", "base directory to use for secret keys")

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// showsecretCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// showsecretCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

func showSecretCmdFunc(cmd *cobra.Command, args []string) {
	directory, _ := cmd.Flags().GetString("directory")
	if directory == "" {
		exitWithMessage("Directory must be spcified for adding a secret")
	} else {
		fi, err := os.Stat(directory)
		if err != nil || !fi.IsDir() {
			exitWithMessage("Invalid directory provided")
		}
	}
	key := args[0]
	showSecretKeyVal(directory, key)
}

func showSecretKeyVal(directory, key string) {
	if len(key) == 0 {
		exitWithMessage("Invalid key specified")
	}
	storer := utils.NewDiskStorer(directory, 10*1024*1024)
	val, err := storer.GetObject(key)
	if err != nil {
		exitWithError(err)
	}
	str := decryptVal(val)
	fmt.Println("Value of key ", key, " = ", str)
}

func decryptVal(val []byte) string {
	secret := ""

	getInput("Enter key to be used for encryption of values: ", &secret)

	fmt.Println("lenth of secret", len(secret))
	if len(secret) != 16 && len(secret) != 32 {
		exitWithMessage("Invalid secret key specified. Secret should be 8, 16 or 32 characters")
	}

	c, err := aes.NewCipher([]byte(secret))
	// if there are any errors, handle them
	if err != nil {
		exitWithError(err)
	}

	gcm, err := cipher.NewGCM(c)
	// if any error generating new GCM
	// handle them
	if err != nil {
		exitWithError(err)
	}

	nonceSize := gcm.NonceSize()

	if len(val) < nonceSize {
		exitWithMessage("Error decrypting key, malformed ciphertext")
	}

	strToRet, err := gcm.Open(nil, val[:nonceSize], val[nonceSize:], nil)
	if err != nil {
		exitWithError(err)
	}
	return string(strToRet)
}
