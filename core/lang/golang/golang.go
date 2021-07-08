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
	"errors"
	"fmt"
	"io"
	"text/template"

	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/lang"
)

var (
	// 种子不要改！！
	SourceByte = []byte{
		10, 9, 112, 97, 99, 107, 97, 103, 101, 32, 109, 111, 100, 101,
		108, 10, 9, 10, 9, 105, 109, 112, 111, 114, 116, 32, 40, 10, 9,
		9, 34, 101, 110, 99, 111, 100, 105, 110, 103, 47, 106, 115, 111,
		110, 34, 10, 9, 41, 10, 10, 9, 116, 121, 112, 101, 32, 123, 123,
		32, 46, 83, 116, 114, 117, 99, 116, 78, 97, 109, 101, 32, 124, 32,
		84, 111, 67, 97, 109, 101, 108, 67, 97, 115, 101, 32, 125, 125, 32,
		115, 116, 114, 117, 99, 116, 32, 123, 10, 9, 9, 123, 123, 32, 114, 97,
		110, 103, 101, 32, 46, 67, 111, 108, 117, 109, 110, 115, 32, 125, 125,
		10, 9, 9, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 124, 32, 84, 111,
		67, 97, 109, 101, 108, 67, 97, 115, 101, 32, 125, 125, 9, 123, 123, 32, 46,
		84, 121, 112, 101, 32, 125, 125, 32, 123, 123, 32, 46, 84, 97, 103, 32, 32,
		125, 125, 10, 9, 9, 123, 123, 32, 101, 110, 100, 32, 125, 125, 10, 9, 125, 10,
		10, 9, 102, 117, 110, 99, 32, 40, 123, 123, 32, 46, 83, 116, 114, 117, 99, 116,
		78, 97, 109, 101, 32, 125, 125, 32, 42, 123, 123, 32, 46, 83, 116, 114, 117, 99,
		116, 78, 97, 109, 101, 32, 124, 32, 84, 111, 67, 97, 109, 101, 108, 67, 97, 115, 101,
		32, 125, 125, 41, 32, 84, 97, 98, 108, 101, 78, 97, 109, 101, 40, 41, 32, 115, 116, 114,
		105, 110, 103, 32, 123, 10, 9, 9, 114, 101, 116, 117, 114, 110, 32, 34, 123, 123, 32, 46,
		83, 116, 114, 117, 99, 116, 78, 97, 109, 101, 32, 124, 32, 84, 111, 67, 97, 109, 101, 108,
		67, 97, 115, 101, 32, 125, 125, 34, 10, 9, 125, 10, 10, 9, 102, 117, 110, 99, 32, 40, 123,
		123, 32, 46, 83, 116, 114, 117, 99, 116, 78, 97, 109, 101, 32, 125, 125, 32, 42, 123, 123,
		32, 46, 83, 116, 114, 117, 99, 116, 78, 97, 109, 101, 32, 124, 32, 84, 111, 67, 97, 109, 101,
		108, 67, 97, 115, 101, 32, 125, 125, 41, 32, 84, 111, 74, 115, 111, 110, 40, 41, 32, 115, 116,
		114, 105, 110, 103, 32, 123, 10, 9, 9, 115, 116, 114, 44, 95, 32, 58, 61, 32, 106, 115, 111, 110,
		46, 77, 97, 114, 115, 104, 97, 108, 40, 123, 123, 32, 46, 83, 116, 114, 117, 99, 116, 78, 97, 109,
		101, 32, 125, 125, 41, 10, 9, 9, 114, 101, 116, 117, 114, 110, 32, 115, 116, 114, 10, 9, 125, 10, 9}
)

type Field struct {
	tag     string
	field   string
	kind    string
	comment string
}

func (gf *Field) Field() string {
	return gf.field
}

func (gf *Field) Type() string {
	return gf.kind
}

func (gf *Field) Comment() string {
	return gf.comment
}

func (gf *Field) Tag() string {
	return gf.tag
}

type Assembly struct {
	lang.DataType
	source []byte
}

func (gas *Assembly) ToField(tcs []*db.TableColumn) []core.Field {
	fieldColumn := make([]core.Field, 0, len(tcs))
	for _, column := range tcs {
		fieldColumn = append(fieldColumn, &Field{
			tag:     fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName),
			field:   column.ColumnName,
			kind:    gas.Table[column.DataType],
			comment: column.ColumnComment,
		})
	}

	return fieldColumn
}

func (gas *Assembly) Parse(wr io.Writer, tabName string, cs []core.Field) error {

	if tabName == "" || len(cs) <= 0 {
		return errors.New("table name info or []core.Field is empty")
	}

	// 生成模板准备解析
	tpl := template.Must(template.New("s2s_golang").Funcs(
		template.FuncMap{
			"ToCamelCase": core.CamelCaseFunc,
		},
	).Parse(string(gas.source)))

	type (
		structure struct {
			StructName string
			Columns    []core.Field
		}
	)

	return tpl.Execute(wr, structure{
		StructName: tabName,
		Columns:    cs,
	})
}

func NewAssembly() *Assembly {
	var gas Assembly
	gas.Lang = lang.Golang
	gas.source = SourceByte
	gas.Table = lang.TypeSystem{
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
	return &gas
}

func New() *core.Structure {
	sts := new(core.Structure)
	sts.SetLang(NewAssembly())
	return sts
}
