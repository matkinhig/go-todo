package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/labstack/echo"
	"github.com/matkinhig/go-todo/config"
	"github.com/matkinhig/go-todo/handler"
	"github.com/matkinhig/go-todo/types"
	_ "github.com/mattn/go-oci8"
	_ "gopkg.in/goracle.v2"
)

func main() {
	fmt.Println("start golang")
	fmt.Println(config.Config)

	fmt.Println(config.Config.OracleDB.Uri)
	testConnection()
	GetAllUsersFromDB()
	handler.GetUserJson()
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
	rows, err := db.Query("Select * from halong.user_account")
	s := []types.UserAccount{}
	if err != nil {
		fmt.Println(err)
		panic("cant connect to oracle")
	}
	for rows.Next() {
		var us types.UserAccount
		err = rows.Scan(&us.USER_NAME, &us.PASSWORD, &us.EMAIL, &us.AGE)
		if err != nil {
			log.Fatal(err)
			panic(err)
		}
		s = append(s, us)
	}

	fmt.Println(s[1].USER_NAME)
}
