package main

import (
	"fmt"
	"os"

	"RankEdge/internal/controllers"
	"RankEdge/internal/middlewares"
	"RankEdge/internal/utils"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/session"
	"github.com/joho/godotenv"
	fiberSwagger "github.com/swaggo/fiber-swagger"
	_ "RankEdge/docs"
)

func init() {
	godotenv.Load()
	utils.ConnectToRedis()
	utils.ConnectToDb()
	utils.Migrate()
}

// @title			RankEdge API
// @version		1.0
// @description	This is a sample server for GoCommerce.

// @contact.name	Numan Burak Fidan
// @contact.url	https://numanburakfidan.com
// @contact.email	numanburakfidan@yandex.com

// @host		localhost:8080
// @BasePath	/
func main() {
	app := fiber.New()
	store := session.New()

	app.Get("/docs/*", fiberSwagger.WrapHandler)
	authRoutes := app.Group("/auth")
	authRoutes.Post("/signup", controllers.SignUp)
	authRoutes.Post("/signin", controllers.SignIn)
	authRoutes.Get("/signout", controllers.SignOut)

	leaderboardRoutes := app.Group("/leaderboard", func(c *fiber.Ctx) error {
		return middlewares.AuthMiddleware(c, store)
	})
	leaderboardRoutes.Put("/:name", controllers.CreateLeaderBoard)
	leaderboardRoutes.Get("/", controllers.GetLeaderboards)
	leaderboardRoutes.Get("/:name", controllers.GetLeaderboard)
	leaderboardRoutes.Post("/:name", controllers.UpdateLeaderboard)
	leaderboardRoutes.Get("/:name/top/:n", controllers.GetTopNUsers)
	leaderboardRoutes.Get("/:name/user/:userID", controllers.GetUserRankAndScore)
	leaderboardRoutes.Delete("/:name/user/:userID", controllers.RemoveUser)

	app.Listen(fmt.Sprintf(":%s", os.Getenv("HTTP_PORT")))
}
