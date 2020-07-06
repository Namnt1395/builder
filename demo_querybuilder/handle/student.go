package handle

import (
	"demo_querybuilder/model/entity"
	"demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
	"strconv"
)

var tool mysql.Query

type Object struct {
	field map[string]string
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// get data input param
	queries := r.URL.Query()
	code := queries.Get("code")
	name := queries.Get("name")
	classId := queries.Get("class_id")

	params := map[string]interface{}{
		"code":     code,
		"name":     name,
		"class_id": classId,
	}
	id, err := mysql.AddStudent(params)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	if id > 0 {
		fmt.Println("Isert thanh cong")
	}
}
func InsertObject(w http.ResponseWriter, r *http.Request) {
	//p := entity.Student{Code: "1256", Name: "Namnt455", ClassID: 2}
	//mysql.SaveObject(p)
	//fmt.Println(err.Error())
}

func SelectOneStudent(w http.ResponseWriter, r *http.Request) {
	result, err := mysql.SelectStudentId(1)
	if err != nil {
		fmt.Println("lỗi xảy ra....", err.Error())
	}
	fmt.Println("-----------------\n")
	fmt.Println("HO ten :" + result.Name)
}
func SelectStudent(w http.ResponseWriter, r *http.Request) {
	mysql.SelectStudent()
	//mysql.SelectStudentWhereIn()
	//mysql.SelectStudentWhereJoin()
}
func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	idUpdate, _ := strconv.Atoi(queries.Get("id"))
	code := queries.Get("code")
	name := queries.Get("name")
	classId, _ := strconv.Atoi(queries.Get("class_id"))
	model := mysql.StudentModel{ID: int64(idUpdate), ClassID: classId, Code: code, Name: name}
	mysql.UpdateStudent(model)
}
func UpdateObjectStudent(w http.ResponseWriter, r *http.Request) {
	rs, _ := mysql.FindStudent()
	rs = &entity.Student{
		ID: rs.ID,
		Code:    rs.Code,
		Name:    "Tuấn 123",
		ClassID: rs.ClassID,
	}

	mysql.UpdateObject(*rs)
}
