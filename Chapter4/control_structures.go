package main

func main() {
	// each number has own line
	/*
		fmt.Println(1)
		fmt.Println(2)
		fmt.Println(3)
		fmt.Println(4)
		fmt.Println(5)
		fmt.Println(6)
		fmt.Println(7)
		fmt.Println(8)
	*/

	// Another way || pretty tedious
	//	fmt.Println(`1
	//2
	//3
	//4
	//5
	//6
	//7
	//8
	//`)

	// The for statement (loop)
	//i := 1
	//for i <= 10 {
	//	fmt.Println(i)
	//	i += 1
	//}

	// My way :D
	//for i := 1; i <= 10; i++ {
	//	fmt.Println(i)
	//}

	// The if statement
	//for i := 1; i <= 10; i++ {
	//	if i%2 == 0 {
	//		fmt.Println("Even")
	//	} else if i%5 == 0 {
	//		fmt.Println("My fovurite one")
	//	} else {
	//		fmt.Println("Odd")
	//	}
	//}

	// The switch Statement
	/*
		var x int
		fmt.Println("Enter the number: ")
		fmt.Scanf("%d", &x)

		switch x {
		case 0:
			fmt.Println("Zero")
		case 1:
			fmt.Println("One")
		case 2:
			fmt.Println("Two")
		case 3:
			fmt.Println("Three")
		case 4:
			fmt.Println("Four")
		default:
			fmt.Println("My lovely")
		}
	*/

	// Exercises
	//#1
	//i := 25
	//if i > 10 {
	//	fmt.Println("Big")
	//} else {
	//	fmt.Println("Small")
	//}

	//#2
	//for i := 3; i <= 100; i += 3 {
	//	fmt.Println(i)
	//}

	//#3 FizzBuzz algorithm
	//for i := 1; i <= 100; i++ {
	//	if i%3 == 0 && i%5 == 0 {
	//		fmt.Println("FizzBuzz")
	//	} else if i%3 == 0 {
	//		fmt.Println("Fizz")
	//	} else if i%5 == 0 {
	//		fmt.Println("Buzz")
	//	} else {
	//		fmt.Println("Pretty")
	//	}
	//}
}
