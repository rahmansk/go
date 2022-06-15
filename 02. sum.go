// Write a Go Program to find sum of digits in a number

package main

import "fmt"

func main() {
	var num, sum, reminder int

	fmt.Print("Enter the Number to find the sum of digits = ")
	fmt.Scan(&num)

	for sum = 0; num > 0; num = num / 10 {
		reminder = num % 10
		sum = sum + reminder
	}
	fmt.Println("The sum of digit : ", sum)
}
