package main

import (
    "database/sql"

    _ "github.com/goinaction/code/chapter3/dbdriver/postgres"
)

func main() {
    sql.Open("postgres", "mydb")
}
