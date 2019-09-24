package main

import (
	"fmt"
)

type Anyg struct {}

func (a *Anyg) Any() {
	fmt.Println("Any:any")
}

type Someg struct{}

func (a *Someg) Some() {
	fmt.Println("Some:some")
}

type All struct{}

func (a *All) Any() {
	fmt.Println("All:any")
}

func (a *All) Some() {
	fmt.Println("All:some")
}

// a struct include these sturct which have same name methods
type Collection struct {
	Someg
	p All
	Anyg
}

func (c *Collection) Any(){
	fmt.Println("collection:any")
}

func (c *Collection) Some() {
	fmt.Println("collection:some")
}

// Documents some line to test.
// Then know how can use embed struct.
func TestStruct(){
	c  := &Collection{
		Someg{},
		All{},
		Anyg{},
	}
	c.Any()
	c.Some()
	c.p.Any()
	c.p.Some()
}

// Let's learn this skills how to work in RealWorld
// Type a interface first.
type Ready interface {
	Any()
	Some()
	Readiness() bool
}

// There is a struct implements Ready
type AlwaysReady struct{}
func (a *AlwaysReady)Any() {}
func (a *AlwaysReady)Some() {}
func (a *AlwaysReady) Readiness()bool {
	return true
}

// We use All struct to implements our own methods logic.
// and use AlwaysReady to implement Ready interface.
type RealWorldUse struct{
	p All
	AlwaysReady
}


func EmbedInRealWorld(embedReady Ready){
	fmt.Println(embedReady.Readiness())
}

func main() {

	TestStruct() 


	// ----------------------------------------------------------------
	embed := &RealWorldUse{
		All{},
		AlwaysReady{},
	}
 
	EmbedInRealWorld(embed)
}
