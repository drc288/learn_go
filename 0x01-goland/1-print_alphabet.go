package main

import "fmt"

/*
main - print alphabet
*/

func main() {
	for i := 'a'; i <= 'z'; i++ {
		fmt.Print(string(i))
	}
	fmt.Println()
}
