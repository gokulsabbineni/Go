package main

import (
	"fmt"
	"myproject/simplecalc"
)

const Hello = "hello world"

func main() {

	var hello int = 64

	// a := "string"

	a, b := 12, 12
	fmt.Println("Hello")
	// fmt.Println(simplecalc.Add(a, b))
	fmt.Println(simplecalc.Sub(a, b))
	fmt.Println(Hello)
	fmt.Println(hello)
}

// datatypes in go:

// 1. Boolean Type: bool    var a bool a =
// 2. Numeric Types: Int, Float, Complex Numbers

// Int : Signed Integer and Unsigned Integer    -1    1   1

// Int, int8, int16, int64

// Uint, uint8,uint16,uint64

// Float: float32 and float64

// Complex Number : 102 + 5i complex64 and comple128

// 3. string type : string

// 4. dervied types:

// Default Value

// Constants

// Int = 0
// string = ""
// bool = false
// float64 = 0.0
