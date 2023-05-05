package lib

import (
	"backend/models"
	"database/sql"
	"fmt"
)

var db *sql.DB

func GetAllData() []models.Data {
	var data []models.Data

	rows, _ := db.Query("SELECT * FROM data")

	defer rows.Close()
	for rows.Next() {
		var dataa models.Data
		rows.Scan(&dataa.Id_data, &dataa.Pertanyaan, &dataa.Jawaban)
		data = append(data, dataa)
	}

	return data
}
func GetPertanyaan(data []models.Data) []string {
	var pertanyaan []string
	for i := 0; i < len(data); i++ {
		pertanyaan = append(pertanyaan, data[i].Pertanyaan)
	}
	return pertanyaan
}

func AddToDatabase(pertanyaan string, jawaban string, listPertanyaan []string) string {
	if ValidatePertanyaan(pertanyaan, listPertanyaan) == -1 {
		rows, err := db.Query("UPDATE data SET jawaban = ? WHERE pertanyaan = ?", jawaban, pertanyaan)
		if err != nil {
			return "Gagal mengupdate jawaban"
		}
		defer rows.Close()
		return "Pertanyaan " + pertanyaan + " sudah ada!. jawaban di update ke " + jawaban
	} else {
		rows, err := db.Query("INSERT INTO data (pertanyaan, jawaban) VALUES (?, ?)", pertanyaan, jawaban)
		if err != nil {
			return "Gagal menambahkan pertanyaan dan jawaban"
		}
		defer rows.Close()
		return "Berhasil menambahkan pertanyaan dan jawaban"
	}
}

func RemoveFromDatabase(pertanyaan string, listPertanyaan []string) string {
	fmt.Println(IsPertanyaanExist(pertanyaan, listPertanyaan))
	if IsPertanyaanExist(pertanyaan, listPertanyaan) == -1 {
		return "Tidak ada pertanyaan "+pertanyaan+" pada database."
	} else {
		rows, err := db.Query("DELETE FROM data WHERE pertanyaan = ?", pertanyaan)
		if err != nil {
			return "Gagal menghapus pertanyaan"
		}
		defer rows.Close()
		return "Pertanyaan "+pertanyaan+" berhasil dihapus"
	}
}

func ValidatePertanyaan(pertanyaan string, listPertanyaan []string) int {
	for i := 0; i < len(listPertanyaan); i++ {
		if pertanyaan == listPertanyaan[i] {
			return -1
		}
	}
	return 0
}


func IsPertanyaanExist(pertanyaan string, listPertanyaan []string) int{
	for i := 0; i < len(listPertanyaan); i++ {
		if pertanyaan == listPertanyaan[i] {
			fmt.Println(i)
			return i
		}
	}
	return -1
}