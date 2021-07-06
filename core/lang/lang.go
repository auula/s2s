// Open Source: MIT License
// Author: Jaco Ding <deen.job@qq.com>
// Date: 2021/7/6 - 4:15 下午 - UTC/GMT+08:00

// todo...

package lang

type (
	Languages int8
	DBType    int8
)

const (
	Java Languages = iota
	Rust
	Golang
)

type DataType struct {
	Table map[string]string
	Lang  Languages
}
