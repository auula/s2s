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
	"os"

	"github.com/c-bata/go-prompt"
	"github.com/higker/s2s/core/app"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/emoji"
	"github.com/higker/s2s/core/lang/java"
	"github.com/spf13/cobra"
)

var (
	lang          string
	commandSymbol = "😃:shell>"
)

// consoleCmd represents the console command
var consoleCmd = &cobra.Command{
	Use:   "console",
	Short: "Console interaction",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(app.Info.Banner)
		fmt.Println()
		emoji.Success("You have entered the command line mode!")
		// for {
		// 	emoji.Info("Press the 'tab' key to get a prompt ")
		// 	t := prompt.Input(commandSymbol, completer)
		// 	commands.ParseInput(t, args)
		// }

		structure := java.New()

		if err := structure.OpenDB(
			&db.Info{
				HostIPAndPort: os.Getenv("HostIPAndPort"), // 数据库IP
				UserName:      "root",                     // 数据库用户名
				Password:      os.Getenv(""),              // 数据库密码
				Type:          db.MySQL,                   // 数据库类型 PostgreSQL Oracle
				Charset:       "utf8",
			},
		); err != nil {
			emoji.Error("Database Establishment Connection Fails, please detect configuration information!ment Connection Fails, please detect configuration information!ment Connection Fails, please detect configuration information!ment Connection Fails, please detect configuration information!")
			return
		}

		defer structure.Close()

		// 结果输出到标准输出   "数据库名"   "表名"
		structure.Parse(os.Stdout, "emp_db", "user_info")
	},
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "tables", Description: "Show tables infomation command."},
		{Text: "databases", Description: "Show database infomation command."},
		{Text: "use", Description: "Database of select using."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

func init() {
	rootCmd.AddCommand(consoleCmd)

}
