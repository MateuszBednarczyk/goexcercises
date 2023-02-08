package first

import (
	"fmt"
	"io"
	"log"
	"net/http"
)

const (
	prefix = "http://"
)

func Fetch() {
	fmt.Println("Pass url")
	var url string
	fmt.Scan(&url)
	url = prefix + url

	res, err := http.Get(url)
	if err != nil {
		log.Fatal("Can't fetch")
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal("Error while reading")
	}
	fmt.Printf("%s", body)
	fmt.Println("\n", "status code:", res.StatusCode)
}
