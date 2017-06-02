程序中记录日志的首要目的就是`Troubleshooting`，通过记录程序中对外部系统与模块的依赖调用、重要状态信息的变化、关键变量、关键逻辑等，显示基于时间轴的程序运行轨迹，显示业务是否正常、是否存在非预期执行，在出问题时方便还原现场，推断程序运行过程、理清问题的方向。

#基础

##后台输出
```
package main

import (
	"fmt"
)

func main(){
	fmt.Println("------hello world-----")
}
```

##There are no exceptions in Golang, only errors.

Go语言不支持传统的 try…catch…finally 这种异常，因为Go语言的设计者们认为，将异常与控制结构混在一起会很容易使得代码变得混乱。因为开发者很容易滥用异常，甚至一个小小的错误都抛出一个异常，替代方案是使用多值返回来返回错误。当然Go并不是全面否定异常的存在，或者用recover+panic语法实现，只是极力不鼓励多用异常。
```
package main

import (
	"log"
	"errors"
	"fmt"
)

func main() {
	/* local variable definition */
	// ...

	/* function for division which return an error if divide by 0 */
	ret,err = div(a, b)
	if err != nil {
		log.Fatal(err)
    	}
    	
    	fmt.Println(ret)
}
```

##写入日志文件:
```
package main

import (
	"log"
	"os"
)

func main(){
	f,err :=os.OpenFile("test.log",os.O_WRONLY|os.O_CREATE|os.O_APPEND,0644)
	if err !=nil{
		log.Fatal(err)
  	}
 
 	defer f.Close()
 	log.SetOutput(f)
	log.Println("==========works==============")
}
```
```
$ more test.log
2017/05/24 21:46:25 ==========works==============
```

#格式化

推荐日志工具库：`logrus`
```
$ go get github.com/Sirupsen/logrus
```

##`JSON format`
```
package main

import (
	log "github.com/Sirupsen/logrus"
	"github.com/logmatic/logmatic-go"
)

func main() {
	// use JSONFormatter
	log.SetFormatter(&logmatic.JSONFormatter{})
	// log an event as usual with logrus
	log.WithFields(log.Fields{"string": "foo", "int": 1, "float": 1.1 }).Info("My first ssl event from golang")
}
```

日志输出样式：
```
{
	"@marker":["sourcecode","golang"],
	"date":"2017-05-24T15:27:40+08:00",
	"float":1.1,"int":1,"level":"info",
	"message":"My first ssl event from golang",
	"string":"foo"
}
```

##附加上下文

通过`logrus`库可以加入一些上下文信息，例如：主机名称，程序名称或者会话参数等。
```
contextLogger := log.WithFields(log.Fields{
	"common": "XXX common content XXX",
	"other": "YYY special context YYY",
})

contextLogger.Info("AAAAAAAAAAAA")
contextLogger.Info("BBBBBBBBBBBB")
```

日志输出样式：
```
$ go run LogMatic.go
{"@marker":["sourcecode","golang"],"common":"XXX common content XXX","date":"2017-05-24T17:00:08+08:00","level":"info","message":"AAAAAAAAAAAA","other":"YYY special context YYY"}
{"@marker":["sourcecode","golang"],"common":"XXX common content XXX","date":"2017-05-24T17:00:08+08:00","level":"info","message":"BBBBBBBBBBBB","other":"YYY special context YYY"}
```

##`Hooks`

我们还可以利用`Hook`机制实现日志功能扩展，例如`Syslog hook`，将输出的日志发送到指定的`Syslog`服务。
```
package main

import (
	log "github.com/sirupsen/logrus"
	"gopkg.in/gemnasium/logrus-airbrake-hook.v2" // the package is named "aibrake"
	logrus_syslog "github.com/sirupsen/logrus/hooks/syslog"
	"log/syslog"
)

func main(){
	hook, err := logrus_syslog.NewSyslogHook("udp", "59.37.0.1:514", syslog.LOG_INFO, "")
	if err != nil {
		log.Error("Unable to connect to local syslog daemon")
	} else {
		log.AddHook(hook)
	}
}
```

验证是否发送`Syslog`：
```
$ sudo tcpdump | grep 59.37.0.1
tcpdump: data link type PKTAP
tcpdump: verbose output suppressed, use -v or -vv for full protocol decode
listening on pktap, link-type PKTAP (Apple DLT_PKTAP), capture size 262144 bytes
18:51:05.663612 IP 192.168.199.15.58819 > 59.37.0.1.syslog: SYSLOG kernel.info, length: 314
18:51:05.663657 IP 192.168.199.15.58819 > 59.37.0.1.syslog: SYSLOG kernel.info, length: 314
```

##可视化

在真实场景中日志数据体量非常庞大，日志存储只是第一步，更多的情况是需要查看特定指标或者能够快速检索信息，此时日志分析平台就发挥作用了。以`logmatic`为例，可以在它的官网注册`https://logmatic.io/`，免费体验。

在使用`logmatic`之前，需要下载它的`hook`支持：
```
$ go get github.com/logmatic/logmatic-go
```

```
func main() {
	// instantiate a new Logger with your Logmatic APIKey
	// 国内访问比较慢
	log.AddHook(logmatic.NewLogmaticHook("p53uTkOhSEqI3-116DynkQ"))

	// ..........
}
```