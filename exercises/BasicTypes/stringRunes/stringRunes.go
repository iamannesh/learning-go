package main

import "fmt"
func main(){
	var str string = "hello"
	var l int = len(str)
	fmt.Printf("Length of String:- %v\n", l)
	for i := 0; i < l; i++ {
		fmt.Printf("bytes:- %v\n", str[i])
	}

	for i, v := range str {
		fmt.Printf("index:- %v, value:- %v\n", i, v)
	}
}