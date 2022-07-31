package main

import "fmt"

func main() {
	fmt.Println(add(15, 4))
}

func add(x int, y int) int {
	fmt.Println("İşlem Başarılı!")
	if x < 6 && x > 0 || y > 16 {
		return x + y
	} else if x > 10 || y < 5 {
		return x / y
	} else if x == 10 || y == 4 {
		return x * y
	} else if x < 0 && y > 0 {
		return (x * y * -1)
	} else {
		return 0
	}
}
