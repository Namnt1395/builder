package main

import (
	"demo_querybuilder/handle"
	"demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
)

func main() {
	mysql.OpenDatabase()
	http.HandleFunc("/add", handle.AddStudent)
	http.HandleFunc("/update", handle.UpdateStudent)
	http.HandleFunc("/select", handle.SelectStudent)

	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("START")
}
