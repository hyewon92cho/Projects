package main

import (
	"fmt"
)

func main() {
	var num int
	print("Enter a number: ")
	fmt.Scanln(&num)
	if num <= 0 {
		println("Invalid input!")
	} else {
		fmt.Println(fibonacci(num))
	}
}

func fibonacci(size int) []int {
	arr := make([]int, size)
	for index := 0; index < size; index++ {
		switch {
		case index == 0:
			arr[index] = 1
		case index == 1:
			arr[index] = 1
		case arr[index] == 0:
			arr[index] = arr[index-1] + arr[index-2]
		default:
		}
	}
	return arr
}
