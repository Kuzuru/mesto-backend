package server

import (
	"os"

	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/swagger"
	"github.com/rs/zerolog/log"
)

func Run() {
	app := fiber.New(fiber.Config{
		Prefork:       true,
		CaseSensitive: true,
		JSONEncoder:   sonic.Marshal,
		JSONDecoder:   sonic.Unmarshal,
	})

	app.Use(cors.New())

	address := ":" + os.Getenv("PORT")

	// Registering Swagger for API documentation
	app.Get("/swagger/*", swagger.HandlerDefault)

	// Registering all available HTTP endpoints
	RegisterHTTPEndpoints(app)

	go func() {
		if err := app.Listen(address); err == nil {
			log.Fatal().Stack().AnErr("app.Listen: %s", err)
		}
	}()

	if !fiber.IsChild() {
		log.Info().Msgf("Server listening on %s", address)
	}
}
