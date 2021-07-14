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
	"io"
	"strings"

	"github.com/higker/s2s/core/db"

	"github.com/higker/s2s/core/db/mysql"
)

var (
	// 驼峰英文转换函数
	CamelCaseFunc = func(str string) string {
		str = strings.Replace(str, "_", " ", -1)
		str = strings.Title(str)
		return strings.Replace(str, " ", "", -1)
	}
)

// 结构体解析接口
type Assembly interface {
	Parse(wr io.Writer, tabName string, cs []Field) error
	ToField(tcs []*db.TableColumn) []Field
}

// 抽象的结构体Field接口
type Field interface {
	Field() string
	Type() string
	Comment() string
}

type WebServer struct {
	Port string
}

type Structure struct {
	assembly Assembly
	db       db.DataBase
}

func (s *Structure) OpenDB(info *db.Info) error {

	switch info.Type {
	case db.MySQL:
		s.db = mysql.New()
	case db.PostgreSQL:
		s.db = nil
	}

	s.db.SetInfo(info)

	return s.db.Connect()
}

func (s *Structure) Close() error {
	return s.db.Close()
}

func (s *Structure) SetLang(ass Assembly) {
	s.assembly = ass
}

func (s *Structure) Parse(wr io.Writer, dbName, tabName string) error {
	columns, err := s.db.GetColumns(dbName, tabName)
	if err != nil {
		return err
	}
	return s.assembly.Parse(wr, tabName, s.assembly.ToField(columns))
}

func New() *Structure {
	sts := new(Structure)
	return sts
}

func (sts *Structure) DataBases() ([]string, error) {
	return sts.db.DataBases()
}

func (sts *Structure) Tables() ([]string, error) {
	return sts.db.Tables()
}

func (sts *Structure) SetSchema(schema string) {
	sts.db.SetSchema(schema)
}
