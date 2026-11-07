package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func main() {

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error while creating the file:", err)
		return
	}
	defer file.Close()

	content := " hello world by P Thakre"
	// byte, err := io.WriteString(file, contend+"\n")
	// fmt.Println("Byte written", byte)
	// if err != nil {
	// 	fmt.Println("Error while writing a file", err)
	// 	return
	// }

	// fmt.Println("Successfully crated a file")

	// file read using io.utils : Read entire file into byte
	content, err := ioutil.ReadFile("example.txt")
	if err != nil {
		fmt.Println("Error while reading the file", err)
		return
	}
	fmt.Println(string(content))
}
