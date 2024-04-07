package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	url := "https://api.github.com"
	client := http.Client{}

	_, _ = client, url

	response, err := client.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Println(response.StatusCode)

	bytes, err := io.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(bytes))

}
