package main

import "fmt"

func main() {

	//use a map
	studentGrades := make(map[string]int)

	studentGrades["Price"] = 100
	studentGrades["Ashu"] = 98
	studentGrades["Suhash"] = 102

	fmt.Println("Mark of that student", studentGrades["Ashu"])
	studentGrades["Ashu"] = 99
	fmt.Println("New mark of that student", studentGrades["Ashu"])

	// to delete the mark of that student
	delete(studentGrades, "Bob")
	fmt.Println("Deleted mark of student Bob", studentGrades["Bob"])

	student := map[string]int{
		"Alice": 90,
	}
	fmt.Println("Mark of student Alice is", student["Alice"])
}
