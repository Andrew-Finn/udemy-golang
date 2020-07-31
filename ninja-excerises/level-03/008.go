package main

import "fmt"

func main() {
	switch {
	case false:
		fmt.Println("No print")
	case true:
		fmt.Println("Prints")
	}
}
