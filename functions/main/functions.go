package main

import "fmt"

func main() {
	fmt.Println(greet1("Prashanth"))
	fmt.Println(greet2("Shruthi", "Shashank"))
	fmt.Println(average(11, 22.0, 24, 45, 21.3))
	// Variadic arguments
	data := []float64{11, 33.2, 2, 43.1, 4}
	fmt.Println(average(data...))
}

// Named return
func greet1(name string) (hello string) {
	hello = fmt.Sprint("Hello ", name)
	return
}

// Multiple return
func greet2(name1 string, name2 string) (string, string) {
	return fmt.Sprint("Hello ", name1), fmt.Sprint("Hello ", name2)
}

// Variadic inputs
func average(sf ...float64) float64 {
	total := 0.0
	for _, v := range sf {
		total += v
	}
	return total / float64(len(sf))
}
