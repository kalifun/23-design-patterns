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
