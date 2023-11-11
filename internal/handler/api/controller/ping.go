package controller

import "github.com/gofiber/fiber/v2"

// HealthCheck godoc
// @Summary 	HealthCheck
// @Description	HealthCheck
// @Tags		health-check
// @Accept		json
// @Produce		json
// @Success		200	{object}	JsonResponse[bool]{}
// @Router		/ping	[get]
func (api *API) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.JSON("pong")
}
