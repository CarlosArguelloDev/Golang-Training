package main

import "fmt"

func CheckParity(number int) string {
	switch {
	case number == 0:
		return "Zero"
	case number%2 == 0:
		return "Even"
	default:
		return "Odd"
	}
}

func main() {
	num := 10
	fmt.Printf("The number %d is : %s\n", num, CheckParity(num))

}
