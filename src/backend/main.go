package main

import (
	// "bufio"
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

var db *sql.DB

// type Album struct {
// 	ID     int64
// 	Title  string
// 	Artist string
// 	Price  float32
// }

type Dataa struct {
	id_data    int64
	pertanyaan string
	jawaban    string
}

type History struct {
	id_history int64
	pertanyaan string
	jawaban    string
}

// func albumsByArtist() []Album {
// 	var albums []Album

// 	rows, _ := db.Query("SELECT * FROM album")

// 	defer rows.Close()
// 	for rows.Next() {
// 		var alb Album
// 		rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price)
// 		albums = append(albums, alb)
// 	}

// 	return albums
// }

func getAllData() []Dataa {
	var data []Dataa

	rows, _ := db.Query("SELECT * FROM data")

	defer rows.Close()
	for rows.Next() {
		var dataa Dataa
		rows.Scan(&dataa.id_data, &dataa.pertanyaan, &dataa.jawaban)
		data = append(data, dataa)
	}

	return data
}

// func getAllHistory() []History {
// 	// An albums slice to hold data from returned rows.
// 	var albums []History

// 	rows, _ := db.Query("SELECT * FROM history")

// 	defer rows.Close()
// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var alb History
// 		rows.Scan(&alb.id_history, &alb.pertanyaan, &alb.jawaban)
// 		albums = append(albums, alb)
// 	}

// 	return albums
// }

// func read_file(namaFile string) []string {
// 	file, err := os.Open(namaFile)
// 	if err != nil {
// 		panic(err.Error())
// 	}

// 	defer file.Close()
// 	scanner := bufio.NewScanner(file)

// 	var isi []string
// 	for scanner.Scan() {
// 		isi = append(isi, scanner.Text())
// 	}
// 	return isi
// }

func getEnv(key string) string {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func printData(data []Dataa) {
	for i := 0; i < len(data); i++ {
		fmt.Println(data[i])
	}
}

func main() {

	var err error
	db, err = sql.Open("mysql", getEnv("DBUSER")+":"+getEnv("DBPASS")+"@tcp(localhost:"+getEnv("DBPORT")+")/"+getEnv("DBNAME"))
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rows := getAllData()
	printData(rows)

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
