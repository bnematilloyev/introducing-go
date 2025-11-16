package main

import "fmt"

func main() {
	/*
		// Maps
		mp := make(map[string]int)
		mp["index"] = 10
		mp["age"] = 20
		fmt.Println(mp)

		// delete function
		delete(mp, "age")
		fmt.Println(mp)
	*/
	//week_days := make(map[string]string)
	//week_days["Mo"] = "Monday"
	//week_days["Te"] = "Tuesday"
	//week_days["Wn"] = "Wensday"
	//week_days["Th"] = "Thursday"
	//week_days["Fr"] = "Friday"
	//week_days["St"] = "Saturday"
	//week_days["Sn"] = "Sunday"

	// reflect.TypeOf(week_days["Fo"]) // -> This is for getting Type
	// import reflect for usage

	// fmt.Println(week_days["Fo"])

	/*
		name, ok := week_days["Mo"]
		fmt.Println(name, ok)
		// Monday true

		name, ok = week_days["Mn"]
		if ok == true {
			fmt.Println(name)
		} else {
			fmt.Println("not exist")
		}
		// not exist
	*/

	//if name, ok := week_days["Mo"]; ok {
	//	fmt.Println(name, ok)
	//}

	// Another way to create maps
	/*
		week_days := map[string]string{
			"Monday":    "Mon",
			"Tueasday":  "Tue",
			"Wednesday": "Wed",
			"Thursday":  "Thu",
			"Friday":    "Fri",
			"Saturday":  "Sat",
			"Sunday":    "Sun",
		}
		fmt.Println(week_days)
	*/
	/*
		elements := map[string]map[string]string{
			"H": map[string]string{
				"name":  "Hydrogen",
				"state": "gas",
			},
			"He": map[string]string{
				"name":  "Helium",
				"state": "gas",
			},
			"Li": map[string]string{
				"name":  "Lithium",
				"state": "solid",
			},
			"Be": map[string]string{
				"name":  "Beryllium",
				"state": "solid",
			},
		}
		fmt.Println(elements["Li"]["state"])
	*/

	// Exercises
	//#1
	//var arr = []int{1, 2, 3, 4, 5, 6}
	//fmt.Println(arr[3]) // 4

	//#2
	//var arr = make([]int, 3, 9)
	//fmt.Println(len(arr), cap(arr))
	// length -> 3 |  capacity -> 9

	//#3
	//x := [6]string{"a", "b", "c", "d", "e", "f"}
	//fmt.Println(x[2:5])
	// [c d e]

	// #4
	x := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	minimum := x[0]
	for _, v := range x {
		if v < minimum {
			minimum = v
		}
	}
	fmt.Println(minimum)
}
