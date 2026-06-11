package main

import "fmt"
func main() {
	var price float64
	var quantity int
	var total float64
	var discountApplied bool

	fmt.Print("Enter the price and quantity:-")
	fmt.Scan(&price, &quantity)

	total = price * float64(quantity)
	if total>100.0 {
		discountApplied = true
		total = total * 0.9
	}
	if discountApplied==true {
		fmt.Printf("%v x $%v = $%v (discount applied)\n", quantity, price, total)
	} else {
		fmt.Printf("%v x $%v = $%v \n", quantity, price, total)
	}
}