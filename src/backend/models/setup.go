package models

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"github.com/joho/godotenv"
	"os"
)

type Data struct {
	Id_data         int64  `gorm:"primaryKey" json:"id_data"`
	Pertanyaan string `gorm:"type:varchar(255)" json:"pertanyaan"`
	Jawaban   string `gorm:"type:varchar(255)" json:"jawaban"`
}

var DB *gorm.DB

func getEnv(key string) string {
	err := godotenv.Load("models/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
	return os.Getenv(key)
}

func ConnectDatabase() {
	database, err := gorm.Open(mysql.Open(getEnv("DBUSER")+":"+getEnv("DBPASS")+"@tcp(localhost:"+getEnv("DBPORT")+")/"+getEnv("DBNAME")))
	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&Data{})

	DB = database
}