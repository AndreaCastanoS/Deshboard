package auth

import (
	"challengeFinal/database"
	"challengeFinal/models"
	"challengeFinal/shared"
	"encoding/json"
	"fmt"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gofrs/uuid"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func Register(c *gin.Context) {
	var userInput models.UserInput
	if err := c.ShouldBindJSON(&userInput); err != nil {
		c.JSON(http.StatusBadRequest, err.Error())
		return
	}
	user := models.User{
		Email:         userInput.Email,
		Password:      userInput.Password,
		DisplayConfig: "{}",
	}
	database.CreateDbConnection()
	if tx := database.DBConn.Create(&user); tx.Error != nil {
		if strings.Contains(tx.Error.Error(), "duplicate") || strings.Contains(tx.Error.Error(), "UNIQUE") {
			c.JSON(http.StatusConflict, gin.H{"error": "Email already exists"})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": tx.Error.Error()})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"email": user.Email,
		"id":    user.ID,
	})
}

func Login(c *gin.Context) {
	var input models.UserInput
	var user models.User

	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid input"})
		return
	}

	database.CreateDbConnection()
	if err := database.DBConn.Where("email = ?", input.Email).First(&user).Error; err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(input.Password)); err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid credentials"})
		return
	}

	sessionToken := uuid.NewV5(uuid.UUID{}, "session").String()

	session := shared.Session{
		Uid:        user.ID,
		ExpiryTime: time.Now().Add(24 * time.Hour),
	}
	shared.Sessions[sessionToken] = session

	claims := shared.Payload{
		MapClaims: jwt.MapClaims{
			"iat": jwt.NewNumericDate(time.Now()),
			"exp": jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
		},
		Session: sessionToken,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signinKey := []byte(os.Getenv("JWT_SECRET_KEY"))
	tokenString, err := token.SignedString(signinKey)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not create token"})
		return
	}
	var displayConfig map[string]interface{}
	if err := json.Unmarshal([]byte(user.DisplayConfig), &displayConfig); err != nil {
		displayConfig = nil
	}

	c.JSON(http.StatusOK, gin.H{
		"token":          tokenString,
		"session_token":  sessionToken,
		"display_config": displayConfig,
	})
}

func GetUserConfig(userID uint) (map[string]interface{}, error) {
	var user models.User

	database.CreateDbConnection()
	if err := database.DBConn.First(&user, userID).Error; err != nil {
		fmt.Printf("Error al buscar usuario con ID %d: %v\n", userID, err)
		return nil, err
	}

	var config map[string]interface{}
	if err := json.Unmarshal([]byte(user.DisplayConfig), &config); err != nil {
		fmt.Printf("Error al deserializar DisplayConfig para userID %d: %v\n", userID, err)
		return nil, err
	}

	fmt.Printf("DisplayConfig obtenido para userID %d: %v\n", userID, config)
	return config, nil
}

func UpdateUserDisplayConfig(userID uint, config map[string]interface{}) error {
	var user models.User
	database.CreateDbConnection()

	if err := database.DBConn.First(&user, userID).Error; err != nil {
		return err
	}

	configJSON, err := json.Marshal(config)
	if err != nil {
		return err
	}

	return database.DBConn.Model(&user).Update("display_config", string(configJSON)).Error
}
