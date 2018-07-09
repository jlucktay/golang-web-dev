package main

import "fmt"

func main() {
	one()
	two()
	four()
	five()
	six()
	eight()
	nine()
	ten()
}

func one() {
	xi := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(xi)

	for a := range xi {
		fmt.Println("index:", a)
	}

	for a, b := range xi {
		fmt.Printf("xi[%v]: %v\n", a, b)
	}
}

func two() {
	m := map[string]int{
		"One":   1,
		"Two":   2,
		"Three": 3,
		"Four":  4,
		"Five":  5,
		"Six":   6,
		"Seven": 7,
	}

	for a := range m {
		fmt.Println(a)
	}

	for a, b := range m {
		fmt.Printf("%v = %v\n", a, b)
	}
}

type person struct {
	fName, lName string
}

func four() {
	p1 := person{"Firstname", "Lastname"}

	fmt.Println(p1)
	fmt.Println(p1.fName)
}

type personFive struct {
	person
	favFood []string
}

func five() {
	p1 := personFive{
		person{
			"Firstname",
			"Lastname",
		},
		[]string{
			"pizza",
			"dumplings",
			"rice",
		},
	}

	fmt.Println(p1.favFood)

	for f, g := range p1.favFood {
		fmt.Printf("food '%d' = '%s'\n", f, g)
	}
}

type personSix struct {
	personFive
}

func (ps personSix) walk() string {
	return fmt.Sprintf("%s is walking.", ps.fName)
}

func six() {
	s := personSix{personFive{person{"Firstname", "Lastname"}, []string{"apples", "oranges"}}}.walk()

	fmt.Println(s)
}

type vehicle struct {
	doors int
	color string
}

type sedan struct {
	vehicle
	luxury bool
}

type truck struct {
	vehicle
	fourWheel bool
}

func eight() {
	t := truck{vehicle{2, "red"}, true}
	s := sedan{vehicle{4, "blue"}, true}

	fmt.Println(t)
	fmt.Println(s)
	fmt.Println(t.color)
	fmt.Println(s.luxury)
}

func (t truck) transportationDevice() string {
	return fmt.Sprintf("I am a %s truck with %d doors. Do I have four-wheel drive? %v", t.color, t.doors, t.fourWheel)
}

func (s sedan) transportationDevice() string {
	return fmt.Sprintf("I am a %s sedan with %d doors. Am I a luxury model? %v", s.color, s.doors, s.luxury)
}

func nine() {
	t := truck{vehicle{2, "red"}, true}
	s := sedan{vehicle{4, "blue"}, true}

	fmt.Println(t.transportationDevice())
	fmt.Println(s.transportationDevice())
}

type transportation interface {
	transportationDevice() string
}

func report(t transportation) {
	fmt.Println(t.transportationDevice())
}

func ten() {
	t := truck{vehicle{2, "red"}, true}
	s := sedan{vehicle{4, "blue"}, true}

	report(t)
	report(s)
}
