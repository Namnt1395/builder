package entity

import "builder/demo_querybuilder/model/mysql"

type StudentEntity struct {
	ID      string `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	ClassID string `json:"class_id"`
}

func StudentQuery() *mysql.Query {
	return mysql.New("students", "id")
}
