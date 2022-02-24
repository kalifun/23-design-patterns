package composite

import (
	"fmt"
)

// file
type file struct {
	name string
}

// search
//  @receiver f
//  @param keyword
func (f *file) search(keyword string) {
	fmt.Printf("Searching for keyword %s in file %s\n", keyword, f.name)
}

// getName
//  @receiver f
//  @return string
func (f *file) getName() string {
	return f.name
}

// component
type component interface {
	search(string)
}

// folder
type folder struct {
	components []component
	name       string
}

// search
//  @receiver f
//  @param keyword
func (f *folder) search(keyword string) {
	fmt.Printf("Serching recursively for keyword %s in folder %s\n", keyword, f.name)
	for _, composite := range f.components {
		composite.search(keyword)
	}
}

// add 
//  @receiver f 
//  @param c 
func (f *folder) add(c component) {
	f.components = append(f.components, c)
}

