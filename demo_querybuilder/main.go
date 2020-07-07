package main

import (
	"builder/demo_querybuilder/handle"
	"builder/demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
)

func main() {

	//params := map[string]interface{}{
	//	"id" : "123",
	//	"username" : "namnt",
	//	"email" : "nam.nt@bidgear.com",
	//}

	//q := &handle.User{
	//	Attributes: handle.Attributes{},
	//	TableName:  "",
	//	PrimaryKey: "",
	//}
	//result := handle.SetData(params, handle.Attributes{})
	//
	//var user handle.Attributes
	//err := mapstructure.Decode(result, &user)
	//if err != nil {
	//	// error
	//}
	//fmt.Print("Email...", user.Email)

	//handle.SetDataNew(params, handle.Attributes{})
	//fmt.Println("email....", result["Email"])

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
