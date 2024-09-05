package controllers

import (
	"net/http"
	"strconv"

	"RankEdge/internal/models"
	"RankEdge/internal/repositories"
	"RankEdge/internal/services"
	"RankEdge/internal/utils"

	"github.com/gofiber/fiber/v2"
)

// CreateLeaderBoard godoc
//	@Summary		Create a new leaderboard
//	@Description	Create a new leaderboard
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Success		201		{object}	string
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name} [put]
func CreateLeaderBoard(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)

	name := c.Params("name")
	err := services.NewLeaderboardService(leaderboardRepo).CreateLeaderboard(name)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusCreated).JSON(fiber.Map{
		"leaderboard": name,
	})
}

// GetLeaderboards godoc
//	@Summary		Get all leaderboards
//	@Description	Get all leaderboards
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Success		200	{object}	string
//	@Failure		500	{object}	string
//	@Router			/leaderboard [get]
func GetLeaderboards(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)

	leaderboards, err := services.NewLeaderboardService(leaderboardRepo).GetLeaderboards()
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"leaderboards": leaderboards,
	})
}

// GetLeaderboard godoc
//	@Summary		Get a leaderboard
//	@Description	Get a leaderboard
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Success		200		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name} [get]
func GetLeaderboard(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)
	boardName := c.Params("name")

	leaderboard, err := services.NewLeaderboardService(leaderboardRepo).GetLeaderboard(boardName)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"leaderboard": leaderboard,
	})
}

// UpdateLeaderboard godoc
//	@Summary		Update a leaderboard
//	@Description	Update a leaderboard
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name} [post]
func UpdateLeaderboard(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)
	boardName := c.Params("name")
	var entry models.LeaderboardEntry
	if err := c.BodyParser(&entry); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	err := services.NewLeaderboardService(leaderboardRepo).UpdateLeaderboard(boardName, entry)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"leaderboard": entry,
	})
}

// GetTopNUsers godoc
//	@Summary		Get top N users
//	@Description	Get top N users
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Param			n		path		int		true	"Number of users"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name}/top/{n} [get]
func GetTopNUsers(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)
	boardName := c.Params("name")
	n, err := strconv.Atoi(c.Params("n"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "Invalid input"})
	}

	users, err := services.NewLeaderboardService(leaderboardRepo).GetTopNUsers(boardName, n)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"users": users,
	})
}

// GetUserRankAndScore godoc
//	@Summary		Get user rank and score
//	@Description	Get user rank and score
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Param			userID	path		string	true	"User ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name}/user/{userID} [get]
func GetUserRankAndScore(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)
	boardName := c.Params("name")
	userID := c.Params("userID")

	rank, score, err := services.NewLeaderboardService(leaderboardRepo).GetUserRankAndScore(boardName, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"rank":  rank,
		"score": score,
	})
}

// RemoveUser godoc
//	@Summary		Remove user
//	@Description	Remove user
//	@Tags			leaderboard
//	@Accept			json
//	@Produce		json
//	@Param			name	path		string	true	"Leaderboard name"
//	@Param			userID	path		string	true	"User ID"
//	@Success		200		{object}	string
//	@Failure		400		{object}	string
//	@Failure		500		{object}	string
//	@Router			/leaderboard/{name}/user/{userID} [delete]
func RemoveUser(c *fiber.Ctx) error {
	leaderboardRepo := repositories.NewLeaderboardRepository(utils.RedisClient)
	boardName := c.Params("name")
	userID := c.Params("userID")

	err := services.NewLeaderboardService(leaderboardRepo).RemoveUser(boardName, userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.Status(http.StatusOK).JSON(fiber.Map{
		"message": "User removed successfully",
	})
}
