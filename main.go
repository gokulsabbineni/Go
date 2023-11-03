package main

import "fmt"

func main() {
	g := map[string]float64{
		"x": 92.3,
		"y": 99.8,
		"z": 91.2,
	}

	g["w"] = 85.1

	xvalue := g["x"]
	fmt.Println(xvalue)

	g["x"] = 90.5

	delete(g, "y")

	for student, grade := range g {
		fmt.Println(student, grade)
	}

	if grade, ok := g["z"]; ok {
		fmt.Println(grade, ok)
	} else {
		fmt.Println("not found")
	}

}

// Pointers

// Note:
// default value of a pointer is nil
// we cann't increment or decrement a pointer
// j : 0x20e456734 = 20
// p : 0x20e456734

// * -> holds the value
// & -> holds the address

// Struct

// var bookTitle = "hello"
// var
// 0 1 2 3 4 -- index of an arr
// 1 2 3 4 5 -- values which are holding in the index

// Slices:
// length and capacity

// length: 3

// s:=make([]int, 3, 5)

// s = {1,2,3}    ----- lenght 3 and capacity 3
// s = {1,2,3,4,5} ----- length 5 and capacity 6

// s = {1,2,3,4,5,6} --- length 6 and capacity 10

// 4 ways you can declare a slice:

// 1. var s []int
// 2. s:=make([]int, 3, 5)
// 	s := make([]int,3)
// 3. s := []int{1,2,3}
// 4.

// 0 1 2 3 4  ----Index

// 1 2 3 4 5------ Elements

// Maps:
// a := make(map[string]string)
// map literals
// a := map[string]int{
// 	"one" : 1
// 	"two" : 2
// }
