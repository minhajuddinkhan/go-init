package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p *Person) talk() {
	fmt.Println("My Name is", p.name)
}

type coolPerson interface {
	sayYolo()
}

func main() {

	var a Person = Person{"Min", 20}
	a.talk()
	fmt.Println(a.age)
}
