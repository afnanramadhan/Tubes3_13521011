package lib

// BM the pattern in the text
func BM(text, pattern string) int {
	charTable := [256]int{}
	// initialise default values for bad char table
	for i := range charTable {
		charTable[i] = len(pattern)
	}
	// creating the bad char table based on the pattern
	// using the formula len(pattern) - i - 1
	for i := 0; i < len(pattern)-1; i++ {
		charTable[pattern[i]] = len(pattern) - i - 1
	}
	// i is the index for text
	i := len(pattern) - 1

	// scan all the text
	for i < len(text) {
		j := len(pattern) - 1
		for j >= 0 && pattern[j] == text[i] {
			j--
			i--
		}
		if j < 0 {
			return i + 1
		}
		// shift the pattern based on badchar table
		i = i + charTable[text[i]]
	}

	// nothing has been found
	return -1
}

// func main() {
// 	fmt.Println("crazy brown fx", "own", BM("crazy brown fx", "own"))
// 	fmt.Println("crazy brown fox", "fox", BM("crazy brown fox", "fox"))
// 	fmt.Println("crazy brown fx", "fox", BM("crazy brown fx", "fox"))
// 	fmt.Println("crazy brown fox", "fx", BM("crazy brown fox", "fx"))
// }