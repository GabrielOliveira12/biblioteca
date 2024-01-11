package db

import (
	"biblioteca/model"
	"log"
	"os"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Database *gorm.DB

func Connect() {

	dsn := "user=gabriel password=123 dbname=biblioteca host=localhost port=5432 sslmode=disable TimeZone=UTC"

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags),
			logger.Config{
				LogLevel: logger.Info,
			},
		),
	})

	if err != nil {
		panic("Erro ao conectar ao banco de dados: " + err.Error())
	}

	Database = db

	Database.AutoMigrate(&model.Livro{}, &model.Usuario{})
}
