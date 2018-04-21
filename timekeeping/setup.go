package timekeeping

import (
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
)

var db *gorm.DB

// Setup opens the database and creates the tables
func Setup() error {
	var openErr error
	db, openErr = gorm.Open("postgres", "user=timekeeper password=test dbname=timekeeper")
	if openErr != nil {
		return openErr
	}

	if openErr = db.DB().Ping(); openErr != nil {
		return openErr
	}
	db.CreateTable(&Tag{})
	db.Create(&Interval{})
	db.Create(&Task{})
	return nil
}

// Cleanup closes the database
func Cleanup() error {
	return db.Close()
}
