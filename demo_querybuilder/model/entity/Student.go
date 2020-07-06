package entity

type Student struct {
	ID      int    `json:"id"`
	Code    string `json:"code"`
	Name    string `json:"name"`
	ClassID int    `json:"class_id"`
}
