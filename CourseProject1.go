// This is a profit calculator application
package main

import (
	"fmt"
)

func main() {
	var revenue float64
	var expenses float64
	var tax_rate float64

	//fmt.Print("Enter revenue : ")
	OutputText("Enter revenue : ")
	fmt.Scan(&revenue)

	//fmt.Print("Enter expenses : ")
	OutputText("Enter expenses : ")
	fmt.Scan(&expenses)

	//fmt.Print("Enrer tax rate : ")
	OutputText("Enter tax rate : ")
	fmt.Scan(&tax_rate)
	EBT, profit, ratio := Calculations(revenue, expenses, tax_rate)
	//var EBT float64 = revenue - expenses
	//var profit float64 = EBT * (1 - tax_rate)

	//var ratio float64 = EBT / profit

	fmt.Printf("The earnings before taxes are : %.1f \n", EBT)
	fmt.Printf("The profit is : %.1f \n", profit)
	fmt.Printf("The ratio of EBT to profit is : %.3f \n", ratio)
}

func OutputText(Text string) {
	fmt.Print(Text)
}

func Calculations(revenue, expenses, tax_rate float64) (EBT, profit, ratio float64) {
	EBT = revenue - expenses
	profit = EBT * (1 - tax_rate)
	ratio = EBT / profit
	return
}
