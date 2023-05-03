package main

import (
	"backend/lib"
	"bufio"
	"database/sql"
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/texttheater/golang-levenshtein/levenshtein"
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

func getPertanyaan(data []Dataa) []string {
	var pertanyaan []string
	for i := 0; i < len(data); i++ {
		pertanyaan = append(pertanyaan, data[i].pertanyaan)
	}
	return pertanyaan
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
		fmt.Println("pertanyaan :", data[i].pertanyaan)
		fmt.Println("jawaban :", data[i].jawaban)
	}
}

func searchHighestPercentage(source string, listPertanyaan []string) (float64, int) {
	var highest float64
	var index int
	for i := 0; i < len(listPertanyaan); i++ {
		distance := levenshtein.DistanceForStrings([]rune(source), []rune(listPertanyaan[i]), levenshtein.DefaultOptions)
		maxx := math.Max(float64(len(source)), float64(len(listPertanyaan[i])))
		percentage := 100 - (float64(distance) / maxx * 100)
		if percentage > highest {
			highest = percentage
			index = i
		}
	}
	return highest, index
}

func main() {

	var err error
	var regexCalcu *regexp.Regexp
	var regexCalen *regexp.Regexp
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
	pertanyaan := getPertanyaan(rows)
	// target := "apa ibukota indonsa"

	fmt.Println("Masukkan Pertanyaan :")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	text := scanner.Text()

	regexCalcu, err = regexp.Compile(`[-+]?[0-9]*\.?[0-9]+([-+*/]?([0-9]*\.?[0-9]+))*`)
	regexCalen, err = regexp.Compile(`[0-9]{1,2}/[0-9]{1,2}/[0-9]{1,4}`)
	if err != nil {
		fmt.Println(err.Error())
	}
	var hasilCalcu = regexCalcu.FindAllString(text, -1)
	var hasilCalen = regexCalen.FindAllString(text, -1)
	if len(hasilCalen) != 0 {
		fmt.Println("ini kalender")
		if lib.IsDateValid(hasilCalen[0]) {
			fmt.Println("Hari ",lib.GetDay(hasilCalen[0]))
		} else {
			fmt.Println("Invalid Date")
		}
	} else if len(hasilCalcu) != 0 {
		fmt.Println("ini kalkulator")
		fmt.Println(hasilCalcu[0])
		lib.Calculator(hasilCalcu[0])
	} else {
		fmt.Println("ini pertanyaan")
		percentage, index := searchHighestPercentage(text, pertanyaan)
		fmt.Println(percentage, index)
		fmt.Println(rows[index].jawaban)
	}

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

	// main bm
	// fmt.Println("crazy brown fx", "own", lib.BM("crazy brown fx", "own"))
	// fmt.Println("crazy brown fox", "fox", lib.BM("crazy brown fox", "fox"))
	// fmt.Println("crazy brown fx", "fox", lib.BM("crazy brown fx", "fox"))
	// fmt.Println("crazy brown fox", "fx", lib.BM("crazy brown fox", "fx"))

}
