package main

import "fmt"
func main() {
	var i float64 = 3.9
	var f int = int(i)
	var u rune = rune(i)

	fmt.Printf("Value of f:- %v\n", f)
	fmt.Printf("Value of u:- %v\n", u)
}