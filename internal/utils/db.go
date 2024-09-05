package utils

import (
	"crypto/rand"
	"encoding/hex"
	"errors"
	"fmt"
	"log"
	"os"
	"time"

	"RankEdge/internal/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// ConnectToDb connects to the postgresql database
func ConnectToDb() {
	var err error
	host := os.Getenv("MYSQL_HOST")
	port := os.Getenv("MYSQL_PORT")
	user := os.Getenv("MYSQL_USER")
	password := os.Getenv("MYSQL_PASSWORD")
	dbName := os.Getenv("MYSQL_DB")

	if host == "" {
		log.Fatal(errors.New("MYSQL_HOST environment variable is not set"))
	}

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Asia%%2FShanghai", user, password, host, port, dbName)
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		time.Sleep(5 * time.Second)
		ConnectToDb()
	}
	DB.Debug()
	log.Println("Connected to database")
}

// Migrate migrates the models
func Migrate() {
	DB.AutoMigrate(
		&models.User{},
	)
}

// GenerateRandomID generates a random ID as a hexadecimal string
func GenerateRandomID() string {
	// Define the length of the ID in bytes (e.g., 16 bytes for 128-bit ID)
	length := 16
	// Create a byte slice to hold the random data
	randomBytes := make([]byte, length)
	// Generate random bytes
	rand.Read(randomBytes)
	// Encode the random bytes as a hexadecimal string
	userID := hex.EncodeToString(randomBytes)

	return userID
}
