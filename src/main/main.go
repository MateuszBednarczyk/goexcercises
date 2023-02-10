package main

import (
	"fmt"

	first "goexcercises/src/main/first"
	second "goexcercises/src/main/second"
)

func main() {
	first.ExFirstOfFirst()
	first.ExSecondOfFirst()
	first.Fetch()
	fmt.Println(first.GetGrade(30, 30, 30))
	second.Convert()
	first.ExampleSwtich("red")
	first.ExampleSwtich("black")
	first.ExampleSwtich("red")
}
