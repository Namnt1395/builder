package mysql

import (
	"fmt"
	"strconv"
)

type StudentModel struct {
	ID         int64
	ClassID    int
	Code       string
	Name       string
	OtherField map[string]string
}

const (
	ID      = "id"
	ClassID = "class_id"
	Code    = "code"
	Name    = "name"
)

// Query returns a new query for pages
func StudentQuery() *Query {
	return New("students", "id")
}

func GetValueStudent(cols map[string]interface{}) *StudentModel {
	student := &StudentModel{}

	student.ID = cols["id"].(int64)
	fmt.Println("cols ", student.ID)
	if cols["code"] != nil {
		student.Code = cols["code"].(string)
	}
	if cols["name"] != nil {
		student.Name = cols["name"].(string)
	}
	return student
}

func SelectStudentId(ID int) (*StudentModel, error) {
	result, err := StudentQuery().Where("id", "=", strconv.Itoa(ID)).FirstResult()
	if err != nil {
		return nil, err
	}
	return GetValueStudent(result), nil
}
func SelectStudent() ([]*StudentModel, error) {
	var models []*StudentModel
	q := StudentQuery().Select("id", "name").Where("id", ">", "1").
		Where("name", "=", "123").
		OrWhere("name", "=", "thanh nam")

	results, err := q.Results()

	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := GetValueStudent(r)
		models = append(models, m)
	}
	fmt.Println("model...", models)
	for _, v := range models {
		fmt.Println("value :", v.Name)
	}
	return models, nil
}

func SelectStudentWhereIn() ([]*StudentModel, error) {
	var models []*StudentModel
	var ids []int64
	ids = append(ids, 2)
	ids = append(ids, 4)
	ids = append(ids, 3)
	q := StudentQuery().Select("id", "name").WhereIn("id", ids)
	results, err := q.Results()

	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := GetValueStudent(r)
		models = append(models, m)
	}
	for _, v := range models {
		fmt.Println("value :", v.Name)
	}
	return models, nil
}

func SelectStudentWhereJoin() ([]*StudentModel, error) {
	var models []*StudentModel

	q := StudentQuery().Select("id", "name").Join("class", "class_id", "id")

	results, err := q.Results()

	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := GetValueStudent(r)
		models = append(models, m)
	}
	for _, v := range models {
		fmt.Println("value :", v.Name)
	}

	return models, nil

}
func UpdateStudent(student StudentModel) {
	params := map[string]string{
		Code: student.Code,
		Name: student.Name,
	}
	err := StudentQuery().Where("id", "=", student.ID).Where("class_id", "=", student.ClassID).Update(params)
	if err != nil {
		fmt.Println("Co loi xay ra")
	}
}
func AddStudent(param map[string]string) (int64, error) {
	id, err := StudentQuery().Insert(param)

	return id, err
}
//func UpdateData(id int)  {
//	result, _ :=StudentQuery().Select("id", "code", "name").Where("id", "=", id).Result()
//
//}
