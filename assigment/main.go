package main

import "fmt"

type numbers []int

func main() {
	numbers := numbers{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	numbers.printData()
}

func (n numbers) printData() {
	for _, number := range n {
		if (number % 2) == 0 {
			fmt.Println(number, "is even")
		} else {
			fmt.Println(number, "is odd")
		}
	}
}
