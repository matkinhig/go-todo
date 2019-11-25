package handler

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"os"
	"reflect"

	"github.com/labstack/echo"
	"github.com/matkinhig/go-todo/config"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK strong health")
}

func InitDb() *sql.DB {
	db, err := sql.Open("oci8", config.Config.OracleDB.Uri)
	if err != nil {
		log.Fatalln(err)
	}
	return db
}

func QuerryToJson(db *sql.DB, query string, args ...interface{}) ([]byte, error) {
	var objects []map[string]interface{}
	rows, err := db.Query(query, args...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		columns, err := rows.ColumnTypes()
		if err != nil {
			return nil, err
		}
		values := make([]interface{}, len(columns))
		object := map[string]interface{}{}
		for i, column := range columns {
			object[column.Name()] = reflect.New(column.ScanType()).Interface()
			values[i] = object[column.Name()]
		}
		err = rows.Scan(values...)
		if err != nil {
			return nil, err
		}
		objects = append(objects, object)
	}
	return json.MarshalIndent(objects, "", "\t")
}

func GetUserJson() {
	log.SetFlags(log.Lshortfile)
	db := InitDb
	b, err := QuerryToJson(db, "Select * from halong.user_account")
	if err != nil {
		log.Fatalln(err)
	}
	os.Stdout.Write(b)
}

// func GetAllUsersFromDB(c echo.Context) error {
// 	db, err := sql.Open("oci8", config.Config.OracleDB.Uri)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()

// 	rows, exception := db.Query("Select * from user_account")
// 	if exception != nil {
// 		log.Fatal(exception)
// 	}
// 	for rows.Next() {
// 		println(rows.Next())
// 	}
// 	defer rows.Close()
// }
