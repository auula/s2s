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
package core

import (
	"errors"

	"github.com/higker/s2s/core/db/mysql"
)

type (
	Languages int8
	DBType    int8
)

const (
	Java Languages = iota
	Rust
	Golang
)

const (
	MySQL DBType = iota
	PostgreSQL
	DataSourceFormat = "%s:%s@tcp(%s)/information_schema?charset=%s&parseTime=True&loc=Local"
	QuerySQL         = "SELECT COLUMN_NAME, DATA_TYPE, COLUMN_KEY, " +
		"IS_NULLABLE, COLUMN_TYPE, COLUMN_COMMENT " +
		"FROM COLUMNS WHERE TABLE_SCHEMA = ? AND TABLE_NAME = ? "
)

// 结构体解析接口
type Assembly interface {
	Columns(tc []*TableColumn) []Columns
	Parse(tabName string, cs []Columns) error
}

// 抽象的结构体接口
type Columns interface {
	Tag() string
	Name() string
	Type() string
	Comment() string
}

type DataBase interface {
	Connect() error
	SetInfo(info *DBInfo)
	Close() error
	Type() DBType
}

type DBInfo struct {
	HostIPAndPort string
	UserName      string
	Password      string
	Charset       string
	Type          DBType
}

type Agrs struct{}

type TableColumn struct {
	ColumnName    string
	DataType      string
	IsNullable    string
	ColumnKey     string
	ColumnType    string
	ColumnComment string
}

type DataType struct {
	Table map[string]string
	Lang  Languages
}
type WebServer struct {
	Port string
}

type Structer struct {
	assembly Assembly
	db       DataBase
}

func (s *Structer) Open(info *DBInfo) error {

	if info == nil {
		return errors.New("database info is empty")
	}

	switch info.Type {
	case MySQL:
		s.db = mysql.New()
	case PostgreSQL:
		s.db = nil
	}

	s.db.SetInfo(info)

	return s.db.Connect()
}

func (s *Structer) Close() error {
	return s.db.Close()
}

func (s *Structer) SetLang(ass Assembly) {
	s.assembly = ass
}

func (s *Structer) Columns(tc []*TableColumn) []Columns {
	return s.assembly.Columns(tc)
}

func (s *Structer) Parse(tabName string, cs []Columns) error {
	return s.assembly.Parse(tabName, cs)
}
