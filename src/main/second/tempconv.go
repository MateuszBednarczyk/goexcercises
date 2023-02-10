package second

import "fmt"

const (
	cBoilign  = 100
	cFreezing = 0
)

func Convert() {
	var cTemp float64

	fmt.Println("Pass c temp to convert")
	fmt.Scan(&cTemp)

	converted := (cTemp * 1.8) + 32
	fmt.Println(converted)
}
