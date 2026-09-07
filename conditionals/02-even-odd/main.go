package main

import "fmt"

func main() {
	number := 0
	if number == 0 {
		fmt.Println("It´s Zero")
	} else if number%2 == 0 {
		fmt.Println("It's Even")
	} else {
		fmt.Println("it's Odd")
	}
}
