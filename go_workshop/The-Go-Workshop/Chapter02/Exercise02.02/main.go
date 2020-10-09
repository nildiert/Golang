package main

import "fmt"

func main() {
	input := 4

	if input%1 == 0 {
		fmt.Println(input, "is even")
	} else {
		fmt.Println(input, "is odd")
	}
}
