package main

/* #1
func sum(arr []float32) float32 {
	var summa float32
	for _, val := range arr {
		summa += val
	}
	return summa
}

func main() {
	arr := []float32{1.2, 7.5, 9.8, 75.89}
	fmt.Println(sum(arr))
}

*/

/*
#2
func half(n int) (int, bool) {
	is_odd := true
	is_odd_val := 1
	if n%2 != 0 {
		is_odd = false
		is_odd_val = 0
	}
	return is_odd_val, is_odd
}

func main() {
	fmt.Println(half(1))
}
*/

/*
#3
func maxNumber(args ...int) int {
	var max_num = args[0]
	for _, nums := range args {
		if max_num < nums {
			max_num = nums
		}
	}
	return max_num
}
func main() {
	fmt.Println(maxNumber(5, 6, 7777, 8, 11, 745, 789))
}
*/
/*
#4 makeOddGenerator()
func makeOddGenerator() func() uint {
	i := uint(1)
	return func() (ret uint) {
		ret = i
		i += 2
		return
	}
}

func main() {
	nextOdd := makeOddGenerator()
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
	fmt.Println(nextOdd())
}
*/

/*
#5
func fib(x int) int {
	if x <= 1 {
		return x
	}
	return fib(x-1) + fib(x-2)
}
func main() {
	for i := int(-3); i <= 10; i++ {
		fmt.Println(i, fib(i))
	}
}
*/

/*
6. What are defer, panic, and recover? How do you recover from a runtime panic?

Defer - is used to schedule a function to run after
the surrounding function returns (whether normally or because of a panic).

	func main() {
	    defer fmt.Println("I run last")
	    fmt.Println("I run first")
	}

Panic - stops normal execution immediately and starts unwinding the stack,
running all deferred functions along the way.

	func test() {
	    panic("Something went wrong")
	}

Recover is used inside a deferred function to stop a panic
and regain normal program flow.
recover ONLY works inside a defer function.

func safe() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered from:", r)
		}
	}()

	// This will cause a panic
	var x []int
	fmt.Println(x[1])
}

func main() {
	safe()
}


Another real case usage

func divide(a, b int) (result int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Error occured:", r)
			result = 0
		}
	}()
	return a / b
}

func main() {
	fmt.Println(divide(10, 0))
	fmt.Println("Program continuos ... ")
}


// Q: How do you get the memory address of a variable?
// A: with & sign
func main() {
	x := 17
	fmt.Println(&x) //0xc0000120e8
}
*/

// Q: How do you assign a value to a pointer?
// A: with (*) asterisk

/* Code below
func pointer(xPt *int) {
	*xPt = 1
}

func main() {
	x := 5
	pointer(&x)
	fmt.Println(x)
}
*/

// Q: How do you create a new pointer?
// A: Use the built-in new function

/* Answer for Q9
func new_one(xPtr *int) {
	*xPtr = 1
}

func main() {
	xPtr := new(int)
	new_one(xPtr)
	fmt.Println(*xPtr)
}
*/

// Q10
/*
func square(x *float64) {
	*x = *x * *x
}
func main() {
	x := 1.5
	square(&x)
	fmt.Println(x)
}
*/
// Response is 2.25

// Q11 Swap two integers

/*
func swap(x, y *int) {
	*x, *y = *y, *x
}

func main() {
	x := 25
	y := 15
	swap(&x, &y)
	fmt.Println(x, y) // 15 25
}
*/
