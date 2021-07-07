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

import "github.com/higker/s2s/core/lang"

var (
	SourceByte = `
public class User {

    private Integer id;

    private String account;

    private String nickName;


    public Integer getId() {
        return id;
    }

    public void setId(Integer id) {
        this.id = id;
    }

    public String getAccount() {
        return account;
    }

    public void setAccount(String account) {
        this.account = account;
    }

    public String getNickName() {
        return nickName;
    }

    public void setNickName(String nickName) {
        this.nickName = nickName;
    }

    @Override
    public String toString() {
        return "User{" +
                "id=" + id +
                ", account='" + account + '\'' +
                ", nickName='" + nickName + '\'' +
                '}';
    }
}
	`
)

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
	source []byte
}

func NewAssembly() *Assembly {
	var jas Assembly
	jas.Lang = lang.Java
	jas.source = []byte(SourceByte)
	jas.Table = map[string]string{
		"int":        "Integer",
		"tinyint":    "Integer",
		"smallint":   "Integer",
		"mediumint":  "Integer",
		"bigint":     "java.math.BigInteger",
		"bit":        "java.lang.Boolean",
		"bool":       "java.lang.Boolean",
		"enum":       "String",
		"set":        "String",
		"varchar":    "String",
		"char":       "String",
		"tinytext":   "String",
		"mediumtext": "String",
		"text":       "String",
		"longtext":   "String",
		"blob":       "java.lang.byte[]",
		"tinyblob":   "String",
		"mediumblob": "String",
		"longblob":   "String",
		"date":       "java.sql.Date",
		"datetime":   "java.sql.Timestamp",
		"timestamp":  "java.sql.Timestamp",
		"time":       "java.sql.Time",
		"float":      "java.lang.Float",
		"double":     "java.lang.Double",
	}
	return &jas
}
