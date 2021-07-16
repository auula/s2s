/*
Copyright © 2021 Jarvib Ding

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
	"errors"
	"fmt"
	"os"
	"os/exec"
	"runtime"
	"strings"

	"github.com/apcera/termtables"
	"github.com/c-bata/go-prompt"
	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/app"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/emoji"

	"github.com/spf13/cobra"
)

var (
	commandSymbol  = "😃:s2s>"
	ConnectionInfo *db.Info
)

type Args struct {
	args []string
	sts  *core.Structure
}

type (
	Command func(args *Args) error
)

var (
	Use = func(args *Args) error {
		if len(args.args) == 1 {
			return errors.New("You do not choose which database, 👉 `use database` ")
		}
		if len(strings.TrimSpace(args.args[1])) == 0 {
			return errors.New("You do not choose which database, 👉 `use database` ")
		}
		args.sts.SetSchema(args.args[1])
		return nil
	}
	Tables = func(args *Args) error {
		t := termtables.CreateTable()
		t.AddHeaders("*", "Tables")
		ts, err := args.sts.Tables()
		if err != nil {
			return err
		}
		for i, v := range ts {
			t.AddRow(i+1, v)
		}
		fmt.Println(t.Render())
		return nil
	}
	Database = func(args *Args) error {
		t := termtables.CreateTable()
		t.AddHeaders("*", "Database")
		ds, err := args.sts.DataBases()
		if err != nil {
			return err
		}
		for i, v := range ds {
			t.AddRow(i+1, v)
		}
		fmt.Println(t.Render())
		return nil
	}
	Generate = func(args *Args) error {
		if len(args.args) == 1 {
			return errors.New("You do not choose which table, 👉 `gen table` ")
		}
		if len(strings.TrimSpace(args.args[1])) == 0 {
			return errors.New("You do not choose which table, 👉 `gen table` ")
		}
		args.sts.Parse(os.Stdout, args.args[1])
		return nil
	}
	Info = func(args *Args) error {
		t := termtables.CreateTable()
		t.AddRow("HOST", ConnectionInfo.HostIPAndPort)
		t.AddRow("USER", ConnectionInfo.UserName)
		t.AddRow("TYPE", ConnectionInfo.Type)
		fmt.Println(t.Render())
		return nil
	}
	shell = map[string]Command{
		"use":       Use,
		"info":      Info,
		"tables":    Tables,
		"databases": Database,
		"generate":  Generate,
	}
)

func ParseInput(cmd string, args *Args) {
	switch cmd {
	case "tables":
		if err := shell[cmd](args); err != nil {
			emoji.Error(err.Error())
		}
	case "databases":
		if err := shell[cmd](args); err != nil {
			emoji.Error(err.Error())
		}
	case "use":
		if err := shell[cmd](args); err != nil {
			emoji.Error(err.Error())
		} else {
			emoji.Info(fmt.Sprintf("Selected as database 👉 `%s`！", args.args[1]))
		}
	case "gen":
		if err := shell["generate"](args); err != nil {
			emoji.Error(err.Error())
		}
	case "info":
		if err := shell["info"](args); err != nil {
			emoji.Error(err.Error())
		}
	case "clear":
		clearFunc()
	case "EXIT":
		fallthrough
	case "exit":
		emoji.Info("Bye👋 :)")
		os.Exit(0)
	default:
		emoji.Error("Currently do not support commands you entered！！！")
	}
}

func shellPrompt(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "databases", Description: "Displays a list of all the database names."},
		{Text: "tables", Description: "Displays the current database of all table names."},
		{Text: "use", Description: "Specify the database name used."},
		{Text: "gen", Description: "Generate the structure corresponding to the data table."},
		{Text: "info", Description: "Current database connection information."},
		{Text: "exit", Description: "Enter `exit` to exit the program."},
		{Text: "clear", Description: "Content on the clean-up screen."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}

var clear map[string]func() //create a map for storing clear funcs

func init() {
	clear = make(map[string]func()) //Initialize it
	clear["linux"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["darwin"] = func() {
		cmd := exec.Command("clear") //Linux example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls") //Windows example, its tested
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

func clearFunc() {
	value, ok := clear[runtime.GOOS] //runtime.GOOS -> linux, windows, darwin etc.
	if ok {                          //if we defined a clear func for that platform:
		value() //we execute it
	} else { //unsupported platform
		panic("Your platform is unsupported! I can't clear terminal screen :(")
	}
}

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s2s",
	Short: "s2s is a command line database tool",
	Long:  app.Info.Description,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	ConnectionInfo = &db.Info{
		HostIPAndPort: os.Getenv("s2s_host"), // 数据库IP
		UserName:      os.Getenv("s2s_user"), // 数据库用户名
		Password:      os.Getenv("s2s_pwd"),  // 数据库密码
		Type:          db.MySQL,              // 数据库类型 PostgreSQL Oracle
		Charset:       os.Getenv("s2s_charset"),
	}
}
