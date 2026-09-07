package main

import "fmt"

func main() {
	number := 10

	if number == 0 {
		fmt.Println("its zero D:")
	} else if number%2 == 0 {
		fmt.Println("its even")
	} else {
		fmt.Println("its odd")
	}
}
