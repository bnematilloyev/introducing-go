package main

/*
func first(x int) {
	x += 1
}

func second(y int) {
	y += 1
}

func main() {
	x := 1
	y := 1
	// defer -> kechiktirmoq
	defer second(y)
	first(x)
	defer second(y)
	first(x)
	fmt.Println(x)
	fmt.Println(y)
}
*/

// panic and recover

func main() {
	/*
		panic("Panic")
		str := recover()
		fmt.Println(str)
	*/
	// panic: Panic
	// panic immediately stops execution of the function

	/*
			defer func() {
				str := recover()
				fmt.Println(str)
			}()
			panic("PANIC")
			// PANIC
		// A panic generally indicates a programmer error
		// or an exceptional
		//condition that there’s no easy way to recover from (hence the name panic)
	*/

}
