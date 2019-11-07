package main

import "fmt"

/*
main - print the z to a
*/

func main() {
	for i := 'z'; i >= 'a'; i-- {
		fmt.Print(string(i))
	}
	fmt.Println()
}
