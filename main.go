package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/matkinhig/go-todo/config"
	_ "github.com/mattn/go-oci8"
	_ "gopkg.in/goracle.v2"
)

func main() {
	fmt.Println("start golang")
	fmt.Println(config.Config)

	fmt.Println(config.Config.OracleDB.Uri)
	GetAllUsersFromDB()
}

func buildServer() {
	e := echo.New()
	err := e.Start(":1234")
	if err != nil {
		fmt.Println("Cannot start server", err)
	}

}

func testConnection() {
	db, err := sql.Open("oci8", config.Config.OracleDB.Uri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	println("Connection succcess!!")
	rows, err := db.Query("SELECT sysdate  FROM dual")
	if err != nil {
		log.Fatalln("err:", err.Error)
	}
	defer rows.Close()
	var (
		sysdate string
	)
	for rows.Next() {
		if err = rows.Scan(&sysdate); err != nil {
			log.Fatalln("error fetching", err)
		}
		log.Println(sysdate)
	}
}

func GetAllUsersFromDB() {
	db, err := sql.Open("oci8", config.Config.OracleDB.Uri)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	rows, exception := db.Query("Select * from user_account")
	if exception != nil {
		log.Fatal(exception)
	}
	for rows.Next() {

	}
	defer rows.Close()
}
