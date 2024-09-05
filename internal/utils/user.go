package utils

import (
	"os"
	"time"

	"RankEdge/internal/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// createToken creates a JWT token for a user
//
// Parameters:
//
// - user: The user for whom to create the token
//
// Returns:
//
// - string: The token string
//
// - error: An error if the token could not be created
func createToken(user models.User) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.UserID,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func SetToken(c *fiber.Ctx, user models.User) error {
	tokenString, err := createToken(user)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "No token provided"})
	}

	cookie := new(fiber.Cookie)
	cookie.Name = "token"
	cookie.Value = tokenString
	cookie.Expires = time.Now().Add(24 * time.Hour)

	c.Cookie(cookie)

	return nil
}

type TokenString struct {
	Token string `json:"token"`
}
func GetToken(c *fiber.Ctx) string {
	tokenString := new(TokenString)
	if err := c.CookieParser(tokenString); err != nil {
		return ""
	}

	return tokenString.Token
}
