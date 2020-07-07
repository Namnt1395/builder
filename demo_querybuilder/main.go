package main

import (
	"builder/demo_querybuilder/handle"
	"builder/demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
)

func main() {
	_ = mysql.OpenDatabase()

	http.HandleFunc("/ins", handle.InsertObject)
	http.HandleFunc("/add", handle.AddStudent)
	http.HandleFunc("/update", handle.UpdateStudent)
	http.HandleFunc("/updobj", handle.UpdateObjectStudent)
	http.HandleFunc("/select", handle.SelectStudent)
	http.HandleFunc("/selectone", handle.SelectOneStudent)
	//
	err := http.ListenAndServe(":9080", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("START")
}
