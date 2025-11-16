package main

import "fmt"

func main() {
	/* // First way to make slices
	x := make([]float64, 5, 10) // 3rd parameter for capacity
	fmt.Println(x)
	*/

	// Second way

	/*
		arr := [5]float64{1, 2, 3, 4, 5}
		x := arr[0:5]
		fmt.Println(x)
		fmt.Println(arr[1:4])

		fmt.Println(arr[0:]) // fmt.Println(arr[0:len(arr)]) are same
		fmt.Println(arr[:5]) // fmt.Println(arr[0:5]) are same
		fmt.Println(arr[:])  // fmt.Println(arr[0:len(arr)]) are same
	*/

	// append
	/*
		slice1 := []int{1, 2, 3}
		slice2 := append(slice1, 4, 5)
		fmt.Println(slice1, slice2)
	*/

	// copy takes two arguments: dst and src
	slice1 := []int{1, 2, 3}
	slice2 := make([]int, 7)
	fmt.Println(slice2)
	copy(slice2, slice1)
	fmt.Println(slice1, slice2)
}
