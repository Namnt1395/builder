package mysql

import "fmt"

var myClass *Class

type StudentModel struct {
	ID           string                 `json:"id" builder:"id"`
	Code         string                 `json:"code" builder:"code"`
	Name         string                 `json:"name" builder:"name"`
	ClassID      string                 `json:"class_id" builder:"class_id"`
	Relationship map[string]interface{} `builder:"rel"`
}

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

func StudentWithClass() ([]*StudentModel, error) {
	var models []*StudentModel
	q := StudentQuery().Select("id", "class_id", "code", "name", "c.class_name").InnerJoin("class c", "students.class_id=c.id")

	results, err := q.Results()
	if err != nil {
		return nil, err
	}
	for _, r := range results {
		rs := StudentQuery().SetData(r, *&StudentModel{})
		student := rs.(StudentModel)
		models = append(models, &student)
	}
	for _, v := range models {
		fmt.Println("value...", v.Code)
	}
	return models, nil
}
