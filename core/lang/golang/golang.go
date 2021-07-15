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
	// code template
	SourceByte = `
	package model
	
	import (
		"encoding/json"
	)

	type {{ .StructName | ToCamelCase }} struct {
		{{ range .Columns }}
		{{ .Field | ToCamelCase }}	{{ .Type }} {{ .Tag  }}
		{{ end }}
	}

	func ({{ .StructName }} *{{ .StructName | ToCamelCase }}) TableName() string {
		return "{{ .StructName | ToCamelCase }}"
	}

	func ({{ .StructName | ToCamelCase }} *{{ .StructName | ToCamelCase }}) ToJson() string {
		str,_ := json.Marshal({{ .StructName }})
		return str
	}
`
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
	Source string
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
	).Parse(gas.Source))

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
	gas.Source = SourceByte
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

func ToString() {
	fmt.Println(string(SourceByte))
}
