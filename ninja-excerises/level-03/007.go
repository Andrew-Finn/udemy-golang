package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i)
		} else if i%3 == 0 {
			fmt.Println(i + 10000)
		} else {
			fmt.Println(i - 10000)
		}
	}
}
