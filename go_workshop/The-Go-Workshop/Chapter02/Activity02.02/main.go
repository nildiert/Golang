package main

import "fmt"

func main() {
	words := map[string]int{
		"Gonna": 3,
		"You":   3,
		"Give":  2,
		"Never": 1,
		"Up":    4,
	}

	count := -1
	mostPopularWord := ""

	for key, value := range words {
		if value > count {
			mostPopularWord = key
			count = value
		}
	}
	fmt.Println("Most popular word: ", mostPopularWord)
	fmt.Println("With a count of :", count)
}
