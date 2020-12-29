# 抽象工厂模式
抽象工厂模式(Abstra Factory Pattern)是围绕一个超级工厂，创建其他的工厂。这种类型的设计模式属于创建型模式。  
## 优点和缺点
### 优点：
当一个产品族中多个对象被设计成一起工作时，它可以保证客户始终只使用同一个产品族中的对象
### 缺点:
产品族拓展非常困难，需要修改很多代码
### 应用场景
当需要创建的对象是一系列相互关联或相互依赖的产品族时，便可以使用抽象工厂模式。说的更明白一点，就是一个继承体系中，如果存在着多个等级结构（即存在着多个抽象类），并且分属各个等级结构中的实现类之间存在着一定的关联或者约束，就可以使用抽象工厂模式。假如各个等级结构中的实现类之间不存在关联或约束，则使用多个独立的工厂来对产品进行创建，则更合适一点。
## demo
```go
package main

import "fmt"

type XiaoMiMobile interface {
	MakingChips()
	MakingCameras()
	ManufacturingSystem()
	CompleteProduction()
}

type XiaoMiFactory interface {
	CreateMobile() XiaoMiMobile
}

type XiaoMi struct {
}

func (xiaomi *XiaoMi) MakingChips() {
	fmt.Println("小米 正在制造芯片。。。。。。。")
}

func (xiaomi *XiaoMi) MakingCameras() {
	fmt.Println("小米 正在制造相机。。。。。。。")
}

func (xiaomi *XiaoMi) ManufacturingSystem() {
	fmt.Println("小米 正在制造系统。。。。。。。")
}

func (xiaomi *XiaoMi) CompleteProduction() {
	fmt.Println("--- 小米工厂完成生产 ---")
}

func NewXiaoMiFactory() *XiaoMi {
	fmt.Println("---正在创建小米工厂---")
	return &XiaoMi{}
}

func main() {
	mi := NewXiaoMiFactory()
	mi.MakingChips()
	mi.MakingCameras()
	mi.ManufacturingSystem()
	mi.CompleteProduction()
}
