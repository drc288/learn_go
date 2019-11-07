package main

import "fmt"

/*
main - print 0 to 9 using while ..... the while not exists
using for but.. while syntax is iqual
*/

func main() {
	i := 0
	for i < 10 {
		fmt.Print(i)
		i += 1
	}
	fmt.Println()
}
