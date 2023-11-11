package controller

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/swagger"
	_ "github.com/mariojuzar/fleet-inventory/docs"
	"github.com/mariojuzar/fleet-inventory/internal/usecases"
	"time"
)

type API struct {
	readTimeout   time.Duration
	writeTimeout  time.Duration
	enableSwagger bool
	scUc          usecases.SpaceCraftUseCases
}

type Options struct {
	ReadTimeout   time.Duration
	WriteTimeout  time.Duration
	EnableSwagger bool
	ScUc          usecases.SpaceCraftUseCases
}

func New(opts *Options) *API {
	return &API{
		readTimeout:   opts.ReadTimeout,
		writeTimeout:  opts.WriteTimeout,
		enableSwagger: opts.EnableSwagger,
		scUc:          opts.ScUc,
	}
}

func (api *API) RegisterRoute() *fiber.App {
	config := fiber.Config{
		ReadTimeout:           api.readTimeout,
		WriteTimeout:          api.writeTimeout,
		DisableStartupMessage: true,
	}

	app := fiber.New(config)

	if api.enableSwagger {
		app.Get("/docs/*", swagger.HandlerDefault)
	}

	// health check
	app.Get("/ping", api.HealthCheck)

	appV1 := app.Group("/v1")

	spaceCraft := appV1.Group("/space-craft")
	spaceCraft.Post("/", api.CreateSpaceCraft)
	spaceCraft.Put("/", api.EditSpaceCraft)
	spaceCraft.Get("/:id", api.GetSpaceCraft)
	spaceCraft.Get("/", api.FetchSpaceCraft)
	spaceCraft.Delete("/", api.DeleteSpaceCraft)

	return app
}
