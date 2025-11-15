package main

import "fmt"

func main() {
	//	var name type = value
	/*
		var ism string = "Botir, Nematilloyev"
		var age int
		age = 20
		fmt.Println("Ismi: ", ism, "\nYosh: ", age)
	*/

	//	Resetting value to variable
	/*
		var x string
		x = "First"
		fmt.Println(x)
		x = "Second"
		fmt.Println(x)
	*/

	//First
	//Second

	/*
		var x string
		x = "First "
		fmt.Println(x)
		x += "Second"
		fmt.Println(x)
	*/

	//First
	//First Second

	//Assigning value with another way
	// type was not set-up, but go compiler identify by self

	//x := "Botir"
	//fmt.Println(x)

	//Constants
	//const MUSIC_PLAYER string = "KONSTA"
	//fmt.Println(MUSIC_PLAYER)

	// Defining multiple variables or constants
	//var (
	//	a = 5
	//	b = 7
	//	c = "hello"
	//)
	//fmt.Println(a, b, c)
	//
	//const (
	//	d = 5
	//	e = 7
	//	f = "hello"
	//)
	//fmt.Println(d, e, f)

	// Scanf usage
	//fmt.Println("Son kiriting: ")
	//var input float64
	//fmt.Scanf("%f", &input)
	//
	//output := input * 2
	//fmt.Println(output)

	//Exercises
	//#5
	//var F float32
	//fmt.Println("Farangentni kiriting: ")
	//fmt.Scanf("%f", &F)
	//Celsius := (F - 32) * 5 / 9
	//fmt.Println(Celsius)

	//#6
	var metr float64
	fmt.Println("Metrni kiriting: ")
	fmt.Scanf("%f", &metr)
	ft := metr * 0.3048
	fmt.Println(ft)
}
