package main

import (
	"database/sql"
	_ "database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	UserID       int
	UserName     string
	UserPassword string
	Email        string
}

func SelectUser(username string) string {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s"+" password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}
	defer db.Close()
	var myUser User
	sqlStatement := `
	SELECT user_id, username, password, email FROM users WHERE username = $1`

	err = db.QueryRow(sqlStatement, username).Scan(&myUser.UserID, &myUser.UserName, &myUser.UserPassword, &myUser.Email)

	if err != nil {
		log.Fatal("Failed to execute query: ", err)
	}

	return myUser.UserPassword
}

func Login(username, userpassword string) bool {
	hash := SelectUser(username)
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(userpassword))
	if err == nil {
		fmt.Println("Login realizado com sucesso!", hash)
	} else {
		fmt.Println("Falha ao realizar login!", err)
	}
	return err == nil
}
