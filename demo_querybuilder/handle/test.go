package handle

import (
	"fmt"
	"reflect"
)

type Attributes struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
	Email    string `json:"email"`
}

type User struct {
	Attributes
	TableName  string
	PrimaryKey string
}

func (u *User) LoadDefaultData() {
	u.Attributes.Id = 0
	u.Attributes.Username = ""
	u.Attributes.Email = "admin@hotmail.com"
}

func SetData(data map[string]interface{}, object interface{}) map[string]interface{}{

	result := make(map[string]interface{})

	st := reflect.TypeOf(*&object)
	num := st.NumField()

	// for 1
	for i := 0; i < num; i++ {
		item := st.Field(i)
		//ps := reflect.ValueOf(u).Elem()

		// for in data
		for v, _ := range data {
			// check theo tag
			if item.Tag.Get("json") == v {
				// switch
				switch item.Type.Kind() {
				case reflect.Int:
					val, _ := data[v].(int)
					result[item.Name] = val
				case reflect.String:
					result[item.Name] = fmt.Sprintf("%v", data[v])
				default:
				} // end switch

			} // end if check name

		} // end for data

	} // end for 1
   return result
}


func (u *User) SetData12(data map[string]interface{}, object interface{}) {

	st := reflect.TypeOf(*&object)
	num := st.NumField()

	// for 1
	for i := 0; i < num; i++ {
		item := st.Field(i)
		ps := reflect.ValueOf(u).Elem()

		// for in data
		for v, _ := range data {

			// check theo tag
			if item.Tag.Get("json") == v {
				f := ps.FieldByName(item.Name)
				fmt.Println("item.Name......f", f)


				if !f.IsValid() || !f.CanSet() {
					continue
				}

				fmt.Println("kind....", f.Kind())
				// switch
				switch f.Kind() {
				case reflect.Int:
					val, _ := data[v].(int)
					f.SetInt(int64(val))
				case reflect.String:
					f.SetString(fmt.Sprintf("%v", data[v]))
				default:
				} // end switch

			} // end if check name

			// check theo key
			if item.Name == v {

			}

		} // end for data

	} // end for 1

}

func SetDataNew(data map[string]interface{}, object interface{}) {

	st := reflect.TypeOf(*&object)
	num := st.NumField()

	// for 1
	for i := 0; i < num; i++ {
		item := st.Field(i)
		ps := reflect.ValueOf(&object).Elem()
		fmt.Println("Item pssssssss....", ps)
		// for in data
		for v, _ := range data {
			fmt.Println("v data:", v)
			// check theo tag
			if item.Tag.Get("json") == v {
				f := ps.FieldByName(item.Name)
				fmt.Println("Item name...", item.Type)

				switch item.Type.Kind() {
				case reflect.Int:
					val, _ := data[v].(int)
					f.SetInt(int64(val))
				case reflect.String:
					f.SetString(fmt.Sprintf("%v", data[v]))
				default:
				}

				// switch
				//switch f.Kind() {
				//case reflect.Int:
				//	val, _ := data[v].(int)
				//	f.SetInt(int64(val))
				//case reflect.String:
				//	f.SetString(fmt.Sprintf("%v", data[v]))
				//default:
				//} // end switch

			} // end if check name

			// check theo key
			if item.Name == v {

			}

		} // end for data

	} // end for 1

}
