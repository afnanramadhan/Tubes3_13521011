package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
	"time"
)

func isCalValid(text string) bool {

	kalender := strings.Split(text, "/")
	
	tanggal,err := strconv.Atoi(kalender[0])
	bulan,err := strconv.Atoi(kalender[1])
	tahun,err := strconv.Atoi(kalender[2])
	if err != nil {
		fmt.Println(err.Error())
	}

	//Cek batas tanggalan
	if(tanggal<1 || tanggal>31){
		return false
	}
	if(bulan<1 || bulan>12){
		return false
	}
	if(tahun<0){
		return false
	}

	//Cek bulan 30 hari
	if(bulan==4 || bulan==6 || bulan==9 || bulan==11){
		if(tanggal>30){
			return false
		}
	}
	//Cek bulan februari
	if(bulan==2){
		if(isKabisat(tahun)){
			if(tanggal>29){
				return false
			}
		}else{
			if(tanggal>28){
				return false
			}
		}
	}

	return true
}

func getDay(text string) string{
	kalender := strings.Split(text, "/")
	if(len(kalender[0])==1){
		kalender[0] = "0"+kalender[0]
	}
	if(len(kalender[1])==1){
		kalender[1] = "0"+kalender[1]
	}
	if(len(kalender[2])!=4){
		for(len(kalender[2])!=4){
			kalender[2] = "0"+kalender[2]
		}
	}


	date := fmt.Sprintf("%s-%s-%s", kalender[1], kalender[0], kalender[2])
	fmt.Println(date)
	t, err := time.Parse("01-02-2006", date)
	if err != nil {
		panic(err)
	}
	return t.Weekday().String()

}

func isKabisat(tahun int) bool {
	if(tahun%4==0){
		if(tahun%100==0){
			if(tahun%400==0){
				return true
			}
			return false
		}
		return true
	}
	return false
}

func main() {
	var text string;
	fmt.Scanln(&text)
	var regex, err = regexp.Compile(`[0-9]{1,2}/[0-9]{1,2}/[0-9]{1,4}`)

	if err != nil {
		fmt.Println(err.Error())
	}

	var hasil = regex.FindAllString(text, -1)
	if(isCalValid(hasil[0])){
		fmt.Println(getDay(hasil[0]))
	}else{
		fmt.Println("Invalid Date")
	}
}
