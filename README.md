## Time4Go

Go 语言的时间工具库。

## 帮助 
在集成的过程中有遇到问题，欢迎加 QQ 群 564704807 讨论。

## 安装
```bash
$ go get github.com/smartwalle/time4go
```

## 开始
```go
package main

import (
	"fmt"
	"github.com/smartwalle/time4go"
	"time"
)

func main() {
	var now = time4go.Now()
	fmt.Println(now)

	var d = time4go.Date(2018, time.May, 20, 13, 14, 0, 0, time.Local)
	fmt.Println(d)
}
```

#### 获取当前日期所在周的第一天和最后一天
```go
var now = time4go.Now()
now.BeginningDateOfWeek()
now.EndDateOfWeek()
```


#### 获取当前日期所在月的第一天和最后一天
```go
var now = time4go.Now()
now.BeginningDateOfMonth()
now.EndDateOfMonth()
```

#### JSON

* 设置序列化成 JSON 字符串时的格式

```go
time4go.JSONFormatter = time4go.DefaultFormatter{Layout: "2006-01-02 15:04:05"}
```

例子:

```
package main

import (
	"encoding/json"
	"fmt"
	"github.com/smartwalle/time4go"
)

func main() {
	time4go.JSONFormatter = time4go.DefaultFormatter{Layout: "2006-01-02 15:04:05"}

	var s = &Schedule{}
	s.Begin = time4go.Now()
	s.End = s.Begin.AddDate(0, 1, 0)

	sBytes, _ := json.Marshal(s)
	fmt.Println(string(sBytes))

	var ts = `{"begin":"2019-11-10 09:59:21","end":"2019-12-10 09:59:21"}`
	var s2 *Schedule
	json.Unmarshal([]byte(ts), &s2)
	fmt.Println(s2.Begin, s2.End)
}

type Schedule struct {
	Begin *time4go.Time `json:"begin"`
	End   *time4go.Time `json:"end"`
}

```

* 自定义 Formatter

当然你也可以自定义 Formatter，只需要实现以下接口即可：

```go
type TimeFormatter interface {
	Format(t time.Time) ([]byte, error)
	Parse(data []byte) (time.Time, error)
}
```

比如：

```go
type MyFormatter struct {
}

func (this MyFormatter) Format(t time.Time) ([]byte, error) {
	...
}

func (this MyFormatter) Parse(data []byte) (time.Time, error) {
	...
}

time4go.JSONFormatter = MyFormatter{}
```

#### 支持 SQL 类数据库

```go
db, err := sql.Open("mysql", "xxx")
if err != nil {
	fmt.Println("连接数据库出错：", err)
	return
}
defer db.Close()
db.Exec("INSERT INTO `user` (`name`, `age`, `created_on`) VALUES (?, ?, ?)", "test", 18, time4go.Now())
```

## License
This project is licensed under the MIT License.