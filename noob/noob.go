package main

import (
	"fmt"
	"time"
)

func main() {
	defer func(start time.Time) {
		fmt.Printf("elapsed: %s\t", time.Since(start))
	}(time.Now())
	// var n int = 19
	// fmt.Println(n)

	// var greeting string = "hej"
	// fmt.Println(greeting)

	// var numbers []int = []int{1, 2, 3, 4}
	// fmt.Println(numbers)

	var greetings []string = []string{"Hej", "Hallå", "Hi", "Hola", "Hello"}
	var names []string = []string{"Ove", "Bob", "Alice", "Ben", "Bertil"}

	for _, greeting := range greetings {
		for _, name := range names {
			fmt.Println(greeting, name)
		}
	}
}
