package handle

import (
	"builder/demo_querybuilder/entity"
	"builder/demo_querybuilder/model/mysql"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"net/http"
)

func AddStudentV2(w http.ResponseWriter, r *http.Request) {
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
	id, err := entity.StudentQuery().Save(params)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	if id > 0 {
		fmt.Println("Insert thanh cong")
	}
}
func InsertObjectV2(w http.ResponseWriter, r *http.Request) {
	model := &entity.StudentEntity{Code: "1256", Name: "Namnt455", ClassID: "2"}
	id, err := entity.StudentQuery().SaveObject(model)
	if err != nil {
		fmt.Println("err", err.Error())
	}
	if id > 0 {
		fmt.Println("Insert thanh cong")
	}
}

func SelectOneStudentV2(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	id := queries.Get("code")

	result, err := entity.StudentQuery().Where("id", "=", id).FirstResult()
	if err != nil {
		return
	}
	resultData := mysql.SetData(result, entity.StudentEntity{})
	var student *entity.StudentEntity
	err2 := mapstructure.Decode(resultData, &student)
	if err2 != nil {
		fmt.Println("lỗi xảy ra....", err.Error())
		return
	}
	fmt.Println("-----------------\n")
	fmt.Println("HO ten :" + student.Name)
}
func SelectStudentV2(w http.ResponseWriter, r *http.Request) {
	mysql.SelectStudent()
	var models []*entity.StudentEntity
	q := entity.StudentQuery().Select("id", "name")
	results, err := q.Results()
	if err != nil {
		fmt.Println("co loi xay ra")
	}
	for _, r := range results {
		m := mysql.SetData(r, entity.StudentEntity{})
		var student *entity.StudentEntity
		err2 := mapstructure.Decode(m, &student)
		if err2 == nil {
			models = append(models, student)
		}
	}
	for _, v := range models {
		fmt.Println("value :", v.Name)
	}
	//mysql.SelectStudentWhereIn()
	//mysql.SelectStudentWhereJoin()
}
func UpdateStudentV2(w http.ResponseWriter, r *http.Request) {
	queries := r.URL.Query()
	idUpdate := queries.Get("id")
	code := queries.Get("code")
	name := queries.Get("name")
	classId := queries.Get("class_id")
	model := entity.StudentEntity{
		ID:      idUpdate,
		Code:    code,
		Name:    name,
		ClassID: classId,
	}
	err := entity.StudentQuery().Where("id", "=", idUpdate).Where("class_id", "=", classId).UpdateObject(model)
	if err != nil {
		fmt.Println("Co loi xay ra")
	}
}
func UpdateObjectStudentV2(w http.ResponseWriter, r *http.Request) {
	q, _ := entity.StudentQuery().Select("id", "name", "code", "class_id").FirstResult()
	m := mysql.SetData(q, entity.StudentEntity{})
	var student *entity.StudentEntity
	err2 := mapstructure.Decode(m, &student)
	if err2 != nil {
		fmt.Println("Co loi xay ra", err2.Error())
	}
	rs := &entity.StudentEntity{
		ID:      student.ID,
		Code:    student.Code,
		Name:    "Tuấn 123",
		ClassID: student.ClassID,
	}
	entity.StudentQuery().Where("id", "=", rs.ID).UpdateObject(rs)
}
