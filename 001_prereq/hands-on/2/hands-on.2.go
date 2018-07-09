package main

import (
	"fmt"
)

type person struct {
	firstName string
	lastName  string
}

type secretAgent struct {
	person
	licenseToKill bool
}

func (p person) pSpeak() {
	fmt.Println("My name is", p.firstName, p.lastName, "and I am a person.")
}

func (sa secretAgent) saSpeak() {
	fmt.Printf("%s. %s %s. Secret agent.\n", sa.lastName, sa.firstName, sa.lastName)
}

func main() {
	p := person{"James", "Lucktaylor"}
	sa := secretAgent{person{"Alec", "Trevelyan"}, true}

	fmt.Println(p.firstName)

	p.pSpeak()

	fmt.Println(sa.lastName)

	sa.saSpeak()
	sa.pSpeak()
}
