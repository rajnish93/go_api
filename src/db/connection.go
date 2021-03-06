package db

import (
	"fmt"
	"log"
	"os"

	"github.com/rajnish93/go_api/src/modules/product"
	"github.com/rajnish93/go_api/src/modules/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	var (
		host     = os.Getenv("POSTGRES_HOST")
		user     = os.Getenv("POSTGRES_USER")
		port     = os.Getenv("POSTGRES_PORT")
		password = os.Getenv("POSTGRES_PASSWORD")
		name     = os.Getenv("POSTGRES_DB")
	)
	connectionstr := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, name)
	DB, err := gorm.Open(postgres.Open(connectionstr), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent),
	})

	if err != nil {
		log.Fatalf("Error in connect the DB %v", err)
		return nil
	}

	if DB.Error != nil {
		log.Fatalln("Any Error in connect the DB " + err.Error())
		return nil
	}
	log.Print("Connection to DB: Success")

	return DB
}

// Migrations
func InitialMigration(db *gorm.DB) {
	log.Print("Running Migrations...")
	db.AutoMigrate(&user.User{}, &product.Product{})
}