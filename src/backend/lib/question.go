package lib

import "regexp"

func FindPrefixQ(text string) string {
	knowledge_base := map[string]string{
		"[A!a]pakah (.*)":   "X",
		"[A|a]pa (.*)":      "X",
		"[S|s]iapakah (.*)": "X",
		"[S|s]iapa (.*)":    "X",
		"(.*)?":             "X",
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