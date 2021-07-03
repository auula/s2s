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
package golang

import (
	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/db/mysql"
)

type Assembly struct {
	core.DataType
	structTpl string
}

func (goas *Assembly) Columns(tc []*core.TableColumn) []core.Columns {
	return nil
}

func (goas *Assembly) Parse(tabName string, cs []core.Columns) error {
	return nil
}

func NewAssembly() *Assembly {
	var goas Assembly
	goas.Lang = core.Golang
	goas.structTpl = ""
	goas.Table = map[string]string{
		"int":        "int32",
		"tinyint":    "int8",
		"smallint":   "int",
		"mediumint":  "int64",
		"bigint":     "int64",
		"bit":        "int",
		"bool":       "bool",
		"enum":       "string",
		"set":        "string",
		"varchar":    "string",
		"char":       "string",
		"tinytext":   "string",
		"mediumtext": "string",
		"text":       "string",
		"longtext":   "string",
		"blob":       "string",
		"tinyblob":   "string",
		"mediumblob": "string",
		"longblob":   "string",
		"date":       "time.Time",
		"datetime":   "time.Time",
		"timestamp":  "time.Time",
		"time":       "time.Time",
		"float":      "float64",
		"double":     "float64",
	}
	return &goas
}

func New() (sts *core.Structer) {
	sts.Assembly = NewAssembly()
	sts.DB = mysql.New()
	return
}
