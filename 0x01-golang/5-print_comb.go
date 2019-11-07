package main

import "fmt"

/*
main - print to 0 from 9 and print
*/

func main() {
	for i := 0; i < 10; i++ {
		fmt.Print(i)
		if i != 9 {
			fmt.Print(", ")
		} else {
			fmt.Println()
		}
	}
}
