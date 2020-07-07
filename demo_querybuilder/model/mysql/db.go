package mysql

import (
	"database/sql"
	"fmt"
)

var debug bool

func init() {

}

// Query SQL execute - NB caller must call use defer rows.Close() with rows returned
func QuerySql(query string, args ...interface{}) (*sql.Rows, error) {
	if MySQL == nil {
		return nil, fmt.Errorf("No database available.")
	}

	if debug {
		fmt.Println("QUERY:", query, "ARGS", args)
	}
	stmt, err := MySQL.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query(args...)

	if err != nil {
		return nil, err
	}

	// Caller is responsible for closing rows with defer rows.Close()
	return rows, err
}

// Exec - use this for non-select statements
func Exec(query string, args ...interface{}) (sql.Result, error) {
	if MySQL == nil {
		return nil, fmt.Errorf("No database available.")
	}
	if debug {
		fmt.Println("QUERY:", query, "ARGS", args)
	}

	stmt, err := MySQL.Prepare(query)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)

	if err != nil {
		return result, err
	}
	return result, err
}

// QuoteField quotes a table name or column name
func QuoteField(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func Insert(query string, args ...interface{}) (id int64, err error) {

	tx, err := MySQL.Begin()
	if err != nil {
		return 0, err
	}

	// Execute the sql using db
	result, err := Exec(query, args...)
	if err != nil {
		return 0, err
	}

	id, err = result.LastInsertId()
	if err != nil {
		return 0, err
	}

	err = tx.Commit()
	if err != nil {
		return 0, err
	}

	return id, nil

}
func ReplaceArgPlaceholder(sql string, args []interface{}) string {
	return sql
}
func Placeholder(i int) string {
	return "?"
}
