// package main

// import (
// 	"fmt"
// 	"net/http"
// )

// func main() {
// 	// The fake API
// 	url := "https://jsonplaceholder.typicode.com/posts/1"

// 	// Create a DELETE request. No request body is needed here.
// 	request, err := http.NewRequest(http.MethodDelete, url, nil)
// 	if err != nil {
// 		fmt.Println("Error creating request:", err)
// 		return
// 	}

// 	// Send the request to the API.
// 	response, err := http.DefaultClient.Do(request)
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer response.Body.Close()

// 	// Print the status returned by the API.
// 	fmt.Println("Status:", response.Status)
// }

package main

import (
	"fmt"
	"net/http"
)

func main() {
	//the fake uri'
	url := "https://jsonplaceholder.typicode.com/posts/1"

	//now creating a delete request

	request, err := http.NewRequest(http.MethodDelete, url, nil)
	if err != nil {
		fmt.Println("Error while creating a request ", err)
		return
	}
	// send the request to the api
	response, err := http.DefaultClient.Do(request)
	if err != nil {
		fmt.Println("Error while sending the request ", err)
		return
	}
	defer response.Body.Close()

	// status return by the api
	fmt.Println("status", response.Status)
}
