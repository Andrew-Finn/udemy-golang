package main

import "fmt"

func main() {
	m := map[string][]string{
		`andrew`: []string{`Random1`, `Random2`},
	}
	fmt.Println(m)

	m["john"] = []string{"random3", `random4`}

	fmt.Println(m)
}