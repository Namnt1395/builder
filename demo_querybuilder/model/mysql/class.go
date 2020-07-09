package mysql

type Class struct {
	ID        int    `builder:"id"`
	ClassName string `builder:"class_name"`
	TestName  string `builder:"test_name"`
}
