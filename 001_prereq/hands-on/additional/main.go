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
	eleven()
	twelve()
	// thirteen()
	fourteen()
	fifteen()
	sixteen()
	seventeen()
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

type gator int

func eleven() {
	var g1 gator

	g1 = 42

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
}

func twelve() {
	var g1 gator
	var x int

	g1 = 42
	x = 27

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func thirteen() {
	// var g1 gator
	// var x int

	// g1 = 42

	// x = g1
}

func fourteen() {
	var g1 gator
	var x int

	g1 = 42
	x = 27

	fmt.Println(g1)
	fmt.Printf("%T\n", g1)
	fmt.Println(x)
	fmt.Printf("%T\n", x)

	x = int(g1)

	fmt.Println(x)
	fmt.Printf("%T\n", x)
}

func (g gator) greeting() {
	fmt.Println("Hello, I am a gator")
}

func fifteen() {
	var g1 gator = 99

	g1.greeting()
}

type flamingo bool

func (f flamingo) greeting() {
	fmt.Println("Hello, I am pink and beautiful and wonderful")
}

type swampCreature interface {
	greeting()
}

func bayou(sc swampCreature) {
	sc.greeting()
}

func sixteen() {
	var g1 gator = 99
	var f1 flamingo = true

	bayou(g1)
	bayou(f1)
}

func seventeen() {
	s := "i'm sorry dave i can't do that"

	fmt.Println(s)
	fmt.Println([]byte(s))
	fmt.Println(string([]byte(s)))

	fmt.Println(s[:14])
	fmt.Printf("'%s'\n", s[10:22])
	fmt.Println(s[17:])

	for _, x := range s {
		fmt.Println(string(x), x)
	}
}
