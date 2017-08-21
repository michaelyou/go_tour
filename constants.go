package main

import "fmt"

const Pi = 3.14

func main() {
	const World = "世界"
	fmt.Println("hello", World)
	fmt.Println("happy", Pi, "day")

	// cannot assign to World
	// World = "好的"

	const Truth = true
	fmt.Println("go rules?", Truth)
}
