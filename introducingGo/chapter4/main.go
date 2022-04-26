package main

import "fmt"

func main() {
	fizzBuzz()
}

func fizzBuzz() {
	for i := 1; i <= 100; i++ {
		is_fizz := i%3 == 0
		is_buzz := i%5 == 0
		if is_fizz && is_buzz {
			fmt.Println(i, "FizzBuzz")
			continue
		}
		if is_fizz {
			fmt.Println(i, "Fizz")
		}
		if is_buzz {
			fmt.Println(i, "Buzz")
		}
	}
}

func multOfThree() {
	for i := 1; i <= 100; i++ {
		if i%3 == 0 {
			fmt.Println(i, " is divisible by 3")
		}
	}
}

func chapLesson() {
	for i := 0; i <= 10; i++ {
		if i%2 == 0 {
			fmt.Println("even")
		} else {

			fmt.Println("odd")
		}

	}
}
