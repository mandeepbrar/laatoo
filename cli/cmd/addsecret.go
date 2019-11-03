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
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"laatoo/sdk/utils"
	"os"

	"github.com/spf13/cobra"
)

// addsecretCmd represents the addsecret command
var addsecretCmd = &cobra.Command{
	Use:   "addsecret",
	Short: "Add secret key to a keys base",
	Long: `Add secret key to a keys base for use by laatoo server. 

		Usage: laatoo addsecret [flags] keyname
		Flags:
		 -d, --directory: Directory where keys base is located
	`,
	Args: func(cmd *cobra.Command, args []string) error {
		if len(args) < 1 {
			return errors.New("Key name not specified")
		}
		return nil
	},
	Run: addSecretCmdFunc,
}

func init() {
	rootCmd.AddCommand(addsecretCmd)

	addsecretCmd.PersistentFlags().StringP("directory", "d", "", "base directory to use for storing secret keys")
	addsecretCmd.PersistentFlags().StringP("file", "f", "", "file containing value of key")
}

func addSecretCmdFunc(cmd *cobra.Command, args []string) {
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
	file, _ := cmd.Flags().GetString("file")
	addSecretKeyVal(directory, key, file)
}

func addSecretKeyVal(directory, key, file string) {
	if len(key) == 0 {
		exitWithMessage("Invalid key specified")
	}

	val := []byte{}
	if file != "" {
		byts, err := ioutil.ReadFile(file)
		if err != nil {
			exitWithError(err)
		} else {
			val = byts
		}
	} else {
		str := ""
		getInput("Enter value of the key: ", &str)
		if len(str) == 0 {
			exitWithMessage("No value provided for the key")
		}
		val = []byte(str)
	}

	storer := utils.NewDiskStorer(directory, 10*1024*1024)
	storer.PutObject(key, encryptVal(val))
	fmt.Println("Secret added ", key)
}

func encryptVal(val []byte) []byte {
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

	nonce := make([]byte, gcm.NonceSize())
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		exitWithError(err)
	}

	return gcm.Seal(nonce, nonce, val, nil)

}
