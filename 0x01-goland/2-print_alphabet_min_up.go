package main

import "fmt"

/*
main - print alphabet min and up
*/

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Print(string(i))
		// t is equal to i, the operator := create the variable
		// and set data
		if t := i; t == 'z' {
			t = 'A'
			for ; t <= 'Z'; t++ {
				fmt.Print(string(t))
			}
		}
	}
	fmt.Println()
}
