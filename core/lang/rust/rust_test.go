// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/7/8 - 4:23 下午 - UTC/GMT+08:00

package rust_test

import (
	"github.com/higker/s2s/core/db"
	"github.com/higker/s2s/core/lang/rust"
	"os"
	"testing"
)

var (
	structure = rust.NewAssembly()
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

	structure.Parse(os.Stdout, "UserInfo", structure.ToField(tcs))
}
