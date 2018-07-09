package main

import "fmt"

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) speak() {
	fmt.Println("My name is", p.firstName, p.lastName, "and I am a person.")
}

func (sa secretAgent) speak() {
	fmt.Printf("%s. %s %s. Secret agent: %v\n", sa.lastName, sa.firstName, sa.lastName, sa.licenseToKill)
}

type human interface {
	speak()
}

func sayStuff(h human) {
	h.speak()
}

func main() {
	p := person{"James", "Lucktaylor"}
	sa := secretAgent{person{"Alec", "Trevelyan"}, true}

	sayStuff(p)
	sayStuff(sa)
}
