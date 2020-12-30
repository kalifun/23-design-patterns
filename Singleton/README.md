# 单例模式
单例模式是一种常用的软件设计模式，其定义是单例对象的类只能允许一个实例存在。

许多时候整个系统只需要拥有一个的全局对象，这样有利于我们协调系统整体的行为。比如在某个服务器程序中，该服务器的配置信息存放在一个文件中，这些配置数据由一个单例对象统一读取，然后服务进程中的其他对象再通过这个单例对象获取这些配置信息。这种方式简化了在复杂环境下的配置管理。

## 优点和缺点
### 优点：
1.在内存中只有一个对象，节省内存空间；  
2.避免频繁的创建销毁对象，可以提高性能；  
3.避免对共享资源的多重占用，简化访问；  
4.为整个系统提供一个全局访问点。
### 缺点:
1.不适用于变化频繁的对象；  
2.滥用单例将带来一些负面问题，如为了节省资源将数据库连接池对象设计为的单例类，可能会导致共享连接池对象的程序过多而出现连接池溢出；  
3.如果实例化的对象长时间不被利用，系统会认为该对象是垃圾而被回收，这可能会导致对象状态的丢失；
### 应用场景
1.需要生成唯一序列的环境。  
2.需要频繁实例化然后销毁的对象。  
3.创建对象时耗时过多或者耗资源过多，但又经常用到的对象。   
4.方便资源相互通信的环境

## demo
```go
package main

import "fmt"

// 定义一个接口
type LoadBalancing interface {
	AddNode(string)
	GetNode() string
}

// 声明一个节点的存储
type LoadBalancingType struct {
	server map[string]struct{}
}

func (s *LoadBalancingType) AddNode(svr string) {
	s.server[svr] = struct{}{}
}

func (s *LoadBalancingType) GetNode() string {
	for srv := range s.server {
		return srv
	}
	return ""
}

var lb LoadBalancing

func GetLoadBalancing() LoadBalancing {
	return lb
}

func NewLoadBalancing() *LoadBalancingType {
	m := make(map[string]struct{})
	return &LoadBalancingType{server: m}
}

func init() {
	lb = NewLoadBalancing()
}

func main() {
	lb := GetLoadBalancing()
	lb.AddNode("server 1")
	lb.AddNode("server 2")
	lb.AddNode("server 3")

	fmt.Println(lb.GetNode())
}
```