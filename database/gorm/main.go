package main

import (
	"fmt"
	"log"

	"gorm.io/driver/mysql" 
	"gorm.io/gorm"
)

// Album represents a music album database model
type Album struct {
	ID     int64   `gorm:"primaryKey;autoIncrement"`
	Title  string  `gorm:"size:100;not null"`
	Artist string  `gorm:"size:100;not null"`
	Price  float32 `gorm:"not null"`
}

func main() {
	dsn := "root:@tcp(127.0.0.1:3306)/recordings?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		fmt.Println("Failed to connect to the database:", err)
		return
	}

	// Automatically migrate schema
	err = db.AutoMigrate(&Album{})
	if err != nil {
		log.Fatal("Failed to migrate schema:", err)
	}

	// Create a new album
	newAlbum := Album{Title: "Fearless", Artist: "Taylor Swift", Price: 19.99}
	result := db.Create(&newAlbum)
	if result.Error != nil {
		log.Fatal("Failed to create album:", result.Error)
	}
	fmt.Printf("New album created with ID: %d\n", newAlbum.ID)

	// Query an album by title
	var album Album
	err = db.First(&album, "title = ?", "Fearless").Error
	if err != nil {
		log.Fatal("Failed to find album:", err)
	}
	fmt.Printf("Album found: %+v\n", album)

	// Update an album's price
	db.Model(&album).Update("Price", 17.99)
	fmt.Printf("Album price updated to: $%.2f\n", album.Price)

	// Delete an album
	db.Delete(&album)
	fmt.Printf("Album deleted with ID: %d\n", album.ID)
}
