package mysql

import (
	"fmt"
)

type StudentModel struct {
	ID         int
	ClassID    int
	Code       string
	Name       string
	OtherField map[string]string
}

// Query returns a new query for pages
func StudentQuery() *Query {
	return New("students", "id")
}

func GetValueStudent(cols map[string]interface{}) *StudentModel {
	student := &StudentModel{}
	student.ID, _ = cols["id"].(int)
	if cols["code"] != nil {
		student.Code = cols["code"].(string)
	}
	if cols["name"] != nil {
		student.Name = cols["name"].(string)
	}
	return student
}

func SelectStudentId(ID int64) (*StudentModel, error) {
	result, err := StudentQuery().Where("id=?", ID).FirstResult()
	if err != nil {
		return nil, err
	}
	return GetValueStudent(result), nil
}
func SelectStudent() ([]*StudentModel, error) {
	var models []*StudentModel
	q := StudentQuery().Select("SELECT students.id,students.name, class.name_class from students").Join("class", "class_id", "id")
	results, err := q.Results()
	fmt.Println("result :" , q.sql)
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := GetValueStudent(r)
		models = append(models, m)
	}
	for _, v := range models {
		fmt.Println("value :", v)
	}
	return models, nil
}
func UpdateStudent(student StudentModel, idUpdate int, classId int) {
	params := map[string]string{
		"code": student.Code,
		"name": student.Name,
	}
	err := StudentQuery().Where("id=?", idUpdate).Where("class_id=?", classId).Update(params)
	if err != nil {
		fmt.Println("Co loi xay ra")
	}
}
