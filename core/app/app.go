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

package app

import (
	"github.com/fatih/color"
)

const (
	version = "0.0.1 Alpha"

	banner = `
	        ______        
	.-----.|__    |.-----.
	|__ --||    __||__ --|
	|_____||______||_____|
`

	description = `
S2S is a command-line database tool that can generate tables 
with the type of Golang,Rust, Java language,with the command line.`
)

var Info *S2SInfo = NewInfo()

type S2SInfo struct {
	Version, Banner, Description string
}

func NewInfo() *S2SInfo {
	return &S2SInfo{
		Version:     color.YellowString(version),
		Banner:      color.MagentaString(banner),
		Description: color.MagentaString(banner) + color.GreenString(description),
	}
}
