package main

import "fmt"

func main() {
	fmt.Println(greet1("Prashanth"))
	fmt.Println(greet2("Shruthi", "Shashank"))
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
