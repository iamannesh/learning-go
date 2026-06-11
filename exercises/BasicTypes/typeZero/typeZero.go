package main

import "fmt"

func main(){
	var a int
	var b string
	var c bool
	var d float64
	var e byte
	var f rune

	fmt.Printf("Type of a:- %T\n", a)
	fmt.Printf("Type of b:- %T\n", b)
	fmt.Printf("Type of c:- %T\n", c)
	fmt.Printf("Type of d:- %T\n", d)
	fmt.Printf("Type of e:- %T\n", e)
	fmt.Printf("Type of f:- %T\n", f)

	fmt.Printf("Value of a:- %v\n", a)
	fmt.Printf("Value of b:- %v\n", b)
	fmt.Printf("Value of c:- %v\n", c)
	fmt.Printf("Value of d:- %v\n", d)
	fmt.Printf("Value of e:- %v\n", e)
	fmt.Printf("Value of f:- %v\n", f)
}