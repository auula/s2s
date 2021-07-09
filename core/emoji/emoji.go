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
package emoji

import (
	"fmt"

	"github.com/fatih/color"
)

const (
	Mark       = "⚠️"
	Programmer = "🤖‍"
	Cry        = "😭"
	Cheer      = "🥳"
)

func Info(msg string) {
	fmt.Println(fmt.Sprintf("%s: %s .", Programmer, color.YellowString(msg)))
	fmt.Println()
}

func Error(msg string) {
	fmt.Println(fmt.Sprintf("%s: %s .", Cry, color.RedString(msg)))
	fmt.Println()
}

func Success(msg string) {
	fmt.Println(fmt.Sprintf("%s: %s .", Cheer, color.GreenString(msg)))
}
