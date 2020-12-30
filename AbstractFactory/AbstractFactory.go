package main

import "fmt"

type MobileGroup interface {
	MakingMobileChips()
	MakingCameras()
	ManufacturingSystem()
	CompleteProduction()
}

type Mi11Mobile struct {
}

func (mi11 *Mi11Mobile) MakingMobileChips() {
	fmt.Println("小米11 正在制造芯片。。。。。。。")
}

func (mi11 *Mi11Mobile) MakingCameras() {
	fmt.Println("小米11 正在制造相机。。。。。。。")
}

func (mi11 *Mi11Mobile) ManufacturingSystem() {
	fmt.Println("小米11 正在制造系统。。。。。。。")
}

func (mi11 *Mi11Mobile) CompleteProduction() {
	fmt.Println("--- 小米11工厂完成生产 ---")
}

type WifiGroup interface {
	MakingWifiChips()
	MakeTheAntenna()
}

type RedMi struct {
}

func (red *RedMi) MakingWifiChips() {
	fmt.Println("红米路由器 正在制造Wi-Fi芯片。。。。。")
}

func (red *RedMi) MakeTheAntenna() {
	fmt.Println("红米路由器 正在制造Wi-Fi天线。。。。。")
}

type HighTechFactory interface {
	CreateMobile() MobileGroup
	CreateWifi() WifiGroup
}

type XiaoMi struct {
}

func (mi *XiaoMi) CreateMobile() MobileGroup {
	fmt.Println("--- 创建小米11手机工厂 ---")
	return new(Mi11Mobile)
}

type HongMi struct {
}

func (hongmi *HongMi) CreateWifi() WifiGroup {
	fmt.Println("--- 创建红米路由器工厂 ---")
	return new(RedMi)
}

func main() {
	a := new(XiaoMi)
	c := a.CreateMobile()
	c.MakingMobileChips()
	c.MakingCameras()
	c.ManufacturingSystem()
	c.CompleteProduction()
	// --- 创建小米11手机工厂 ---
	// 小米11 正在制造芯片。。。。。。。
	// 小米11 正在制造相机。。。。。。。
	// 小米11 正在制造系统。。。。。。。
	// --- 小米11工厂完成生产 ---

	v := new(HongMi)
	t := v.CreateWifi()
	t.MakingWifiChips()
	t.MakeTheAntenna()
	// --- 创建红米路由器工厂 ---
	// 红米路由器 正在制造Wi-Fi芯片。。。。。
	// 红米路由器 正在制造Wi-Fi天线。。。。。
}
