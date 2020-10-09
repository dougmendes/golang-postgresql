package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"time"

	_ "github.com/lib/pq"
)

const (
	host     = "172.17.0.2"
	port     = 5432
	user     = "postgres"
	password = "87070540"
	dbname   = "postgres"
)

func main() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	sqlStatement := `
	INSERT INTO users(username, password, email, create_on, last_login)
	VALUES($1, $2, $3, $4, $5)
	RETURNING user_id`

	id := 0
	password := "87070540"
	hash, _ := HashPassword(password)
	err = db.QueryRow(sqlStatement, "maria", hash, "maria@gmail.com", time.Now(), time.Now()).Scan(&id)
	if err != nil {
		panic(err)
	}
	fmt.Println("New record ID is:")
}
