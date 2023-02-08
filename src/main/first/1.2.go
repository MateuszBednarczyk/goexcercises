package first

import (
	"fmt"
	"os"
)

func ExSecondOfFirst() {
	args := os.Args

	for index, arg := range args {
		fmt.Println(index, " ", arg)
	}
}
