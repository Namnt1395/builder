package handle

import (
	"builder/demo_querybuilder/model/mysql"
	"fmt"
	"net/http"
)

func CreateStudent(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	code := queries.Get("code")
	name := queries.Get("name")
	classId := queries.Get("class_id")
	p := &mysql.StudentModel{Code: code, Name: name, ClassID: classId}
	rs, _ := mysql.CreateStudent(p)
	if rs > 0 {
		fmt.Println("Them thanh cong")
	}
}

func StudentById(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	idStudent := queries.Get("id")
	result, err := mysql.StudentById(idStudent)
	if err != nil {
		fmt.Println("lỗi xảy ra....", err.Error())
	}
	fmt.Println("-----------------\n", result)
	if result != nil {
		fmt.Println("HO ten :" + result.Name)
	}

}
func StudentWithClass() {
	mysql.StudentWithClass()
}
func ListStudent(w http.ResponseWriter, r *http.Request) {
	rs, _ := mysql.ListStudent()
	if len(rs) <= 0 {
		fmt.Println("Khong co du lieu")
	}
	for _, studentInfo := range rs {
		fmt.Println("value :", studentInfo.Name)
	}
}
