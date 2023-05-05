package lib

import (
	"regexp"
)

func ValidateRemoveDatabase(text string) string {
	knowledge_base := map[string]string{
		"[H|h]apus pertanyaan (.*)": "Pertanyaan X berhasil dihapus",
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

func RemoveDatabase(text string, pertanyaan []string) string {
	knowledge_base := map[string]string{
		"[H|h]apus pertanyaan (.*)": "Pertanyaan X berhasil dihapus",
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
				return RemoveFromDatabase(x, pertanyaan)
			}
		}
	}
	return notFound
}
