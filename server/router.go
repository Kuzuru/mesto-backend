package server

import (
	"context"

	"mesto/internal/user"
	psql "mesto/internal/user/db"
	"mesto/pkg/db/postgresql"

	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

func routePatchUserInfo(app fiber.Router, repo user.Storage) {
	app.Patch("/users/me", func(c *fiber.Ctx) error {
		// Checking UUID validity
		_, err := uuid.Parse(c.Get("Authorization"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		u, err := repo.FindOne(context.TODO(), c.Get("Authorization"))
		if err != nil {
			if u == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid token",
				})
			}

			log.Error().Stack().Msgf("[PSQL.FindOne] Something went wrong: %+v", err)

			return c.SendStatus(500)
		}

		if err := c.BodyParser(u); err != nil {
			log.Error().Msgf("Couldn't parse user info: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to parse user info",
			})
		}

		u.AuthID = c.Get("Authorization")

		err = repo.UpdateProfile(context.TODO(), *u)
		if err != nil {
			log.Error().Msgf("Couldn't update user info: %v", err)
			return c.Status(500).JSON(fiber.Map{
				"message": "Failed to update user info",
			})
		}

		return c.JSON(fiber.Map{
			"name":  u.Name,
			"about": u.About,
		})
	})
}

func routeGetUserInfo(app fiber.Router, repo user.Storage) {
	app.Get("/users/me", func(c *fiber.Ctx) error {
		// Checking UUID validity
		_, err := uuid.Parse(c.Get("Authorization"))
		if err != nil {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "Invalid token",
			})
		}

		u, err := repo.FindOne(context.TODO(), c.Get("Authorization"))
		if err != nil {
			if u == nil {
				return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
					"message": "Invalid token",
				})
			}

			log.Error().Stack().Msgf("[PSQL.FindOne] Something went wrong: %+v", err)

			return c.SendStatus(500)
		}

		return c.JSON(fiber.Map{
			"id":     u.ID,
			"name":   u.Name,
			"about":  u.About,
			"avatar": u.Avatar,
		})
	})
}

func RegisterHTTPEndpoints(app *fiber.App) {
	psqlClient, err := postgresql.NewClient(context.TODO(), 3)
	if err != nil {
		log.Error().Stack().Msgf("%v", err)
	}

	repository := psql.NewRepository(psqlClient)

	// Using grouping for versioning
	v1 := app.Group("/v1")

	routeGetUserInfo(v1, repository)
	routePatchUserInfo(v1, repository)
}
