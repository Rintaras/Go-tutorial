package main

import "fmt"

func add(x, y int) int { //add(x int, y int) と同じ
	return x + y
}

func main() {
	fmt.Println(add(42, 13))
}
