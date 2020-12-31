package main

import "fmt"

type Person struct {
	Type string
	Sex  string
}

func (p *Person) SetType(s string) {
	p.Type = s
}

func (p *Person) GetType() string {
	return p.Type
}

func (p *Person) SetSex(s string) {
	p.Sex = s
}

func (p *Person) GetSex() string {
	return p.Sex
}

type BuildPerson interface {
	BuildType()
	BuildSex()
	CreatePerson() *Person
}

type HeroHuman struct {
	person *Person
}

func CreateHeroHuman() *HeroHuman {
	return &HeroHuman{person: new(Person)}
}

func (hero *HeroHuman) BuildType() {
	hero.person.SetType("Hero")
}

func (hero *HeroHuman) BuildSex() {
	hero.person.SetSex("Male")
}

func (hero *HeroHuman) CreatePerson() *Person {
	return hero.person
}

type PersonController struct {
}

func (pc *PersonController) construct(b BuildPerson) *Person {
	b.BuildSex()
	b.BuildType()
	return b.CreatePerson()
}

func main() {
	h := CreateHeroHuman()
	c := new(PersonController)
	f := c.construct(h)
	fmt.Println(f.Sex)
	fmt.Println(f.GetType())
}
