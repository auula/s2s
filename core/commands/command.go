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

package commands

import (
	"fmt"

	"github.com/c-bata/go-prompt"
)

type (
	Command func(args []string) error
)

var (
	Use = func(args []string) error {
		return nil
	}
	Tables = func(args []string) error {
		fmt.Println("tables")
		return nil
	}
	Database = func(args []string) error {
		return nil
	}
	Generate = func(args []string) error {
		return nil
	}
	Execute = map[string]Command{
		"use":       Use,
		"tables":    Tables,
		"databases": Database,
		"generate":  Generate,
	}
)

func ParseInput(cmd string, agrs []string) error {
	switch cmd {
	case "tables":
		Execute[cmd](agrs)
	case "databases":
		Execute[cmd](agrs)
	case "use":
		Execute[cmd](agrs)
	default:
		return nil
	}
	return nil
}

func completer(d prompt.Document) []prompt.Suggest {
	s := []prompt.Suggest{
		{Text: "tables", Description: "Show tables infomation command."},
		{Text: "databases", Description: "Show database infomation command."},
		{Text: "use", Description: "Database of select using."},
	}
	return prompt.FilterHasPrefix(s, d.GetWordBeforeCursor(), true)
}
