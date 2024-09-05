package controllers

import (
	"net/http"

	"RankEdge/internal/models"
	"RankEdge/internal/repositories"
	"RankEdge/internal/services"
	"RankEdge/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// SignUp godoc
//
//	@Summary		Create a new user
//	@Description	Create a new user
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object{email=string,password=string,first_name=string,last_name=string}	true	"User details"
//	@Success		201		{object}	object{user=object{user_id=string,created_at=string,email=string,first_name=string,last_name=string,role=string}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		409		{object}	object{error=string}
//	@Router			/auth/signup [post]
func SignUp(c *fiber.Ctx) error {
	userRepo := repositories.NewUserRepository(utils.DB)
	// Get the JSON body and decode into variables
	var user models.User
	if err := c.BodyParser(&user); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := services.NewUserService(userRepo).SignUp(&user)
	if err != nil {
		return c.Status(fiber.StatusConflict).JSON(fiber.Map{"error": err.Error()})
	}

	utils.SetToken(c, user)

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"user": user,
	})
}

type body struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

// SignIn godoc
//
//	@Summary		Sign in
//	@Description	Signs in the user, returns the user details and sets the jwt token
//	@Tags			Auth
//	@Accept			json
//	@Produce		json
//	@Param			user	body		object{email=string,password=string}	true	"Login details"
//	@Success		200		{object}	object{user=object{user_id=string,email=string,created_at=string,first_name=string,last_name=string,role=string}}
//	@Failure		400		{object}	object{error=string}
//	@Failure		401		{object}	object{error=string}
//	@Failure		404		{object}	object{error=string}
//	@Router			/auth/signin [post]
func SignIn(c *fiber.Ctx) error {
	userRepo := repositories.NewUserRepository(utils.DB)
	// Get the JSON body and decode into variables
	var body body
	if err := c.BodyParser(&body); err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid input"})
	}

	user, err := services.NewUserService(userRepo).SignIn(body.Email, body.Password)
	if err != nil {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "Invalid credentials"})
	}

	utils.SetToken(c, *user)

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"user": user,
	})
}

// SignOut godoc
//
//	@Summary		Sign out
//	@Description	Signs out the user
//	@Tags			Auth
//	@Produce		json
//	@Success		204
//	@Router			/auth/signout [get]
func SignOut(c *fiber.Ctx) error {
	c.ClearCookie("token")
	return c.SendStatus(http.StatusNoContent)
}
