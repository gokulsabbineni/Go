package main

import (
	"fmt"
	"myproject/simplecalc"
)

func main() {
	fmt.Println("HELLO!")
	a, b := 12.0, 3.0
	fmt.Println(simplecalc.Add(a, b))

}

// func sayHello() {
// 	fmt.Println("Hello!")
// }

// func main() {
// 	go sayHello()
// 	fmt.Println("Main function")
// }
