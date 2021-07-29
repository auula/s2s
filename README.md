# s2s

## 什么是`s2s`？
`s2s (sql to structure)`是一款命令行数据库逆向工程工具，它可以通过数据库表生成对应的`Java`、`Go`、`Rust`结构体（`class`），后面将陆续支持更多的语言。

## 配置数据库源

`s2s`依赖于你的数据库，所以需要你配置好你的数据库连接信息，以便`s2s`会正常的运行。配置信息方法很简单你只需要在你的环境变量中加入以下信息即可。


```bash
#s2s 命令的数据库信息
export s2s_host="127.0.0.1:3306"
export s2s_user="root"
export s2s_pwd="you db password"
export s2s_charset="utf8"
```

`windows`的配置`此电脑->属性->高级系统设置->环境变量`，`Mac`和`Linux`则在`~/.profile`或者`~/.zshrc`中添加以上配置信息即可。

## 使用方法

你可以克隆下载本代码库，然后如果你的电脑上已经安装好了`go`的编译器那么就进入主目录即可使用`go build`命令编译生成二进制程序文件。

如果你觉得麻烦即可在下面列表中找到你对应的平台架构下载对应的二进制可执行文件到电脑上，如果你想在系统上随意调用你则只需要把`s2s`的安装目录放入你的环境变量中。

|  平台   | 地址  |
|  ----  | ----  |
| windows-64  | 单元格 |
| Mac-64  | 单元格 |

## 导入包

本库支持你二次开发使用，你只需要导入本包即可在你的代码中进行扩充开发，但是目前仅支持`go`语言！

1. 下载

```bash
go get -u github.com/higker/s2s
```

2. 导入并且使用

```go
package main

import (
	"github.com/higker/s2s/core/lang/java"
	"github.com/higker/s2s/core/db"
)


func main() {

    // 创建一个Java的代码生成器
    structure := java.New()
    
    // 数据库连接信息
    if err := structure.OpenDB(
        &db.Info{
            HostIPAndPort: os.Getenv("s2s_host"), // 数据库IP
            UserName:      os.Getenv("s2s_user"), // 数据库用户名
            Password:      os.Getenv("s2s_pwd"),  // 数据库密码
            Type:          db.MySQL,              // 数据库类型 PostgreSQL Oracle
            Charset:       os.Getenv("s2s_charset"),
        },
    ); err != nil {
        // Failed to establish a connection to the database!
        // .... logic code
    }
		
    defer structure.Close()
    
    structure.SetSchema("选择数据库名")
    
    // 生成结果输出到标准输出
    structure.Parse(os.Stdout,"表名")

}
```
