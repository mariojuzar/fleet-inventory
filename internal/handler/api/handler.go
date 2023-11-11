package api

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/fleet-inventory/config"
	"github.com/mariojuzar/fleet-inventory/internal/handler/api/controller"
	"github.com/mariojuzar/fleet-inventory/internal/usecases"
	"github.com/rs/zerolog/log"
)

type Options struct {
	Cfg  config.ServerConfig
	ScUc usecases.SpaceCraftUseCases
}

type Handler struct {
	router *fiber.App
	port   uint
}

func New(opts *Options) *Handler {
	return &Handler{
		router: controller.New(&controller.Options{
			ReadTimeout:   opts.Cfg.ReadTimeout,
			WriteTimeout:  opts.Cfg.WriteTimeout,
			EnableSwagger: opts.Cfg.EnableSwagger,
			ScUc:          opts.ScUc,
		}).RegisterRoute(),
		port: opts.Cfg.Port}
}

func (h *Handler) Run() {
	err := h.router.Listen(fmt.Sprintf(":%d", h.port))
	if err != nil {
		log.Fatal().Err(err).Msg("Failed to initiate router")
	}
}
