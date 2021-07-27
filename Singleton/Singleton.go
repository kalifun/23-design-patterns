package singleton

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
