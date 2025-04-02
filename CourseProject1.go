// This is a profit calculator application
package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	fmt.Print("Enter revenue : ")
	fmt.Scan(&revenue)

	fmt.Print("Enter expenses : ")
	fmt.Scan(&expenses)

	fmt.Print("Enrer tax rate : ")
	fmt.Scan(&tax_rate)

	var EBT float64 = revenue - expenses
	var profit float64 = EBT * (1 - tax_rate)

	var ratio float64 = EBT / profit

	fmt.Println("The earnings before taxes are : ", EBT)
	fmt.Println("The profit is : ", profit)
	fmt.Println("The ratio of EBT to profit is : ", ratio)
}
