package mysql

import (
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type StudentModel struct {
	ID      string    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	ClassID string    `json:"class_id"`
}

const (
	ID      = "id"
	ClassID = "class_id"
	Code    = "code"
	Name    = "name"
)

func StudentQuery() *Query {
	return New("students", "id")
}

func SelectStudentId(ID int) (*StudentModel, error) {
	result, err := StudentQuery().Where("id", "=", strconv.Itoa(ID)).FirstResult()
	if err != nil {
		return nil, err
	}
	resultData := SetData(result, StudentModel{})
	var student *StudentModel
	err2 := mapstructure.Decode(resultData, &student)
	if err2 != nil {
		return student, err
	}
	return student, err
}

func SelectStudent() ([]*StudentModel, error) {
	var models []*StudentModel
	q := StudentQuery().Select("id", "name")
	results, err := q.Results()
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := SetData(r, StudentModel{})
		var student *StudentModel
		err2 := mapstructure.Decode(m, &student)
		if err2 == nil {
			models = append(models, student)
		}
	}
	for _, v := range models {
		fmt.Println("value :", v.Name)
	}
	return models, nil
}
func FindStudent() (*StudentModel, error) {
	q, _ := StudentQuery().Select("id", "name", "code", "class_id").FirstResult()

	m := SetData(q, StudentModel{})
	var student *StudentModel
	err2 := mapstructure.Decode(m, &student)

	if err2 != nil {
		fmt.Println("Co loi xay ra", err2.Error())
	}
	return student, nil
}
func SaveObject(model interface{}) (int64, error) {
	id, err := StudentQuery().SaveObject(model)
	return id, err
}
func UpdateStudent(student StudentModel) {
	params := map[string]interface{}{
		Code: student.Code,
		Name: student.Name,
	}
	err := StudentQuery().Where("id", "=", student.ID).Where("class_id", "=", student.ClassID).Update(params)
	if err != nil {
		fmt.Println("Co loi xay ra")
	}
}

func UpdateObject(student StudentModel)  {
	StudentQuery().Where("id", "=", student.ID).UpdateObject(student)
}

func AddStudent(param map[string]interface{}) (int64, error) {
	id, err := StudentQuery().Save(param)
	return id, err
}
