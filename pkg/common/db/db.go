package db

import (
	"gorm.io/gorm/logger"
	"log"
	"os"

	"gin-clean-archi/pkg/common/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func Init(url string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(url), &gorm.Config{
		Logger: logger.New(
			log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
			logger.Config{
				LogLevel: logger.Info, // Log level
				Colorful: true,        // Disable color
			},
		),
	})

	if err != nil {
		log.Fatalln(err)
	}

	db.AutoMigrate(&model.User{})

	return db
}
