package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Todo struct {
	UserID    int    `json:"userID"`
	ID        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Learning a CRUD operation in go")

	res, err := http.Get("https://dummyjson.com/products/1")
	if err != nil {
		fmt.Println("Error while getting data from the api", err)
		return
	}
	defer res.Body.Close() // close the connection after reading the data resource

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error while getting the response", res.StatusCode)
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while reading data from the api", err)
	// 	return
	// }
	// fmt.Println("Data", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error while decoding the JSON data", err)
	}

	fmt.Println("Todo data", todo)
	fmt.Println("Todo title", todo.Title)
	fmt.Println("Todo completed", todo.Completed)
}
