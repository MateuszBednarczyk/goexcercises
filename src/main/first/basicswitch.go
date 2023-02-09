package first

import "fmt"

func ExampleSwtich(colour string) {
	switch colour {
	case "red":
		fmt.Println("Given colour is red")
	default:
		fmt.Println("Unknown colour")
	}
}
