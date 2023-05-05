package lib

func KmpMatch(text string, pat string) int {
	border := ComputeBorder(pat)
	i := 0
	j := 0
	m := len(pat)
	n := len(text)

	for i < n {
		if text[i] == pat[j] {
			if j == m-1 {
				return i - m + 1
			}
			i++
			j++
		} else if j > 0 {
			j = border[j-1]
		} else {
			i++
		}
	}
	return -1
}

func ComputeBorder(pat string) []int {
	var border []int = make([]int, len(pat))
	border[0] = 0
	var j int = 0
	var i int = 1
	for i < len(pat) {
		if pat[j] == pat[i] {
			border[i] = j + 1
			j++
			i++
		} else if j > 0 {
			j = border[j-1]
		} else {
			border[i] = 0
			i++
		}
	}
	return border
}

//func main() {
//	text := "ABCDEFGABCDFDF"
//	pat := "BC"
//
//	fmt.Println("Text: ", text)
//	fmt.Println("Pattern: ", pat)
//
//	temp := (kmpMatch(text, pat))
//	fmt.Println("Index: ", temp)
//	if temp == -1 {
//		fmt.Println("Not Matched")
//	} else {
//		fmt.Println("Matched")
//	}
//}