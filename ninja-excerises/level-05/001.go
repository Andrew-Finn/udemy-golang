package main

import "fmt"

type person struct {
	first string
	last string
	favFlavours []string
}
func main() {
	p1 := person{
		first: "John",
		last: "Smith",
		favFlavours: []string{"1", "2", ",3"},
	}
	p2 := person{
		first: "Pat",
		last: "Bond",
		favFlavours: []string{"4", "5", ",6"},
	}

	fmt.Println(p1, p2)
}
