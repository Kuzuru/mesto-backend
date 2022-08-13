package server

import (
	"context"

	"mesto/internal/user"
	psql "mesto/internal/user/db"
	"mesto/pkg/db/postgresql"

	"github.com/gofiber/fiber/v2"
	"github.com/rs/zerolog/log"
)

func userRoutes(app fiber.Router, repo user.Storage) {
	// Get user info
	app.Get("/users/me", func(c *fiber.Ctx) error {
		u, err := repo.FindOne(context.TODO(), c.Get("Authorization"))
		if err != nil {
			log.Fatal().Stack().Msgf("FIND_ALL: %v", err)
			return c.SendStatus(500)
		}

		return c.JSON(fiber.Map{
			"id":     u.ID,
			"name":   u.Name,
			"about":  u.About,
			"avatar": u.Avatar,
		})
	})

	// Update user info
	app.Patch("/users/me", func(c *fiber.Ctx) error {
		var u = new(user.User)

		if err := c.BodyParser(u); err != nil {
			log.Error().Msgf("Couldn't parse user info: %v", err)
			return c.SendStatus(500)
		}

		u.AuthID = c.Get("Authorization")

		err := repo.UpdateProfile(context.TODO(), *u)
		if err != nil {
			log.Error().Msgf("Couldn't update user info: %v", err)
			return c.SendStatus(500)
		}

		return c.JSON(fiber.Map{
			"name":  u.Name,
			"about": u.About,
		})
	})
}

func RegisterHTTPEndpoints(app *fiber.App) {
	psqlClient, err := postgresql.NewClient(context.TODO(), 3)
	if err != nil {
		log.Fatal().Stack().Msgf("%v", err)
	}

	repository := psql.NewRepository(psqlClient)

	// Using grouping for versioning
	v1 := app.Group("/v1")

	userRoutes(v1, repository)
}
