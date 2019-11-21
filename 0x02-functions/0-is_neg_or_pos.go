package main

import "fmt"

/*
pos_or_neg - verifi if the number are positive, negative or is a 0
num: intger positive or negative
*/

func pos_or_neg(num int) {
	if num == 0 {
		fmt.Println("The number is 0")
	} else if num < 0 {
		fmt.Println("The number is negative")
	} else {
		fmt.Println("The number is positive")
	}
}

/*
main - call a function and verify if the number is
positive or negative
*/

func main() {
	pos_or_neg(0)
}
