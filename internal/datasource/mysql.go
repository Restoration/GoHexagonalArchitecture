package datasource

import (
	"fmt"
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// DBConn DB接続用関数
func DBConn() *gorm.DB {
	dbUser := GetEnv("DB_USER_NAME", "")
	dbPass := GetEnv("DB_PASSWORD", "")
	dbHost := GetEnv("DB_HOST", "")
	dbName := GetEnv("DB_NAME", "")
	fmt.Println(dbUser)
	fmt.Println(dbPass)
	dbDriver := "mysql"
	connection := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", dbUser, dbPass, dbHost, dbName)
	db, err := gorm.Open(dbDriver, connection)
	if err != nil {
		log.Fatalf("Couldn't establish database connection: %s", err)
	}
	fmt.Println("MySQL Connected! ")
	return db
}
