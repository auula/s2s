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
	"github.com/higker/s2s/core"
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/lang"
	"io"
	"text/template"
)

var (
	SourceByte = `
package model

{{ range .Pkg }}
import {{ . }};
{{ end }}

public class {{ .StructName | ToCamelCase }} {

	{{ range .Columns }}
	private {{ .Type }} {{ .Field }};
	{{ end }}

	{{ range .Columns }}
	public {{ .Type }} get{{ .Field | ToCamelCase}}() {
        return {{ .Field }};
    }

    public void set{{ .Field | ToCamelCase}}({{ .Type }} {{ .Field }}) {
        this.{{ .Field }} = {{ .Field }};
    }
	{{ end }}

    @Override
    public String toString() {
        return "{{ .StructName }}{" +
				{{ range .Columns }}
                "{{ .Field }}=" + {{ .Field }} + ","+
				{{ end }}
                "}";
    }
}
	`
	//{"BigInteger", "Boolean", "Date", "Timestamp", "Time","Float", "Double"}
	Imports = make([]*AutomaticImports, 7)
)

func init() {
	Imports = []*AutomaticImports{
		&AutomaticImports{Kind: "Time", Pkg: "java.sql.Time"},
		&AutomaticImports{Kind: "Date", Pkg: "java.sql.Date"},
		&AutomaticImports{Kind: "Float", Pkg: "java.lang.Float"},
		&AutomaticImports{Kind: "Double", Pkg: "java.lang.Double"},
		&AutomaticImports{Kind: "Boolean", Pkg: "java.lang.Boolean"},
		&AutomaticImports{Kind: "Timestamp", Pkg: "java.sql.Timestamp"},
		&AutomaticImports{Kind: "BigInteger", Pkg: "java.math.BigInteger"},
	}
}

type AutomaticImports struct {
	Kind string
	Pkg  string
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
	source  []byte
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
	).Parse(string(jas.source)))

	importPkg := make([]string, 0)

	for _, field := range cs {
		for i := range Imports {
			if field.Type() == Imports[i].Kind {
				importPkg = append(importPkg, Imports[i].Pkg)
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
		Pkg:        importPkg,
		StructName: tabName,
		Columns:    cs,
	})
}

func NewAssembly() *Assembly {
	var jas Assembly
	jas.Lang = lang.Java
	jas.source = []byte(SourceByte)
	jas.Table = map[string]string{
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
	}
	return &jas
}

func New() *core.Structure {
	sts := new(core.Structure)
	sts.SetLang(NewAssembly())
	return sts
}
