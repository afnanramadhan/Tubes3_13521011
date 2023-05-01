package main

import ( 
	"fmt"
)

func kmpMatch(input string, pat string) int {
	border := computeBorder(pat)
	i := 0
	j := 0
	m := len(pat)
	n := len(input)

	for i < n{
		if input[i] == pat[j] {
			if j == m-1 {
				return 0
			}
			i++
			j++
		} else if (j > 0) {
			j = border[j-1]
		} else {
			i++
		}
	}
	return -1
}

func computeBorder(pat string) []int  {
	var border []int = make([]int, len(pat))
	border[0] = 0
	var j int = 0
	var i int = 1
	for i<len(pat) {
		if (pat[j]==pat[i]){
			border[i] = j+1
			j++
			i++
		}else if (j>0){
			j = border[j-1]
		}else{
			border[i] = 0
			i++
		}
	}
	return border
}

func main()  {
	text:= "ABCDEFGABCDFDF"
	pat:= "ABCDF"

	fmt.Println("Text: ",text)
	fmt.Println("Pattern: ",pat)
	
	temp:=(kmpMatch(text,pat))
	if(temp==0){
		fmt.Println("Matched")
	}else{
		fmt.Println("Not Matched")
	}
}