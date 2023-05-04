package lib

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	"backend/models"
	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

var db *sql.DB

func GetAllData() []models.Data {
	var data []models.Data

	rows, _ := db.Query("SELECT * FROM data")

	defer rows.Close()
	for rows.Next() {
		var dataa models.Data
		rows.Scan(&dataa.Id_data, &dataa.Pertanyaan, &dataa.Jawaban)
		data = append(data, dataa)
	}

	return data
}
func GetPertanyaan(data []models.Data) []string {
	var pertanyaan []string
	for i := 0; i < len(data); i++ {
		pertanyaan = append(pertanyaan, data[i].Pertanyaan)
	}
	return pertanyaan
}

func SearchHighestPercentage(source string, listPertanyaan []string) (float64, int) {
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
func getEnv(key string) string {
	err := godotenv.Load("models/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func Utama(text string) string {
	var err error
	var regexCalcu *regexp.Regexp
	var regexCalen *regexp.Regexp
	regexCalcu, err = regexp.Compile(`[-+]?[0-9]*\.?[0-9]+([-+*/]?([0-9]*\.?[0-9]+))*`)
	regexCalen, err = regexp.Compile(`[0-9]{1,2}/[0-9]{1,2}/[0-9]{1,4}`)
	if err != nil {
		fmt.Println(err.Error())
	}

	db, err = sql.Open("mysql", getEnv("DBUSER")+":"+getEnv("DBPASS")+"@tcp(localhost:"+getEnv("DBPORT")+")/"+getEnv("DBNAME"))
	if err != nil {
		panic(err.Error())
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}
	fmt.Println("Connected!")

	rows := GetAllData()
	pertanyaan := GetPertanyaan(rows)
	var hasilCalcu = regexCalcu.FindAllString(text, -1)
	var hasilCalen = regexCalen.FindAllString(text, -1)
	if len(hasilCalen) != 0 {
		fmt.Println("ini kalender")
		if IsDateValid(hasilCalen[0]) {
			return ("Hari " + GetDay(hasilCalen[0]))
		} else {
			return ("Invalid Date")
		}
	} else if len(hasilCalcu) != 0 {
		fmt.Println("ini kalkulator")
		fmt.Println(hasilCalcu[0])
		fmt.Println(Calculator(hasilCalcu[0]))
		return Calculator(hasilCalcu[0])
	} else {
		fmt.Println("ini pertanyaan")
		percentage, index := SearchHighestPercentage(text, pertanyaan)
		fmt.Println(percentage, index)
		fmt.Println(rows[index].Jawaban)
		return (rows[index].Jawaban)
	}
}
