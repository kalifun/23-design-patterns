package command

import "fmt"

// command 
type command interface {
	execute()
}

// button 
type button struct {
	command command
}

// press 
//  @receiver b 
func (b *button) press() {
	b.command.execute()
}

// device 
type device interface {
	on()
	off()
}

// onCommand 
type onCommand struct {
	device device
}

// execute 
//  @receiver c 
func (c *onCommand) execute() {
	c.device.on()
}

// offCommand 
type offCommand struct {
	device device
}

// execute 
//  @receiver c 
func (c *offCommand) execute() {
	c.device.off()
}

// tv 
type tv struct {
	isRunning bool
}

// on 
//  @receiver t 
func (t *tv) on() {
	t.isRunning = true
	fmt.Println("Turmomh tv on")
}

// off 
//  @receiver t 
func (t *tv) off() {
	t.isRunning = false
	fmt.Println("Turning tv off")
}


