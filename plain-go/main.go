package main

import "fmt"

func main() {
	fmt.Println("Hello world")
}

func EvenOrOdd(num int) string {
	if num%2 == 0{
		return "EVEN"
	}
	return "ODD"
}

func hello() string  {
	return "Hello World"
}