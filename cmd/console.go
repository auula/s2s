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
	"github.com/higker/s2s/core"
	"log"

	"github.com/higker/s2s/core/lang/golang"
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

		structure := golang.New()

		if err := structure.OpenDB(
			&core.DBInfo{
				HostIPAndPort: "45.76.202.255:3306",
				UserName:      "emp_db",
				Password:      "TsTkHXDK4xPFtCph",
				Type:          core.MySQL,
				Charset:       "utf8",
			},
		); err != nil {
			log.Println(err)
		}

		defer structure.Close()

		structure.Parse("info", "info")
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
