package first

import (
	"bufio"
	"fmt"
	"os"
)

func dup(file *os.File, vis map[string]uint) {
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		vis[scanner.Text()]++
	}
	fmt.Println(vis)
}
