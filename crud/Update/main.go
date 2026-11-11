package main

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

func main() {
	// Update the post with ID 1.
	url := "https://jsonplaceholder.typicode.com/posts/1"

	// The updated post data in JSON format.
	data := `{"id":1,"title":"Updated title","body":"Updated message","userId":1}`

	// Create a PUT request with the new data.
	request, err := http.NewRequest(http.MethodPut, url, strings.NewReader(data))
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	// Tell the API that we're sending JSON.
	request.Header.Set("Content-Type", "application/json")

	// Send the request.
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error sending request:", err)
		return
	}
	defer response.Body.Close()

	// Read and print the response.
	result, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Println("Error reading response:", err)
		return
	}

	fmt.Println("Status:", response.Status)
	fmt.Println("Response:", string(result))
}
