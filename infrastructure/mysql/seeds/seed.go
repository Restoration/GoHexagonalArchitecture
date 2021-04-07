package main

import (
	"fmt"
	"log"
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"hex-go/internal/model"
)

// Article .
type Article struct {
	Title     string
	Body      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func seeds(db *gorm.DB) error {
	for i := 0; i < 10; i++ {
		// 記事
		article := model.Articles{Title: "title", Body: "body", CreatedAt: time.Now(), UpdatedAt: time.Now()}
		if err := db.Create(&article).Error; err != nil {
			fmt.Printf("%+v", err)
		}
	}
	return nil
}

func openConnection() *gorm.DB {
	db, err := gorm.Open("mysql", "root:root@tcp(hex_db:3306)/api_db?charset=utf8mb4&parseTime=True&loc=Local")
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	return db
}

func main() {
	db := openConnection()
	defer db.Close()
	if err := seeds(db); err != nil {
		fmt.Printf("%+v", err)
		return
	}
	fmt.Println("Success to seeding")
}
