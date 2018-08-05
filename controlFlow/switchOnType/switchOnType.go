package switchOnType

import "fmt"

func SwitchOnType(input interface{}) {
	switch input.(type) {
	case int:
		fmt.Println("Its an integer")
	case string:
		fmt.Println("Its a string")
	case bool:
		fmt.Println("Its a boolean")
	default:
		fmt.Println("Not sure what type")
	}
}
