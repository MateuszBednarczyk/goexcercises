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
	first.ExampleSwtich("red")
	first.ExampleSwtich("black")
	fmt.Println(first.GetGrade(30, 30, 30))
	second.Convert()
}
