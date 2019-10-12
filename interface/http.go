package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	response, err := http.Get("https://farizal.dummy.heikaku.com/v1/category")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	fmt.Println(response)
}
