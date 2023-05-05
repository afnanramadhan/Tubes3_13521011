package lib

import (
	"fmt"
	"regexp"
)

func ValidateAddDatabase(text string) string {
	knowledge_base := map[string]string{
		"[T|t]ambah pertanyaan (.*) dengan jawaban (.*)": "Pertanyaan X dan Jawaban Y berhasil ditambahkan",
	}
	notFound := "notFound"
	for key, value := range knowledge_base {
		m := regexp.MustCompile(key)
		if m.MatchString(text) {
			answer := value
			len_groups := len(m.FindString(text))
			if len_groups == 0 {
				fmt.Println("adfad")
				fmt.Println(answer)
				return answer
			} else {
				x := m.FindStringSubmatch(text)[1]
				fmt.Println(x)
				answer = regexp.MustCompile("X").ReplaceAllString(answer, x)
				y := m.FindStringSubmatch(text)[2]
				fmt.Println(y)
				answer = regexp.MustCompile("Y").ReplaceAllString(answer, y)
				
				return answer
			}
		}
	}
	return notFound
}

func AddDatabase(text string, listPertanyaan []string) string {
	knowledge_base := map[string]string{
		"[T|t]ambah pertanyaan (.*) dengan jawaban (.*)": "Pertanyaan X dan Jawaban Y berhasil ditambahkan",
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
				fmt.Println(x)
				answer = regexp.MustCompile("X").ReplaceAllString(answer, x)
				y := m.FindStringSubmatch(text)[2]
				fmt.Println(y)
				answer = regexp.MustCompile("Y").ReplaceAllString(answer, y)
				
				return AddToDatabase(x, y, listPertanyaan)
			}
		}
	}
	return notFound
}