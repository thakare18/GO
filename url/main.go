package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Learning URL")
	myURL := "https://jsonplaceholder.typicode.com/posts/1"
	fmt.Printf("Type of url %T\n", myURL)
	parseURL, err := url.Parse(myURL)
	if err != nil {
		fmt.Println("Erorr while parsing a URL", err)
		return
	}
	fmt.Printf("Type of url: %T\n", parseURL)

	fmt.Println("Scheme", parseURL.Scheme)
	fmt.Println("Host", parseURL.Host)
	fmt.Println("Path", parseURL.Path)
	fmt.Println("RawQuery", parseURL.RawQuery)

	//modifying URL component
	parseURL.Path = "/newPost/2"
	parseURL.Scheme = "http"

	//// constructing a url string from url object
	newUrl := parseURL.String()
	fmt.Println("New url:", newUrl)

}
