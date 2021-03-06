package database

import (
	"fmt"
	"log"

	"github.com/piLinux/GoREST/config"

	"github.com/jinzhu/gorm"

	// Import MySQL database driver
	_ "github.com/jinzhu/gorm/dialects/mysql"

	// Import PostgreSQL database driver
	_ "github.com/jinzhu/gorm/dialects/postgres"

	// Import SQLite3 database driver
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

// DB global variable to access gorm
var DB *gorm.DB
var err error

// InitDB - function to initialize db
func InitDB() *gorm.DB {
	var db = DB

	configureDB := config.ConfigMain()

	driver := configureDB.Database.DbDriver
	username := configureDB.Database.DbUser
	password := configureDB.Database.DbPass
	database := configureDB.Database.DbName
	host := configureDB.Database.DbHost
	port := configureDB.Database.DbPort

	switch driver {
	case "mysql":
		db, err = gorm.Open(driver, username+":"+password+"@tcp("+host+":"+port+")/"+database+"?charset=utf8mb4&parseTime=True&loc=Local")
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Fatalln(err)
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful!")
		}
		break
	case "postgres":
		db, err = gorm.Open(driver, "host="+host+" port="+port+" user="+username+" dbname="+database+" password="+password)
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Fatalln(err)
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful!")
		}
		break
	case "sqlite3":
		db, err = gorm.Open(driver, database)
		if err != nil {
			// fmt.Println("DB err: ", err)
			log.Fatalln(err)
		}
		// Only for debugging
		if err == nil {
			fmt.Println("DB connection successful!")
		}
		break
	default:
		log.Fatalln("The driver " + driver + " is not implemented yet")
	}

	DB = db

	return DB
}

// GetDB - get a connection
func GetDB() *gorm.DB {
	return DB
}
