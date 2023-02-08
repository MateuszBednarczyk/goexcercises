package first

import (
	"fmt"
	"os"
	"strings"
)

func ExFirstOfFirst() {
	fmt.Println(strings.Join(os.Args[0:], " "))
}
