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
	"github.com/higker/s2s/core/assembly"
	"os"
	"text/template"
)

const (
	templateStr = `
	type {{ call ToCamelCase .structName }} struct {
		{{ range _,$value := Fields }}
		{{ value.Name }} {{ value.Type }} {{ value.Tag }}
		{{ end }}
	}
	`
)

type GoField struct {
	tag     string
	field   string
	kind    string
	comment string
}

func (gf *GoField) Field() string {
	return gf.field
}

func (gf *GoField) Type() string {
	return gf.kind
}

func (gf *GoField) Comment() string {
	return gf.comment
}

func (gf *GoField) Tag() string {
	return gf.tag
}

type Assembly struct {
	assembly.DataType
	structTpl string
}

func (goas *Assembly) ToField(tcs []*assembly.TableColumn) []assembly.Field {
	fieldColumn := make([]assembly.Field, 0, len(tcs))
	for _, column := range tcs {
		fieldColumn = append(fieldColumn, &GoField{
			tag:     fmt.Sprintf("`"+"json:"+"\"%s\""+"`", column.ColumnName),
			field:   column.ColumnName,
			kind:    goas.Table[column.DataType],
			comment: column.ColumnComment,
		})
	}

	return fieldColumn
}

func (goas *Assembly) Parse(tabName string, cs []assembly.Field) error {

	if tabName == "" || len(cs) <= 0 {
		return errors.New("table name info or []core.Field is empty")
	}

	// 生成模板准备解析
	tpl := template.Must(template.New("s2s_golang").Funcs(
		template.FuncMap{
			"ToCamelCase": assembly.CamelCaseFunc,
		},
	).Parse(goas.structTpl))

	type (
		structure struct {
			structName string
			Columns    []assembly.Field
		}
	)

	return tpl.Execute(os.Stdout, structure{
		structName: tabName,
		Columns:    cs,
	})
}

func NewAssembly() *Assembly {
	var goas Assembly
	goas.Lang = assembly.Golang
	goas.structTpl = templateStr
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
func New() *assembly.Structure {
	sts := new(assembly.Structure)
	sts.SetLang(NewAssembly())
	return sts
}
