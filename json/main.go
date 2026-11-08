package main

import (
	"encoding/json"
	"fmt"
)

// object
type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	IsAdult bool   `json:"is_adult"`
}

func main() {
	fmt.Println("we are learnig json")
	person1 := Person{Name: "Jasmin", Age: 31, IsAdult: true}
	//fmt.Println("Person1 details", person1)

	// convert object into json string // marshaling
	jsonData, err := json.Marshal(person1)
	if err != nil {
		fmt.Println("Error while marshaling(convering to json string )", err)
	}
	fmt.Println("Person1 deatils in json format", string(jsonData))

	// convert json string into object // unmarshaling
	var personData Person
	err = json.Unmarshal(jsonData, &personData)
	if err != nil {
		fmt.Println("Error while unmarshaling", err)
	}
	fmt.Println("person data after unmarshaling", personData)

}
