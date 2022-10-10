package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"os"
	"strconv"
	"task-manager/database"
	"task-manager/models"
	"time"
)

func SignUpUser(c *gin.Context) {
	var newUser models.User
	if err := c.BindJSON(&newUser); err != nil {
		return
	}
	if newUser.Username == "" {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Username can't be empty"})
		return
	}
	if len(newUser.Password) < 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Password must be more than 5 characters"})
		return
	}
	hash, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), 10)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to hash password"})
		return
	}
	if err := database.DB.Create(&models.User{Username: newUser.Username, Password: string(hash)}).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User with this username already exists"})
		return
	}
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success"})
}

func LoginUser(c *gin.Context) {
	var tempUser models.User
	var user models.User
	if err := c.BindJSON(&tempUser); err != nil {
		return
	}
	database.DB.First(&user, "username = ?", tempUser.Username)
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(tempUser.Password))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid username or password"})
		return
	}
	// Create a new token object, specifying signing method and the claims
	// you would like it to contain.
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		Subject:   strconv.Itoa(int(user.ID)),
		ExpiresAt: &jwt.NumericDate{Time: time.Now().Add(time.Hour * 24)},
	})

	// Sign and get the complete encoded token as a string using the secret
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	c.SetSameSite(http.SameSiteLaxMode)
	c.SetCookie("Authorization", tokenString, 3600*24, "", "", false, true)
	c.IndentedJSON(http.StatusCreated, gin.H{"message": "success"})
}

func ValidateUser(c *gin.Context) {
	user, _ := c.Get("user")
	c.JSON(http.StatusOK, gin.H{"message": user})
}
