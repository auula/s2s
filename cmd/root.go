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
	"fmt"
	"os"
	"os/exec"
	"runtime"

	"github.com/apcera/termtables"
	"github.com/c-bata/go-prompt"
	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/app"
	"github.com/higker/s2s/core/emoji"

	"github.com/spf13/cobra"
)

type (
	Command func(args []string, sts *core.Structure) error
)

var (
	Use = func(args []string, sts *core.Structure) error {
		return nil
	}
	Tables = func(args []string, sts *core.Structure) error {
		fmt.Println("tables")
		return nil
	}
	Database = func(args []string, sts *core.Structure) error {
		t := termtables.CreateTable()
		t.AddHeaders("*", "Database")
		ds, err := sts.DataBases()
		if err != nil {
			return err
		}
		for i, v := range ds {
			t.AddRow(i+1, v)
		}
		fmt.Println(t.Render())
		return nil
	}
	Generate = func(args []string, sts *core.Structure) error {

		return nil
	}
	shell = map[string]Command{
		"use":       Use,
		"tables":    Tables,
		"databases": Database,
		"generate":  Generate,
	}
)

func ParseInput(cmd string, agrs []string, sts *core.Structure) {
	switch cmd {
	case "tables":
		shell[cmd](agrs, sts)
	case "databases":
		shell[cmd](agrs, sts)
	case "use":
		shell[cmd](agrs, sts)
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
		{Text: "tables", Description: "Displays the current database of all table names."},
		{Text: "databases", Description: "Displays a list of all the database names."},
		{Text: "use", Description: "Specify the database name used."},
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

// var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "s2s",
	Short: "s2s is a command line database tool",
	Long:  app.Info.Description,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	// Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	cobra.CheckErr(rootCmd.Execute())
}

func init() {
	//cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	//rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.s2s.yaml)")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	//rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
// func initConfig() {
// 	if cfgFile != "" {
// 		// Use config file from the flag.
// 		viper.SetConfigFile(cfgFile)
// 	} else {
// 		// Find home directory.
// 		home, err := homedir.Dir()
// 		cobra.CheckErr(err)

// 		// Search config in home directory with name ".s2s" (without extension).
// 		viper.AddConfigPath(home)
// 		viper.SetConfigName(".s2s")
// 	}

// 	viper.AutomaticEnv() // read in environment variables that match

// 	// If a config file is found, read it in.
// 	if err := viper.ReadInConfig(); err == nil {
// 		fmt.Fprintln(os.Stderr, "Using config file:", viper.ConfigFileUsed())
// 	}
// }
