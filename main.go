package main

import "fmt"

func main() {
	fmt.Println("Test")
	var a = 2
	var b = 3
	fmt.Println(fmt.Sprintf("%d + %d = %d", a, b, calc(a, b, 1)))
	fmt.Println(fmt.Sprintf("%d * %d = %d", a, b, calc(a, b, 3)))
}

func calc(a int, b int, tipe int) int {
	// 1: tambah, 2: kurang, 3: kali, 4: bagi
	switch tipe {
	case 1:
		return a + b
	case 2:
		return a - b
	case 3:
		return a * b
	default:
		return a / b
	}
}
