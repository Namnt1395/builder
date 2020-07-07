package adapters

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type MysqlAdapter struct {
	options map[string]string
	sqlDB   *sql.DB
	debug   bool
}

// Open this database
func (db *MysqlAdapter) Open() error {
	db.debug = false
	db.options = map[string]string{
		"adapter":  "mysql",
		"user":     "namnt", // sub your user
		"password": "123456",
		"db":       "bg_student",
		"protocol": "tcp",
		"host":     "localhost",
		"port":     "3306",
		"params":   "charset=utf8&parseTime=true",
	}

	//"user:password@tcp(localhost:3306)/dbname?charset=utf8&parseTime=true")
	options := fmt.Sprintf("%s:%s@%s(%s:%s)/%s?%s",
		db.options["user"],
		db.options["password"],
		db.options["protocol"],
		db.options["host"],
		db.options["port"],
		db.options["db"],
		db.options["params"])

	var err error
	db.sqlDB, err = sql.Open(db.options["adapter"], options)
	if err != nil {
		return err
	}

	if db.sqlDB == nil {
		fmt.Printf("Mysql options:%s", options)
		return fmt.Errorf("\nError creating database with options: %v", db.options)
	}

	err = db.sqlDB.Ping()
	if err != nil {
		return err
	}
	return err

}

// Close the database
func (db *MysqlAdapter) Close() error {
	if db.sqlDB != nil {
		return db.sqlDB.Close()
	}
	return nil
}

// Query SQL execute - NB caller must call use defer rows.Close() with rows returned
func (db *MysqlAdapter) Query(query string, args ...interface{}) (*sql.Rows, error) {
	if db.sqlDB == nil {
		return nil, fmt.Errorf("No database available.")
	}

	if db.debug {
		fmt.Println("QUERY:", query, "ARGS", args)
	}
	stmt, err := db.sqlDB.Prepare(query)
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
func (db *MysqlAdapter) Exec(query string, args ...interface{}) (sql.Result, error) {
	if db.sqlDB == nil {
		return nil, fmt.Errorf("No database available.")
	}
	if db.debug {
		fmt.Println("QUERY:", query, "ARGS", args)
	}

	stmt, err := db.sqlDB.Prepare(query)
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
func (db *MysqlAdapter) QuoteField(name string) string {
	return fmt.Sprintf("`%s`", name)
}

func (db *MysqlAdapter) Insert(query string, args ...interface{}) (id int64, err error) {

	tx, err := db.sqlDB.Begin()
	if err != nil {
		return 0, err
	}

	// Execute the sql using db
	result, err := db.Exec(query, args...)
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
func (db *MysqlAdapter) ReplaceArgPlaceholder(sql string, args []interface{}) string {
	return sql
}
func (db *MysqlAdapter) Placeholder(i int) string {
	return "?"
}
