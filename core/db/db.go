// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/7/6 - 4:11 下午 - UTC/GMT+08:00

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
	GetColumns(dbName, tabName string) ([]*TableColumn, error)
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
