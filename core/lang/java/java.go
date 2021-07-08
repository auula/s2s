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
	// Java 字节码源序列化之后的
	SourceByte = []byte{10, 112, 97, 99, 107, 97, 103, 101, 32, 109, 111, 100, 101, 108, 10, 10, 123,
		123, 32, 114, 97, 110, 103, 101, 32, 46, 80, 107, 103, 32, 125, 125, 10, 105, 109, 112, 111, 114,
		116, 32, 123, 123, 32, 46, 32, 125, 125, 59, 10, 123, 123, 32, 101, 110, 100, 32, 125, 125, 10, 10,
		112, 117, 98, 108, 105, 99, 32, 99, 108, 97, 115, 115, 32, 123, 123, 32, 46, 83, 116, 114, 117, 99,
		116, 78, 97, 109, 101, 32, 124, 32, 84, 111, 67, 97, 109, 101, 108, 67, 97, 115, 101, 32, 125, 125,
		32, 123, 10, 10, 9, 123, 123, 32, 114, 97, 110, 103, 101, 32, 46, 67, 111, 108, 117, 109, 110, 115,
		32, 125, 125, 10, 9, 112, 114, 105, 118, 97, 116, 101, 32, 123, 123, 32, 46, 84, 121, 112, 101, 32,
		125, 125, 32, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 125, 125, 59, 10, 9, 123, 123, 32, 101,
		110, 100, 32, 125, 125, 10, 10, 9, 123, 123, 32, 114, 97, 110, 103, 101, 32, 46, 67, 111, 108, 117,
		109, 110, 115, 32, 125, 125, 10, 9, 112, 117, 98, 108, 105, 99, 32, 123, 123, 32, 46, 84, 121, 112,
		101, 32, 125, 125, 32, 103, 101, 116, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 124, 32, 84, 111,
		67, 97, 109, 101, 108, 67, 97, 115, 101, 125, 125, 40, 41, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32,
		114, 101, 116, 117, 114, 110, 32, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 125, 125, 59, 10, 32,
		32, 32, 32, 125, 10, 10, 32, 32, 32, 32, 112, 117, 98, 108, 105, 99, 32, 118, 111, 105, 100, 32, 115, 101,
		116, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 124, 32, 84, 111, 67, 97, 109, 101, 108, 67, 97, 115,
		101, 125, 125, 40, 123, 123, 32, 46, 84, 121, 112, 101, 32, 125, 125, 32, 123, 123, 32, 46, 70, 105, 101,
		108, 100, 32, 125, 125, 41, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 116, 104, 105, 115, 46, 123, 123,
		32, 46, 70, 105, 101, 108, 100, 32, 125, 125, 32, 61, 32, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 125,
		125, 59, 10, 32, 32, 32, 32, 125, 10, 9, 123, 123, 32, 101, 110, 100, 32, 125, 125, 10, 10, 32, 32, 32, 32, 64,
		79, 118, 101, 114, 114, 105, 100, 101, 10, 32, 32, 32, 32, 112, 117, 98, 108, 105, 99, 32, 83, 116, 114, 105,
		110, 103, 32, 116, 111, 83, 116, 114, 105, 110, 103, 40, 41, 32, 123, 10, 32, 32, 32, 32, 32, 32, 32, 32, 114,
		101, 116, 117, 114, 110, 32, 34, 123, 123, 32, 46, 83, 116, 114, 117, 99, 116, 78, 97, 109, 101, 32, 125, 125,
		123, 34, 32, 43, 10, 9, 9, 9, 9, 123, 123, 32, 114, 97, 110, 103, 101, 32, 46, 67, 111, 108, 117, 109, 110, 115,
		32, 125, 125, 10, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 123, 123, 32, 46, 70, 105,
		101, 108, 100, 32, 125, 125, 61, 34, 32, 43, 32, 123, 123, 32, 46, 70, 105, 101, 108, 100, 32, 125, 125, 32,
		43, 32, 34, 44, 34, 43, 10, 9, 9, 9, 9, 123, 123, 32, 101, 110, 100, 32, 125, 125, 10, 32, 32, 32, 32, 32, 32, 32,
		32, 32, 32, 32, 32, 32, 32, 32, 32, 34, 125, 34, 59, 10, 32, 32, 32, 32, 125, 10, 125, 10, 9}

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
	}
	return &jas
}

func New() *core.Structure {
	sts := new(core.Structure)
	sts.SetLang(NewAssembly())
	return sts
}
