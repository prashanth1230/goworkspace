package main

import "fmt"

func alterValue(z *int) {
	*z = 10000
}

func main() {
	x := 20
	fmt.Println("Value of a and memory location of a are: ", x, &x)
	fmt.Scan(&x)
	fmt.Println("Updated Value of a and memory location of a are: ", x, &x)

	a := 300
	fmt.Println("Value and address: ", a, &a)
	var b *int = &a
	fmt.Println("Address, value and the pointer value of b: ", b, &b, *b)
	// Changing the value
	*b = 400
	fmt.Println("Value and address: ", a, &a)
	fmt.Println("Address, value and the pointer value of b: ", b, &b, *b)

	fmt.Println("x before altering ", x)
	alterValue(&x)
	fmt.Println("x after altering ", x)
}
