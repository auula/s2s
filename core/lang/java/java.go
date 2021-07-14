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
package java

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
	SourceByte = `
	package model

	{{ range .Pkg }}
	import {{ . }};
	{{ end }}
	
	public class {{ .StructName | ToCamelCase }} {
	
		{{ range .Columns }}
		private {{ .Type }} {{ .Field | ToCamelCase }};
		{{ end }}
	
		{{ range .Columns }}
		public {{ .Type }} get{{ .Field | ToCamelCase}}() {
			return {{ .Field | ToCamelCase }};
		}
	
		public void set{{ .Field | ToCamelCase}}({{ .Type }} {{ .Field | ToCamelCase}}) {
			this.{{ .Field | ToCamelCase}} = {{ .Field | ToCamelCase}};
		}
		{{ end }}
	
		@Override
		public String toString() {
			return "{{ .StructName }}{" +
									{{ range .Columns }}
					"{{ .Field | ToCamelCase }}=" + {{ .Field | ToCamelCase }} + ","+
									{{ end }}
					"}";
		}
	}
	`
)

var Imports = map[string]string{
	"Time":       "java.sql.Time",
	"Date":       "java.sql.Date",
	"Float":      "java.lang.Float",
	"Double":     "java.lang.Double",
	"Boolean":    "java.lang.Boolean",
	"Timestamp":  "java.sql.Timestamp",
	"BigInteger": "java.math.BigInteger",
	"BigDecimal": "java.math.BigDecimal",
}

type Field struct {
	field   string
	kind    string
	comment string
}

func (j *Field) Field() string {
	return j.field
}

func (j *Field) Type() string {
	return j.kind
}

func (j *Field) Comment() string {
	return j.comment
}

type Assembly struct {
	lang.DataType
	source  string
	imports []string
}

func (jas *Assembly) ToField(tcs []*db.TableColumn) []core.Field {

	fieldColumn := make([]core.Field, 0, len(tcs))

	for _, column := range tcs {
		fieldColumn = append(fieldColumn, &Field{
			field:   column.ColumnName,
			kind:    jas.Table[column.DataType],
			comment: column.ColumnComment,
		})
	}

	return fieldColumn
}

func (jas *Assembly) Parse(wr io.Writer, tabName string, cs []core.Field) error {

	if tabName == "" || len(cs) <= 0 {
		return errors.New("table name info or []core.Field is empty")
	}

	// 生成模板准备解析
	tpl := template.Must(template.New(
		"s2s_java").Funcs(
		template.FuncMap{
			"ToCamelCase": core.CamelCaseFunc,
		},
	).Parse(jas.source))

	packages := make([]string, 0)

	// 自动导包优化
	isContain := func(str string, slice []string) bool {
		for _, v := range packages {
			if v == str {
				return true
			}
		}
		return false
	}

	for _, field := range cs {
		if pkg, ok := Imports[field.Type()]; ok {
			if !isContain(pkg, packages) {
				packages = append(packages, pkg)
			}
		}
	}

	type (
		structure struct {
			Pkg        []string
			StructName string
			Columns    []core.Field
		}
	)

	return tpl.Execute(wr, structure{
		Pkg:        packages,
		StructName: tabName,
		Columns:    cs,
	})
}

func NewAssembly() *Assembly {
	var jas Assembly
	jas.Lang = lang.Java
	jas.source = SourceByte
	jas.Table = lang.TypeSystem{
		"int":        "Integer",
		"tinyint":    "short",
		"smallint":   "short",
		"mediumint":  "Integer",
		"bigint":     "BigInteger",
		"bit":        "Boolean",
		"bool":       "Boolean",
		"enum":       "String",
		"set":        "String",
		"varchar":    "String",
		"char":       "String",
		"tinytext":   "String",
		"mediumtext": "String",
		"text":       "String",
		"longtext":   "String",
		"blob":       "byte[]",
		"tinyblob":   "String",
		"mediumblob": "String",
		"longblob":   "String",
		"date":       "Date",
		"datetime":   "Timestamp",
		"timestamp":  "Timestamp",
		"time":       "Time",
		"float":      "Float",
		"double":     "Double",
		"decimal":    "BigDecimal",
	}
	return &jas
}

func New() *core.Structure {
	sts := new(core.Structure)
	sts.SetLang(NewAssembly())
	return sts
}

func ToString() {
	fmt.Println(string(SourceByte))
}
