package mysql

import (
	"fmt"
)

type StudentModel struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	ClassID string `json:"class_id"`
}

const (
	StudentId = "id"
	ClassID   = "class_id"
	Code      = "code"
	Name      = "name"
)

func StudentQuery() *Query {
	return New("students", "id")
}

func StudentById(id string) (*StudentModel, error) {
	resultQuery, err := StudentQuery().Where("id", "=", id).FirstResult()
	if err != nil {
		return nil, err
	}
	rs := StudentQuery().SetData(resultQuery, *&StudentModel{})
	student := rs.(StudentModel)
	return &student, err
}

func ListStudent() ([]*StudentModel, error) {
	var models []*StudentModel
	q := StudentQuery().Select("id", "name")
	results, err := q.Results()
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		rs := StudentQuery().SetData(r, *&StudentModel{})
		student := rs.(StudentModel)
		models = append(models, &student)
	}

	return models, nil
}

func CreateStudent(model interface{}) (int64, error) {
	id, err := StudentQuery().InsertObject(model)
	return id, err
}

func UpdateStudent(student StudentModel) (int64, error) {
	rs, err := StudentQuery().Where("id", "=", student.ID).UpdateObject(student)
	if err != nil {
		fmt.Println("Co loi xay ra")
		return rs, err
	}
	if rs > 0 {
		fmt.Println("Update du lieu thanh cong")
	}
	return rs, err
}
