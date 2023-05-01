package main

import (
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

type Album struct {
	ID     int64
	Title  string
	Artist string
	Price  float32
}

func main() {

	file, err := os.Open("sqlPrep.txt")
	if err != nil {
		panic(err.Error())
	}

	defer file.Close()
	scanner := bufio.NewScanner(file)

	var isi [4]string
	i := 0
	for scanner.Scan() {
		// do something with a line
		isi[i] = scanner.Text()
		i++
	}

	fmt.Println(len(isi))

	db, err = sql.Open("mysql", isi[0]+":"+isi[1]+"@tcp(localhost:"+isi[2]+")/"+isi[3])
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rows, _ := db.Query("SELECT * FROM album ")
	defer rows.Close()


}
