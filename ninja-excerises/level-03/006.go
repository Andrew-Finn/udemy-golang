package main

import "fmt"

func main() {
	for i := 10; i <= 100; i++ {
		if i%2 == 0 {
			fmt.Println(i, i%4)
		}
	}
}
