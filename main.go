package main

import "fmt"

func main() {
	fmt.Println("Test")
	fmt.Println(fmt.Sprintf("1 + 2 = %d", calc(1, 2)))
}

func calc(a int, b int) int {
	return a + b
}
