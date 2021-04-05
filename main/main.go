package main

import (
	"database/sql"
	"net/http"

	_ "github.com/go-sql-driver/mysql"
)

func defaultHandler(w http.ResponseWriter, r *http.Request) {
	var user_name = getUserName()
	w.Write([]byte(user_name))
}

func main() {
	http.HandleFunc("/", defaultHandler)
	http.ListenAndServe(":8080", nil)
}

func getUserName() string {
	db, err := sql.Open("mysql", `root:root@tcp(127.0.0.1:3306)/test`)

	if err != nil || db.Ping() != nil {
		panic(err.Error())
	}
	defer db.Close()

	var user_name string
	selectErr := db.QueryRow("SELECT user_name FROM user WHERE user_id = 'test' ").Scan(&user_name)

	if selectErr != nil {
		panic(err.Error())
	}

	return user_name
}
