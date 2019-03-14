/*
	单例模式
	故事：就比如在某个服务器程序中，该服务器的配置信息存放在一个文件中，这些配置数据由一个单例对象统一读取，然
	后服务进程中的其他对象再通过这个单例对象获取这些配置信息。这种方式简化了在复杂环境下的配置管理。

*/
package goPattern

import "sync"

var Config *config
var once sync.Once

//某个服务的配置IP和端口
type config struct{
	ip string
	port int
}

func NewConfig(ipValue string,portValue int)*config{
	once.Do(func(){Config=&config{ipValue,portValue}})
	return Config
}

