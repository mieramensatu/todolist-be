// package config

// import (
// 	"fmt"
// 	"log"
// 	"os"

// 	"gorm.io/driver/postgres"
// 	"gorm.io/gorm"
// )

// func CreateDBConnection() *gorm.DB {
// 	user := os.Getenv("USER")
// 	password := os.Getenv("PASSWORD")
// 	host := os.Getenv("HOST")
// 	port := os.Getenv("PORT")
// 	dbname := os.Getenv("DBNAME")
// 	sslmode := os.Getenv("SSL_MODE")
// 	timezone := os.Getenv("TIMEZONE")

// 	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
// 		host, user, password, dbname, port, sslmode, timezone)

// 	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
// 	if err != nil {
// 		log.Fatalf("Failed to connect to database: %v", err)
// 	}

// 	return db
// }


package config

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"os"
	"time"
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("Error loading .env file")
	}
}

func CreateDBConnection() *gorm.DB {
	// Memuat file .env
	LoadEnv()

	// Memuat string connection database dari variabel environment
	dbConfig := os.Getenv("SQLSTRING")

	// Membuat koneksi database
	DB, err := gorm.Open(mysql.Open(dbConfig), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			SingularTable: true,
		},
	})
	if err != nil {
		panic(err)
	}

	// Mengatur konfigurasi kumpulan koneksi database
	dbO, err := DB.DB()
	if err != nil {
		panic(err)
	}
	dbO.SetConnMaxIdleTime(time.Duration(1) * time.Minute)
	dbO.SetMaxIdleConns(2)
	dbO.SetConnMaxLifetime(time.Duration(1) * time.Minute)

	// Munculkan kesalahan jika record tidak ditemukan
	DB.Statement.RaiseErrorOnNotFound = true

	return DB
}