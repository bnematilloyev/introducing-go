package main

import (
	"fmt"
	"sort"
)

/*
func main() {
	// Doubly Linked List
	// NULL ← [Prev | Value | Next] ↔ [Prev | Value | Next] ↔ [Prev | Value | Next] → NULL
		var x list.List
		x.PushBack(1)
		x.PushBack(2)
		x.PushBack(3)

		// Boshiga, o'rtaga va oxiriga qo'shish O(1)
		for e := x.Front(); e != nil; e = e.Next() {
			fmt.Println(e.Value.(int))
		}
		fmt.Println(x) // {{0xc00012a120 0xc00012a180 <nil> <nil>} 3}
}
*/

// Sort
type Person struct {
	Name string
	Age  int
}

// Sorting -> ByName
type ByName []Person

func (ps ByName) Len() int {
	return len(ps)
}

func (ps ByName) Less(i, j int) bool {
	return ps[i].Name < ps[j].Name
}
func (ps ByName) Swap(i, j int) {
	ps[i], ps[j] = ps[j], ps[i]
}

// Sorting ByAge

type ByAge []Person

func (this ByAge) Len() int {
	return len(this)
}

func (this ByAge) Less(i, j int) bool {
	return this[i].Age < this[j].Age // Ascending
}

func (this ByAge) Swap(i, j int) {
	this[i], this[j] = this[j], this[i]
}

func main() {
	people := []Person{
		{"Bahrom", 23},
		{"Botir", 20},
		{"Bahodir", 77},
	}
	sort.Sort(ByName(people))
	fmt.Println(people)
	sort.Sort(ByAge(people))
	fmt.Println(people)
}
