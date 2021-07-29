/*
Copyright © 2021 Jarvib Ding <ding@ibyte.me>
Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:
The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.
THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
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
