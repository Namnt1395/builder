package mysql

import (
	"demo_querybuilder/model/entity"
	"fmt"
	"github.com/mitchellh/mapstructure"
	"strconv"
)

type StudentModel struct {
	ID      int64
	ClassID int
	Code    string
	Name    string
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

func SelectStudentId(ID int) (*entity.Student, error) {
	result, err := StudentQuery().Where("id", "=", strconv.Itoa(ID)).FirstResult()
	if err != nil {
		return nil, err
	}
	resultData := SetData(result, entity.Student{})
	var student *entity.Student
	err2 := mapstructure.Decode(resultData, &student)
	if err2 != nil {
		return student, err
	}
	return student, err
}

func SelectStudent() ([]*entity.Student, error) {
	var models []*entity.Student
	q := StudentQuery().Select("id", "name")
	results, err := q.Results()
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		m := SetData(r, entity.Student{})
		var student *entity.Student
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
func FindStudent() (*entity.Student, error) {
	q, _ := StudentQuery().Select("id", "name", "code", "class_id").FirstResult()

	m := SetData(q, entity.Student{})
	var student *entity.Student
	err2 := mapstructure.Decode(m, &student)

	if err2 != nil {
		fmt.Println("Co loi xay ra", err2.Error())
	}

	fmt.Println("student", student)
	return student, nil
}

//func SelectStudentWhereIn() ([]*StudentModel, error) {
//	var models []*StudentModel
//	var ids []int64
//	ids = append(ids, 2)
//	ids = append(ids, 4)
//	ids = append(ids, 3)
//	q := StudentQuery().Select("id", "name").WhereIn("id", ids)
//	results, err := q.Results()
//
//	if err != nil {
//		return nil, err
//	}
//	for _, r := range results {
//		m := GetValueStudent(r)
//		models = append(models, m)
//	}
//	for _, v := range models {
//		fmt.Println("value :", v.Name)
//	}
//	return models, nil
//}
//
//func SelectStudentWhereJoin() ([]*StudentModel, error) {
//	var models []*StudentModel
//
//	q := StudentQuery().Select("id", "name").Join("class", "class_id", "id")
//
//	results, err := q.Results()
//
//	if err != nil {
//		return nil, err
//	}
//	for _, r := range results {
//		m := GetValueStudent(r)
//		models = append(models, m)
//	}
//	for _, v := range models {
//		fmt.Println("value :", v.Name)
//	}
//
//	return models, nil
//
//}

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

func UpdateObject(student entity.Student)  {
	fmt.Println("update value",student.Name)
	StudentQuery().Where("id", "=", strconv.Itoa(student.ID)).UpdateObject(student)
}

func AddStudent(param map[string]interface{}) (int64, error) {
	fmt.Println(param)
	id, err := StudentQuery().Insert(param)

	return id, err
}
