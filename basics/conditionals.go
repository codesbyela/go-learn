package main

import "fmt"

func main() {
	age := 25
	if age < 18 {
		fmt.Println("You are a minor.")
	} else if age < 65 {
		fmt.Println("You are an adult.")
	} else {
		fmt.Println("You are a senior citizen.")
	}
}