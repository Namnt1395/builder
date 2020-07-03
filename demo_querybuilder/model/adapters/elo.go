package adapters

type User struct{
	Attributes

	TableName string
	PrimaryKey string

}

type Attributes struct {
	Name string
	Age int
}

