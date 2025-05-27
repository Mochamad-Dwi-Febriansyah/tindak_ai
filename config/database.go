package config

import (
	"fmt"
	"log"
	"os"
	"time"
	"tindak_ai/internal/domain"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func InitDB() {
	if err := godotenv.Load(); err != nil {
		// panic(err)
		log.Fatalf("Failed to load .env file: %v", err)
	}

	logFile, err := os.OpenFile("tmp/gorm.log", os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		log.Fatalf("Failed to open log file: %v", err)
	}
	
	newLogger := logger.New(
		log.New(logFile, "\r\n", log.LstdFlags),
		logger.Config{
			SlowThreshold:             time.Second,
			LogLevel:                  logger.Info,
			IgnoreRecordNotFoundError: true,
			Colorful:                  false,
		},
	)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"),
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})

	if err != nil { 
		// panic(err)
		log.Fatalf("Failed to connect to database: %v", err)
	}

	db.AutoMigrate(
		&domain.Users{}, 
		&domain.Role{}, 
		&domain.Permission{}, 
		&domain.RolePermission{}, 
		&domain.UserPermission{}, 
		&domain.Institution{},
		&domain.Logging{},
		&domain.UserInstitution{},
		&domain.Complaint{},
		&domain.ComplaintStatusHistory{},
		&domain.ComplaintRating{},
		&domain.NewsArticle{},
		&domain.InstitutionRating{},
	)

	DB = db
}
