/*
Copyright © 2021 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"fmt"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/higker/s2s/core/app"
	"github.com/higker/s2s/core/emoji"
	"github.com/higker/s2s/core/lang/rust"
	"github.com/spf13/cobra"
)

// rustCmd represents the rust command
var rustCmd = &cobra.Command{
	Use:   "rust",
	Short: "Generate Rust code",
	Long: `
  Produce Rust code mapping according to database table.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Info.Banner)
		fmt.Println()
		emoji.Success("You have entered the command line mode!")

		structure := rust.New()

		if err := structure.OpenDB(
			ConnectionInfo,
		); err != nil {
			emoji.Error("Failed to establish a connection to the database!")
		}

		defer structure.Close()

		emoji.Success("Press the 'tab' key to get a prompt！")
		emoji.Success("Enter `exit` to exit the program!")
		for {
			fmt.Println()
			t := prompt.Input(commandSymbol, shellPrompt)
			if len(strings.TrimSpace(t)) <= 0 {
				continue
			}

			ParseInput(strings.Split(t, " ")[0], &Args{
				sts:  structure,
				args: strings.Split(t, " "),
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(rustCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// rustCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// rustCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
