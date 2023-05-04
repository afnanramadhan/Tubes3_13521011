package lib

import (
	"regexp"
)

func AddDatabase(text string) string {
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
				answer = regexp.MustCompile("X").ReplaceAllString(answer, x)
				y := m.FindStringSubmatch(text)[2]
				answer = regexp.MustCompile("Y").ReplaceAllString(answer, y)
				return answer
			}
		}
	}
	return notFound
}
