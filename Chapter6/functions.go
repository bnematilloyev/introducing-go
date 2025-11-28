package main

import "fmt"

/*
func main() {
	// function also called "procedure" or "subroutine"
	arr := []float64{2, 4, 6}
	avr := average(arr)
	fmt.Println(avr)
}

// My second function
func average(arr []float64) float64 {
	total := 0.0
	for _, val := range arr {
		total += val
	}
	return total / float64(len(arr))
}
*/

// Functions form call stacks
/*
	func main() {
		fmt.Println(f1())
	}
	func f1() int {
		return f2()
	}

	func f2() int {
		return 15
	}
*/

// Return types can have names
/*
	func f2() (r int) {
		r = 1
		return
	}
*/

/*
	// Multiple values can be returned
	func main() {
		x, y := f()
		fmt.Println(x, y)
	}

	func f() (int, string) {
		return 2005, "Botir"
	}
*/

// Variadic functions -> bu funksiya argumentini nechtaligi aniq bo'lmagan holatda ishlatiladi
// Using (...) elipsis

/*
	func add(args ...int) int {
		total := 0
		for _, v := range args {
			total += v
		}
		return total
	}

	func main() {
		fmt.Println(add(1, 2, 3))
		arr := []int{1, 2, 3, 4, 5}
		fmt.Println("Summa:", add(arr...))
	}
*/

// Closure
//func main() {
// First
/*
	add := func(x, y int) int {
		return x + y
	}
	fmt.Println(add(1, 2))
*/
// Second
/*
	x := 06
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
*/
//}

// Nonlocal variables -> closures
/*
func makeEvenGenerator() func() uint {
	i := uint(0)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextEven := makeEvenGenerator()
	fmt.Println(nextEven())
	fmt.Println(nextEven())
	fmt.Println(nextEven())
}
*/

// Recursion
func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func main() {
	fmt.Println(factorial(15))
}

// They both are known as functional programming
// Recursion and closures
