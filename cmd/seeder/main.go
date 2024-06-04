package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/bxcodec/faker/v3"
	"github.com/dqfan2012/learngin/internal/models"
	"github.com/dqfan2012/learngin/pkg/db"
	"github.com/joho/godotenv"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func RandStringBytes(r *rand.Rand, n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[r.Intn(len(letterBytes))]
	}
	return string(b)
}

func main() {
	// Load environment variables from .env file
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	} else {
		log.Println(".env file loaded successfully")
	}

	// Initialize the database
	db.Init()

	// Seed a user
	seedUser()
}

func seedUser() {
	// Use a new random source
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	now := time.Now().UTC() // Get current time in UTC

	user := models.User{
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      "password",
		Role:          "publisher",
		RememberToken: RandStringBytes(r, 10),
		CreatedAt:     now,
		UpdatedAt:     now,
	}

	if err := db.DB.Create(&user).Error; err != nil {
		log.Fatalf("Failed to seed user: %v", err)
	}

	log.Printf("Seeded user: %+v\n", user)
}
