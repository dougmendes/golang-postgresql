package main

import (
	_ "database/sql"

	_ "github.com/lib/pq"
)

func main() {

	InsertUser("maria", "maria110394", "maria@gmail.com")
	Login("maria", "maria110394")

}
