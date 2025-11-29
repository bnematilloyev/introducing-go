package main

import "fmt"

//type Circle struct {
//	x, y, r float64
//}
//
//type Rectangle struct {
//	x1, y1, x2, y2 float64
//}

/*
func main() {
	// Structs -> Obyektga o'xshaydi, lekin obyekt emas.
	// Odatda turli xildagi ma'lumot turlarini saqlovchi quti vazifasini bajaradi

	/* Both are similar
	type Circle struct {
		x float64
		y float64
		z float64
	}
*/

/*
	type Circle struct {
		x, y, z float64
	}
*/

// var c Circle
// c := new(Circle) // Another way

// Pointers are often used with structs so that
// functions can modify their contents.

// Pointer — bu o‘zgaruvchining qiymatini emas, balki uning
// xotiradagi manzilini saqlovchi o‘zgaruvchi.

/*
	x := 10
	p := &x  // p — x ning manzili
*/

// Giving initial value
// c := Circle{x: 0, y: 0, z: 5}
// Second option
// c := Circle{0, 0, 5} // &Circle bu pointerni qaytaradi

// Fields
/*
		fmt.Println(c.x, c.y, c.z)
		c.x = 10
		c.y = 25
		fmt.Println(c.x, c.y, c.z)
}
*/

/*
// CircleArea function
func circleArea(c Circle) float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 0, 5}
	fmt.Println(circleArea(c))
}
*/

// Methods
/*
func (c *Circle) area() float64 {
	return math.Pi * c.r * c.r
}

func main() {
	c := Circle{0, 5, 5}
	fmt.Println(c.area())
}
*/

// With rectangle func
/*
type Rectangle struct {
	x1, y1, x2, y2 float64
}

func distance(x1, y1, x2, y2 float64) float64 {
	a := x2 - x1
	b := y2 - y1
	return math.Sqrt(a*a + b*b)
}

func (r *Rectangle) area() float64 {
	l := distance(r.x1, r.y1, r.x2, r.y2)
	w := distance(r.x1, r.y1, r.x2, r.y1)
	return l * w
}

func main() {
	r := Rectangle{15, 25, 20, 30}
	fmt.Println(r.area())
}
*/

// Embeded Types
// Go tilidagi Embedded Types (yoki embedding, yoki “kompozitsiya”) — bu Go’da meros olish (inheritance)
// yo‘qligi o‘rnini bosadigan mexanizm.

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hi, my name is", p.Name)
}

/*
func main() {
	p := Person{"Botir"}
	p.Talk() //Hi, my name is Botir
}
*/

type Android struct {
	Person
	Model string
}

// First way
/*
func main() {
	a := new(Android)
	a.Person.Name = "Botir"
	a.Person.Talk()
}
*/

// Second way
func main() {
	a := Android{
		Person: Person{Name: "Botir"},
		Model:  "x-500",
	}
	a.Talk()
}
