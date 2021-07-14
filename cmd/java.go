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
	"os"
	"strings"

	"github.com/c-bata/go-prompt"
	"github.com/higker/s2s/core/app"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/emoji"
	"github.com/higker/s2s/core/lang/java"
	"github.com/spf13/cobra"
)

// javaCmd represents the java command
var javaCmd = &cobra.Command{
	Use:   "java",
	Short: "Generate Java code",
	Long: `
  Produce Java code mapping according to database table.
	`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Info.Banner)
		fmt.Println()
		emoji.Success("You have entered the command line mode!")

		structure := java.New()

		if err := structure.OpenDB(
			&db.Info{
				HostIPAndPort: os.Getenv("s2s_host"), // 数据库IP
				UserName:      os.Getenv("s2s_user"), // 数据库用户名
				Password:      os.Getenv("s2s_pwd"),  // 数据库密码
				Type:          db.MySQL,              // 数据库类型 PostgreSQL Oracle
				Charset:       os.Getenv("s2s_charset"),
			},
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
				args: strings.Split(t, " ")[1:],
			})
		}
	},
}

func init() {
	rootCmd.AddCommand(javaCmd)
}
