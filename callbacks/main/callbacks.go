package main

import "fmt"

func main() {
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}
