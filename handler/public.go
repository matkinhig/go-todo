package handler

import (
	"net/http"

	"github.com/labstack/echo"
)

func CheckHealth(c echo.Context) error {
	return c.String(http.StatusOK, "OK strong health")
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
