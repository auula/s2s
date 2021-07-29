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

package db

type (
	Type int8
)

const (
	MySQL Type = iota
	PostgreSQL
	DataSourceFormat = "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	QuerySQL         = "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
)

type Info struct {
	HostIPAndPort string
	UserName      string
	Password      string
	Charset       string
	Type          Type
}

type DataBase interface {
	Type() Type
	Close() error
	Connect() error
	SetInfo(info *Info)
	SetSchema(schema string)
	Tables() ([]string, error)
	DataBases() ([]string, error)
	GetColumns(tabName string) ([]*TableColumn, error)
}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

func (t *TableColumn) Name() string {
	return t.ColumnName
}

func (t *TableColumn) Type() string {
	return t.ColumnType
}

func (t *TableColumn) Comment() string {
	return t.ColumnComment
}
