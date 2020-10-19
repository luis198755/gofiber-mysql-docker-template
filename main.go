package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/gofiber/fiber/v2"
	"gorm.io/driver/mysql"
	_ "gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Apple struct {
	gorm.Model
	Weight  int
	Color   string
	Quality string
}

// database connection
var db *gorm.DB

func init() {
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	database := os.Getenv("DB_DATABASE")

	dsn := fmt.Sprintf("%s:%s@tcp(mariadb:3306)/%s?charset=utf8&parseTime=True&loc=Local", user, password, database)

	var err error

	retries, err := strconv.Atoi(os.Getenv("DB_RECONNECT_RETRIES"))
	if err != nil || retries == 0 {
		retries = 5
	}

	for {
		db, err = gorm.Open(mysql.New(mysql.Config{
			DSN:                       dsn,   // data source name
			DefaultStringSize:         256,   // default size for string fields
			DisableDatetimePrecision:  true,  // disable datetime precision, which not supported before MySQL 5.6
			DontSupportRenameIndex:    false, // drop & create when rename index, rename index not supported before MySQL 5.7, MariaDB
			DontSupportRenameColumn:   true,  // `change` when rename column, rename column not supported before MySQL 8, MariaDB
			SkipInitializeWithVersion: false, // auto configure based on currently MySQL version
		}), &gorm.Config{})

		if err == nil {
			break
		} else if retries > 0 {
			time.Sleep(3 * time.Second)
			retries--
			continue
		} else {
			os.Exit(1)
		}
	}

	err = db.AutoMigrate(&Apple{})
	if err != nil {
		fmt.Println("[ERROR]: failed database migration:", err)
		os.Exit(1)
	}
}

func main() {
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Hello, World 👋!")
	})

	app.Get("/apples", GetApples)
	app.Get("/apples/:id", GetApple)
	app.Post("/apples", AddApple)
	app.Put("/apples/:id", UpdateApple)
	app.Delete("/apples", FlushApples)
	app.Delete("/apples/:id", DeleteApple)

	app.Listen(":3000")

}
