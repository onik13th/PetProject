package repositories

import (
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func NewDatabase() *gorm.DB {
	dsn := fmt.Sprintf(
		"user=%s password=%s dbname=%s port=%s",
		"chmonik", "22134289", "petProgect", "5432",
	)
	db, err := gorm.Open(postgres.New(postgres.Config{
		DSN:                  dsn,
		PreferSimpleProtocol: true,
	}), &gorm.Config{})
	if err != nil {
		panic("Ошибка подключения к базе данных")
	}

	db.AutoMigrate(&Book{})

	return db
}
