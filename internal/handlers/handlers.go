package handlers

import (
	"net/http"
	"time"

	"github.com/dqfan2012/learngin/internal/models"
	"github.com/dqfan2012/learngin/pkg/db"

	"github.com/gin-gonic/gin"
)

type UserResponse struct {
	ID              uint       `json:"id"`
	FirstName       string     `json:"first_name"`
	LastName        string     `json:"last_name"`
	Email           string     `json:"email"`
	Password        string     `json:"password"`
	Role            string     `json:"role"`
	RememberToken   string     `json:"remember_token"`
	CreatedAtUTC    time.Time  `json:"created_at_utc"`
	UpdatedAtUTC    time.Time  `json:"updated_at_utc"`
	CreatedAtLocal  time.Time  `json:"created_at_local"`
	UpdatedAtLocal  time.Time  `json:"updated_at_local"`
	EmailVerifiedAt *time.Time `json:"email_verified_at"`
}

func HomeHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "home.tmpl", gin.H{
		"Title": "Home",
	})
}

func StatusHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "ok",
	})
}

func GetUserHandler(c *gin.Context) {
	var user models.User
	db.DB.First(&user)

	// Convert UTC timestamps to local time
	loc, _ := time.LoadLocation("America/New_York") // Use your desired timezone here

	response := UserResponse{
		ID:              user.ID,
		FirstName:       user.FirstName,
		LastName:        user.LastName,
		Email:           user.Email,
		Password:        user.Password,
		Role:            user.Role,
		RememberToken:   user.RememberToken,
		CreatedAtUTC:    user.CreatedAt.UTC(),   // Ensure UTC time
		UpdatedAtUTC:    user.UpdatedAt.UTC(),   // Ensure UTC time
		CreatedAtLocal:  user.CreatedAt.In(loc), // Convert to local time
		UpdatedAtLocal:  user.UpdatedAt.In(loc), // Convert to local time
		EmailVerifiedAt: user.EmailVerifiedAt,
	}

	c.JSON(http.StatusOK, response)
}
