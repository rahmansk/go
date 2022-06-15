// Taking a year (eg. 2020) as input from user, write a Go program  to see if it is a leap year. Use Switch Case to print leap/not a leap year

package main

import "fmt"

func main() {
	var year int
	fmt.Print("Enter the year to be checked:")
	fmt.Scanf("%d", &year)
	if year%4 == 0 && year%100 != 0 || year%400 == 0 {
		fmt.Println("The year is a leap year!")
	} else {
		fmt.Println("The year isn't a leap year!")
	}
}
