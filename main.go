package main

import (
	"database/sql"
	"fmt"

	"github.com/labstack/echo"
	"github.com/matkinhig/go-todo/config"
	"gopkg.in/goracle.v2"
)

func main() {
	fmt.Println("start golang")
	fmt.Println(config.Config)

	testConnection()
}

func buildServer() {
	e := echo.New()
	err := e.Start(":1234")

	if err != nil {
		fmt.Println("Cannot start server", err)
	}
}

func testConnection() {
	db, err := sql.Open("goracle", "scott/tiger@10.0.1.127:1521/orclpdb1")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer db.Close()

	rows, err := db.Query("select sysdate from dual")
	if err != nil {
		fmt.Println("Error running query")
		fmt.Println(err)
		return
	}
	defer rows.Close()

	var thedate string
	for rows.Next() {

		rows.Scan(&thedate)
	}
	fmt.Printf("The date is: %s\n", thedate)
}
