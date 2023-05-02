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

	// main calculator
	// 	scanner := bufio.NewScanner(os.Stdin)
	// 	scanner.Scan()
	// 	text := scanner.Text()
	// 	var regex, err = regexp.Compile(`[-+]?[0-9]*\.?[0-9]+([-+*/]?([0-9]*\.?[0-9]+))*`)
	// 	if err != nil {
	// 		fmt.Println(err.Error())
	// 	}
	// 	var hasil = regex.FindAllString(text, -1)
	// 	if len(hasil) == 0 {
	// 		fmt.Println("Sintaks persamaan tidak sesuai")
	// 	} else {
	// 		fmt.Println(hasil[0])
	// 		lib.Calculator(hasil[0])
	// 	}

	// main kmp
	// text := "ABCDEFGABCDFDF"
	// pat := "BC"
	// fmt.Println("Text: ", text)
	// fmt.Println("Pattern: ", pat)
	// temp := (lib.KmpMatch(text, pat))
	// fmt.Println("Index: ", temp)
	// if temp == -1 {
	// 	fmt.Println("Not Matched")
	// } else {
	// 	fmt.Println("Matched")
	// }

	// main kalender
	// scanner := bufio.NewScanner(os.Stdin)
	// scanner.Scan()
	// text := scanner.Text()
	// var regex, err = regexp.Compile(`[0-9]{1,2}/[0-9]{1,2}/[0-9]{1,4}`)
	// if err != nil {
	// 	fmt.Println(err.Error())
	// }
	// var hasil = regex.FindAllString(text, -1)
	// if lib.IsCalValid(hasil[0]) {
	// 	fmt.Println(lib.GetDay(hasil[0]))
	// } else {
	// 	fmt.Println("Invalid Date")
	// }

	// main bm
	// fmt.Println("crazy brown fx", "own", lib.BM("crazy brown fx", "own"))
	// fmt.Println("crazy brown fox", "fox", lib.BM("crazy brown fox", "fox"))
	// fmt.Println("crazy brown fx", "fox", lib.BM("crazy brown fx", "fox"))
	// fmt.Println("crazy brown fox", "fx", lib.BM("crazy brown fox", "fx"))

}
