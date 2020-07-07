package main

import (
	"builder/demo_querybuilder/handle"
	"builder/demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
)

func main() {

	mysql.Connect()

	http.HandleFunc("/create_student", handle.CreateStudent)
	http.HandleFunc("/student", handle.StudentById)
	http.HandleFunc("/update", handle.UpdateStudent)
	http.HandleFunc("/list_student", handle.ListStudent)

	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("START")
}
