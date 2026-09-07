package main

import "fmt"

func main() {
	score := 69

	if status := "active"; score >= 70 && status == "active" {
		fmt.Println("User passed the assessment!")
	} else {
		fmt.Println("User failed or is inactive.")
	}
}
