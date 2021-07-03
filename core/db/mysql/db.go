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
package mysql

import (
	"database/sql"
	"errors"
	"fmt"

	"github.com/higker/s2s/core"
)

type DB struct {
	source *sql.DB
	info   *core.DBInfo
}

func New() core.DataBase {
	return new(DB)
}

func (db *DB) Type() core.DBType {
	return db.info.Type
}

func (db *DB) Connect() error {
	var err error
	if db.info == nil {
		return errors.New("database info is empty")
	}
	dsn := fmt.Sprintf(core.DataSourceFormat,
		db.info.UserName,
		db.info.Password,
		db.info.HostIPAndPort,
		db.info.Charset,
	)
	db.source, err = sql.Open("mysql", dsn)
	return err
}

func (db *DB) SetInfo(info *core.DBInfo) {
	db.info = info
}

func (db *DB) Close() error {
	return db.source.Close()
}

func (db *DB) GetColumns(dbName, tableName string) ([]*core.TableColumn, error) {
	rows, err := db.source.Query(core.QuerySQL, dbName, tableName)
	if err != nil {
		return nil, err
	}
	if rows == nil {
		return nil, errors.New("no data")
	}
	defer rows.Close()

	var columns []*core.TableColumn
	for rows.Next() {
		var column core.TableColumn
		err := rows.Scan(&column.ColumnName, &column.ColumnType, &column.ColumnKey, &column.IsNullable, &column.ColumnType, &column.ColumnComment)
		if err != nil {
			return nil, err
		}

		columns = append(columns, &column)
	}

	return columns, nil
}
