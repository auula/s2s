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
package rust

import (
	"errors"
	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/lang"
	"io"
	"text/template"
)

var (
	SourceByte = `
// Special type view: https://docs.rs/chrono

extern crate serde_derive;

#[derive(Serialize, Deserialize)]
pub struct {{ .StructName | ToCamelCase}} {

	{{ range .Columns }}
	// {{ .Comment }}
    #[serde(rename = "{{ .Field }}")]
    pub {{ .Field }}: {{ .Type }},
	{{ end }}

}
`
)

type Field struct {
	kind    string
	field   string
	comment string
}

func (r *Field) Field() string {
	return r.field
}

func (r *Field) Type() string {
	return r.kind
}

func (r *Field) Comment() string {
	return r.comment
}

type Assembly struct {
	lang.DataType
	source  []byte
	comment string
}

func (ras *Assembly) ToField(tcs []*db.TableColumn) []core.Field {

	fieldColumn := make([]core.Field, 0, len(tcs))

	for _, column := range tcs {
		fieldColumn = append(fieldColumn, &Field{
			field:   column.ColumnName,
			kind:    ras.Table[column.DataType],
			comment: column.ColumnComment,
		})
	}

	return fieldColumn
}

func (ras *Assembly) Parse(wr io.Writer, tabName string, cs []core.Field) error {

	if tabName == "" || len(cs) <= 0 {
		return errors.New("table name info or []core.Field is empty")
	}

	// 生成模板准备解析
	tpl := template.Must(template.New(
		"s2s_rust").Funcs(
		template.FuncMap{
			"ToCamelCase": core.CamelCaseFunc,
		},
	).Parse(string(ras.source)))

	importPkg := make([]string, 0)

	type (
		structure struct {
			Pkg        []string
			StructName string
			Columns    []core.Field
		}
	)

	return tpl.Execute(wr, structure{
		Pkg:        importPkg,
		StructName: tabName,
		Columns:    cs,
	})
}

func NewAssembly() *Assembly {
	var jas Assembly
	jas.Lang = lang.Rust
	jas.source = []byte(SourceByte)
	jas.Table = lang.TypeSystem{
		"int":        "i32",
		"tinyint":    "i8",
		"smallint":   "i16",
		"mediumint":  "i64",
		"bigint":     "i64",
		"bit":        "i16",
		"bool":       "bool",
		"enum":       "String",
		"set":        "String",
		"varchar":    "String",
		"char":       "String",
		"tinytext":   "String",
		"mediumtext": "String",
		"text":       "String",
		"longtext":   "String",
		"blob":       "&[u8]",
		"tinyblob":   "String",
		"mediumblob": "String",
		"longblob":   "String",
		"date":       "chrono::NaiveDate",
		"datetime":   "chrono::NaiveDateTime",
		"timestamp":  "chrono::NaiveDateTime",
		"time":       "chrono::NaiveTime",
		"float":      "f32",
		"double":     "f64",
	}
	return &jas
}

func New() *core.Structure {
	sts := new(core.Structure)
	sts.SetLang(NewAssembly())
	return sts
}
