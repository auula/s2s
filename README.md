# s2s
[![DeepSource](https://deepsource.io/gh/higker/s2s.svg/?label=active+issues&show_trend=true)](https://deepsource.io/gh/higker/s2s/?ref=repository-badge)
[![DeepSource](https://deepsource.io/gh/higker/s2s.svg/?label=resolved+issues&show_trend=true)](https://deepsource.io/gh/higker/s2s/?ref=repository-badge)
[![License](https://img.shields.io/badge/license-MIT-db5149.svg)](https://github.com/higker/sessionx/blob/master/LICENSE)
[![Go Reference](https://pkg.go.dev/badge/github.com/higker/s2s.svg)](https://pkg.go.dev/github.com/higker/s2s)

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
 
1. 你可以克隆下载本代码库，然后如果你的电脑上已经安装好了`go`的编译器那么就进入主目录即可使用`go build`命令编译生成二进制程序文件。

2. 如果你觉得麻烦即可在下面列表中找到你对应的平台架构下载对应的二进制可执行文件到电脑上，如果你想在系统上随意调用你则只需要把`s2s`的安装目录放入你的环境变量中。

3. 目前对`Rust`部分数据类型支持不够友好，不过不耽误使用，目前被生成的数据库表名格式必须为`user_info`这样的`snake case`这种格式！！后面会考虑修复这个`bug`。

| 平台       | 地址   |
| ---------- | ------ |
| Windows-x64 | [s2s-windows-x64.zip](https://github.com/higker/s2s/releases/download/v0.0.1/s2s-windows-x64.zip) |
| Mac-x64     | [s2s-darwin-x64.zip](https://github.com/higker/s2s/releases/download/v0.0.1/s2s-darwin-x64.zip) |
| Linux-64     | [s2s-linux-x64.zip](https://github.com/higker/s2s/releases/download/v0.0.1/s2s-linux-x64.zip) |



## 内置命令

**PS: 在命令行模式下按下`tab`键会有命令补全提示！**

| 命令      | 使用方法                 |
| --------- | ------------------------ |
| databases | 显示所有数据库名         |
| use       | 指定使用哪个数据库       |
| tables    | 显示当前数据库所有表     |
| gen       | 生成指定的表，`gen 表名` |
| info      | 显示数据库所有信息       |
| exit      | 退出命令行模式           |
| clear     | 清理屏幕内容             |

**使用案例**

```bash
$:> s2s java

	        ______
	.-----.|__    |.-----.
	|__ --||    __||__ --|
	|_____||______||_____|



🥳: You have entered the command line mode!

🥳: Press the 'tab' key to get a prompt！

🥳: Enter `exit` to exit the program!

😃:s2s>databases
+---+--------------------+
| * | Database           |
+---+--------------------+
| 1 | information_schema |
| 2 | emp_db             |
| 3 | mysql              |
| 4 | performance_schema |
| 5 | sys                |
| 6 | test_db            |
+---+--------------------+


😃:s2s>use emp_db

🤖‍: Selected as database 👉 `emp_db`！

😃:s2s>tables
+---+-----------+
| * | Tables    |
+---+-----------+
| 1 | user_info |
+---+-----------+


😃:s2s>gen user_info

	package model


	import java.sql.Timestamp;

	import java.math.BigDecimal;

	import java.math.BigInteger;


	public class UserInfo {


		// 用户账号
		private String Account;

		// 用户创建时间
		private Timestamp CreateTime;

		// 用户更新时间
		private Timestamp UpdatedDate;

		// 用户年龄
		private short Age;

		// 用户余额
		private BigDecimal Money;

		// 用户ID
		private BigInteger Uid;



		public String getAccount() {
			return Account;
		}

		public void setAccount(String Account) {
			this.Account = Account;
		}

		public Timestamp getCreateTime() {
			return CreateTime;
		}

		public void setCreateTime(Timestamp CreateTime) {
			this.CreateTime = CreateTime;
		}

		public Timestamp getUpdatedDate() {
			return UpdatedDate;
		}

		public void setUpdatedDate(Timestamp UpdatedDate) {
			this.UpdatedDate = UpdatedDate;
		}

		public short getAge() {
			return Age;
		}

		public void setAge(short Age) {
			this.Age = Age;
		}

		public BigDecimal getMoney() {
			return Money;
		}

		public void setMoney(BigDecimal Money) {
			this.Money = Money;
		}

		public BigInteger getUid() {
			return Uid;
		}

		public void setUid(BigInteger Uid) {
			this.Uid = Uid;
		}


		@Override
		public String toString() {
			return "user_info{" +

					"Account=" + Account + ","+

					"CreateTime=" + CreateTime + ","+

					"UpdatedData=" + UpdatedDate + ","+

					"Age=" + Age + ","+

					"Money=" + Money + ","+

					"Uid=" + Uid + ","+

					"}";
		}
	}

😃:s2s>exit

🤖‍: Bye👋 :)
```

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

## 其他

目前仅支持`mysql`数据库，如果有想贡献代码提`issues`！跟多需求： 1. 支持`linux`管道命令这样就可以可编程操作了，前面一个输出就是后面一个程序的输入。
