package main

import "fmt"

func main() {
	visit([]int{1, 2, 3, 4}, func(i int) {
		fmt.Println(i)
	})

	// Filtering even numbers from a list
	fmt.Println(filter([]int{22, 43, 53, 67, 88, 42, 76}, filterEven))
}

func visit(numbers []int, callback func(int)) {
	for _, n := range numbers {
		callback(n)
	}
}

func filter(numbers []int, callback func(i int) bool) (filteredList []int) {
	for _, value := range numbers {
		if callback(value) {
			filteredList = append(filteredList, value)
		}
	}
	return
}

func filterEven(number int) bool {
	return number%2 == 0
}
