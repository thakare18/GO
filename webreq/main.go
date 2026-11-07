package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Learing web request")
	res, err := http.Get("https://jsonplaceholder.typicode.com/posts/10")
	if err != nil {
		fmt.Println("Error while getting a http response")
		return
	}
	// to free resources after geting
	defer res.Body.Close()
	fmt.Printf("Type of response: %T", res)

	// to read a response body

	data, err := ioutil.ReadAll(res.Body)
	//error handling
	if err != nil {
		fmt.Println("Error while reading a response", res)
		return
	}
	fmt.Println("A response data", string(data))

}
