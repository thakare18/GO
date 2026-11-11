package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// The API address
	url := "https://jsonplaceholder.typicode.com/posts"

	// The data we want to send
	data := `{"title":"Hello","body":"My first post","userId":1}`

	// Send the data using POST
	response, err := http.Post(url, "application/json", strings.NewReader(data))
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer response.Body.Close()

	// Read and print the API's response
	result, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Status:", response.Status)
	fmt.Println("Response:", string(result))
}
