package lib

import (
	"fmt"
	"regexp"
	"strconv"
)

func Contains(arr []string, str string) bool {
	for _, a := range arr {
		if a == str {
			return true
		}
	}
	return false
}

func IndexOf(arr []string, str string) int {
	for i, a := range arr {
		if a == str {
			return i
		}
	}
	return -1
}

func ConverArrStrToFloat(arr []string) []float64 {
	var arrFloat []float64
	var temp float64
	for i := 0; i < len(arr); i++ {
		temp, _ = strconv.ParseFloat(arr[i], 64)
		arrFloat = append(arrFloat, temp)
	}
	return arrFloat
}

func FindResult(operand []string, angka []float64) float64 {
	var result float64
	var idx int
	for len(operand) != 0 {
		idxp := IndexOf(operand, "+")
		idxs := IndexOf(operand, "-")
		idxm := IndexOf(operand, "*")
		idxd := IndexOf(operand, "/")

		if idxp < idxs && idxp > -1 {
			if Contains(operand, "-") {
				idx = IndexOf(operand, "-")
				result = angka[idx] - angka[idx+1]
			}
			if Contains(operand, "+") {
				idx = IndexOf(operand, "+")
				result = angka[idx] + angka[idx+1]
			}
		} else {
			if Contains(operand, "+") {
				idx = IndexOf(operand, "+")
				result = angka[idx] + angka[idx+1]
			}
			if Contains(operand, "-") {
				idx = IndexOf(operand, "-")
				result = angka[idx] - angka[idx+1]
			}
		}

		if idxm < idxd && idxm > -1 {
			if Contains(operand, "/") {
				idx = IndexOf(operand, "/")
				result = angka[idx] / angka[idx+1]
			}
			if Contains(operand, "*") {
				idx = IndexOf(operand, "*")
				result = angka[idx] * angka[idx+1]

			}
		} else {
			if Contains(operand, "*") {
				idx = IndexOf(operand, "*")
				result = angka[idx] * angka[idx+1]

			}
			if Contains(operand, "/") {
				idx = IndexOf(operand, "/")
				result = angka[idx] / angka[idx+1]
			}
		}
		operand = append(operand[:idx], operand[idx+1:]...)
		angka = append(angka[:idx+1], angka[idx+2:]...)
		angka[idx] = result
	}
	return result

}

func Calculator(text string) string {
	operand := []string{}
	angka := []string{}
	var temp string
	count := 0

	for i := 0; i < len(text); i++ {
		if text[i] == '+' || text[i] == '-' || text[i] == '*' || text[i] == '/' {
			operand = append(operand, string(text[i]))
		} else {
			if i != len(text)-1 {
				if count == 0 {
					temp = string(text[i])
					if text[i+1] == '+' || text[i+1] == '-' || text[i+1] == '*' || text[i+1] == '/' {
						angka = append(angka, (temp))
						temp = ""
						count = 0
					}
					count++
				} else {
					temp = temp + string(text[i])
					if text[i+1] == '+' || text[i+1] == '-' || text[i+1] == '*' || text[i+1] == '/' {
						angka = append(angka, (temp))
						temp = ""
						count = 0
					}
				}
			} else {
				temp = temp + string(text[i])
				angka = append(angka, (temp))
				temp = ""
				count = 0
			}
		}
	}
	var angkaFloat []float64
	angkaFloat = ConverArrStrToFloat(angka)

	var result = FindResult(operand, angkaFloat)
	return ("Hasilnya adalah " + fmt.Sprintf("%.2f", result))
}

func FindPrefixCalculator(text string) string {
	knowledge_base := map[string]string{
		"Berapa hasil (.*)":    "X",
		"(.*) hasilnya adalah": "X",
		"(.*) berapa hasilnya": "X",
		"Hasil dari (.*)":      "X",
		"(.*)?":                "X",
	}
	notFound := "notFound"
	for key, value := range knowledge_base {
		m := regexp.MustCompile(key)
		if m.MatchString(text) {
			answer := value
			len_groups := len(m.FindString(text))
			if len_groups == 0 {
				return answer
			} else {
				x := m.FindStringSubmatch(text)[1]
				answer = regexp.MustCompile("X").ReplaceAllString(answer, x)
				return answer
			}
		}
	}
	return notFound
}

// func main() {
// 	scanner := bufio.NewScanner(os.Stdin)
// 	scanner.Scan()
// 	text := scanner.Text()
// 	var regex, err = regexp.Compile(`[-+]?[0-9]*\.?[0-9]+([-+*/]?([0-9]*\.?[0-9]+))*$`)

// 	if err != nil {
// 		fmt.Println(err.Error())
// 	}
// 	var hasil = regex.FindAllString(text, -1)
// 	if len(hasil) == 0 {
// 		fmt.Println("Sintaks persamaan tidak sesuai")
// 	} else {
// 		fmt.Println(hasil[0])
// 		calculator(hasil[0])
// 	}
// }
