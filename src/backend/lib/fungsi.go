package lib

import (
	"fmt"
	"log"
	"math"
	"os"
	"regexp"

	"database/sql"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func SearchHighestPercentage(source string, listPertanyaan []string) (float64, int) {
	var highest float64
	var index int
	for i := 0; i < len(listPertanyaan); i++ {
		distance := Levenshtein([]rune(source), []rune(listPertanyaan[i]))
		maxx := math.Max(float64(len(source)), float64(len(listPertanyaan[i])))
		percentage := 100 - (float64(distance) / maxx * 100)
		if percentage > highest {
			highest = percentage
			index = i
		}
	}
	return highest, index
}

func SearchSimilarQuestion(source string, listPertanyaan []string) []int {
	var index []int
	for i := 0; i < len(listPertanyaan); i++ {
		distance := Levenshtein([]rune(source), []rune(listPertanyaan[i]))
		maxx := math.Max(float64(len(source)), float64(len(listPertanyaan[i])))
		percentage := 100 - (float64(distance) / maxx * 100)
		if percentage > 40 {
			index = append(index, i)
		}
	}
	return index
}
func getEnv(key string) string {
	err := godotenv.Load("models/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func findKMP(text string, listPertanyaan []string) int {
	var index int
	for i := 0; i < len(listPertanyaan); i++ {
		index = KmpMatch(text, listPertanyaan[i])
		if index != -1 {
			return i
		}
	}
	return index
}

func findBM(text string, listPertanyaan []string) int {
	var index int
	for i := 0; i < len(listPertanyaan); i++ {
		index = BM(text, listPertanyaan[i])
		if index != -1 {
			return i
		}
	}
	return index
}

func Levenshtein(str1, str2 []rune) int {
	s1len := len(str1)
	s2len := len(str2)
	column := make([]int, len(str1)+1)

	for y := 1; y <= s1len; y++ {
		column[y] = y
	}
	for x := 1; x <= s2len; x++ {
		column[0] = x
		lastkey := x - 1
		for y := 1; y <= s1len; y++ {
			oldkey := column[y]
			var incr int
			if str1[y-1] != str2[x-1] {
				incr = 1
			}

			column[y] = Min(column[y]+1, column[y-1]+1, lastkey+incr)
			lastkey = oldkey
		}
	}
	return column[s1len]
}

func Min(a, b, c int) int {
	if a < b {
		if a < c {
			return a
		}
	} else {
		if b < c {
			return b
		}
	}
	return c
}

func Utama(text string, val bool) string {
	text = text[1:len(text)]
	fmt.Println(text)
	var err error
	var regexCalcu *regexp.Regexp
	var regexCalen *regexp.Regexp

	var textCalen = FindPrefixCalendar(text)
	var textCalcu = FindPrefixCalculator(text)
	regexCalcu, err = regexp.Compile(`[\(]?[-+]?\d*\.?\d+[\)]?\s*([-+*/](\s?)[\(]*?\s*\d*\.?\d+[\)]?\s*)*`)
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
	var hasilCalcu = regexCalcu.FindAllString(textCalcu, -1)
	var hasilCalen = regexCalen.FindAllString(textCalen, -1)
	if len(hasilCalen) != 0 {
		fmt.Println("ini kalender")
		if IsDateValid(hasilCalen[0]) {
			return ("Hari " + GetDay(hasilCalen[0]))
		} else {
			return ("Invalid Date")
		}
	} else if len(hasilCalcu) != 0 {
		if len(hasilCalcu) == 1 {
			fmt.Println("ini kalkulator")
			fmt.Println(hasilCalcu[0])
			fmt.Println(Calculator(hasilCalcu[0]))
			return Calculator(hasilCalcu[0])
		} else {
			return ("Sintaks tidak sesuai")
		}
	} else if ValidateAddDatabase(text) != "notFound" {
		fmt.Println(ValidateAddDatabase(text))
		return AddDatabase(text, pertanyaan)
	} else if ValidateRemoveDatabase(text) != "notFound" {
		fmt.Println(ValidateRemoveDatabase(text))
		return RemoveDatabase(text, pertanyaan)
	} else {
		fmt.Println("ini pertanyaan")
		text = FindPrefixQ(text)
		var retVal int
		if !val {
			retVal = findKMP(text, pertanyaan)
			fmt.Println("Masuk KMP")
		} else {
			retVal = findBM(text, pertanyaan)
			fmt.Println("Masuk BM")
		}
		fmt.Println(retVal)
		if retVal == -1 {
			fmt.Println("Masuk Levenshtein")
			percentage, index := SearchHighestPercentage(text, pertanyaan)
			fmt.Println(percentage)
			if percentage < 41 {
				return ("Maaf, saya tidak mengerti")
			} else if percentage >= 41 && percentage < 80 {
				mirip := SearchSimilarQuestion(text, pertanyaan)
				retMirip := "Apakah maksud anda "
				for i := 0; i < len(mirip); i++ {
					if i == len(mirip)-1 {
						retMirip += rows[mirip[i]].Pertanyaan + " ?"
					} else {
						retMirip += rows[mirip[i]].Pertanyaan + ", "
					}
				}
				fmt.Println(retMirip)
				return retMirip
			} else {
				fmt.Println(percentage, index)
				fmt.Println(rows[index].Jawaban)
				return (rows[index].Jawaban)
			}
		} else {
			fmt.Println("langsung")
			fmt.Println(rows[retVal].Jawaban)
			return (rows[retVal].Jawaban)
		}
	}
}
