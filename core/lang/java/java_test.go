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
package java_test

import (
	"github.com/higker/s2s/core/lang/java"
	"os"
	"testing"

	"github.com/higker/s2s/core/db"
)

var (
	structure = java.NewAssembly()
	tcs       = make([]*db.TableColumn, 3, 3)
)

func TestMain(m *testing.M) {
	tcs[0] = &db.TableColumn{
		ColumnName:    "uid",
		DataType:      "bigint",
		ColumnComment: "用户id",
		ColumnType:    "ct",
		ColumnKey:     "key",
	}
	tcs[1] = &db.TableColumn{
		ColumnName:    "account",
		DataType:      "varchar",
		ColumnComment: "用户账户",
		ColumnType:    "ct",
		ColumnKey:     "key",
	}
	tcs[2] = &db.TableColumn{
		ColumnName:    "create_time",
		DataType:      "datetime",
		ColumnComment: "用户注册时间",
		ColumnType:    "ct",
		ColumnKey:     "key",
	}
	m.Run()
}

func TestGoAssembly_ToField(t *testing.T) {

	fields := structure.ToField(tcs)
	for _, field := range fields {
		t.Log(field)
	}

}

func TestAssembly_Parse(t *testing.T) {

	structure.Parse(os.Stdout, "tableName", structure.ToField(tcs))
}
