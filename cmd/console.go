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
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/lang/rust"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	mode          string
	comoandSymbol = "😃:shell>"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Console interaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {

		structure := rust.New()

		if err := structure.OpenDB(
			&db.Info{
				HostIPAndPort: os.Getenv("HostIPAndPort"), // 数据库IP
				UserName:      "emp_db",                   // 数据库用户名
				Password:      os.Getenv("Password"),      // 数据库密码
				Type:          db.MySQL,                   // 数据库类型 PostgreSQL Oracle
				Charset:       "utf8",
			},
		); err != nil {
			log.Println(err)
		}

		defer structure.Close()

		// 结果输出到标准输出   "数据库名"   "表名"
		structure.Parse(os.Stdout, "emp_db", "user_info")
	},
}

func init() {
	rootCmd.AddCommand(consoleCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// consoleCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// consoleCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
