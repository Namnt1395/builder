package handle

import (
	"demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
	"strconv"
)

var tool mysql.Query

func TableQuery() string {
	return "students"
}
type Object struct {
	field map[string]string
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// get data input param
	queries := r.URL.Query()
	code := queries.Get("code")
	name := queries.Get("name")
	classId := queries.Get("class_id")

	params := map[string]string{
		"code":  code,
		"name":  name,
		"class_id": classId,

	}
	id, err := tool.Insert(params, TableQuery())
	if err != nil {
		fmt.Println("err" , err.Error())
	}
	if id > 0 {
		fmt.Println("Isert thanh cong")
	}
}


func SelectOneStudent(w http.ResponseWriter, r *http.Request)  {
	result, err := mysql.SelectStudentId(1)
	if err != nil {
		fmt.Println("lỗi xảy ra....", err.Error())
	}
	fmt.Println("-----------------\n")
	fmt.Println("HO ten :" +result.Name)
}
func SelectStudent(w http.ResponseWriter, r *http.Request)  {
	 mysql.SelectStudent()

}
func UpdateStudent(w http.ResponseWriter, r *http.Request)  {
	queries := r.URL.Query()
	idUpdate, _ := strconv.Atoi(queries.Get("id"))
	code := queries.Get("code")
	name := queries.Get("name")
	classId, _ := strconv.Atoi(queries.Get("class_id"))
	model := mysql.StudentModel{Code: code, Name: name}
	mysql.UpdateStudent(model, idUpdate, classId)
}
