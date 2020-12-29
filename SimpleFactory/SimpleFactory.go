package main

import "fmt"

func main() {
	// A car has four wheels
	NewTheTrafficTools("c").ToGetParts()
	// A bicycle has two wheels
	NewTheTrafficTools("b").ToGetParts()
}

type TheTrafficTools interface {
	ToGetParts()
}

type Car struct {
}

func (c *Car) ToGetParts() {
	fmt.Println("A car has four wheels")
}

type Bicycle struct {
}

func (b *Bicycle) ToGetParts() {
	fmt.Println("A bicycle has two wheels")
}

func NewTheTrafficTools(name string) TheTrafficTools {
	switch name {
	case "c":
		return &Car{}
	case "b":
		return &Bicycle{}
	default:
		return nil
	}
}
