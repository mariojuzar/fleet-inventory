package controller

import (
	"encoding/json"
	"github.com/gofiber/fiber/v2"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/request"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/response"
	"strconv"
)

// CreateSpaceCraft godoc
// @Summary 	CreateSpaceCraft
// @Description	CreateSpaceCraft
// @Tags		space-craft
// @Accept		json
// @Param		space-craft		body		request.SpaceShipCreateRequest	true	"Create payload"
// @Produce		json
// @Success		200	{object}	JsonResponse[bool]{}
// @Router		/v1/space-craft	[post]
func (api *API) CreateSpaceCraft(ctx *fiber.Ctx) error {
	req := &request.SpaceShipCreateRequest{}
	err := json.Unmarshal(ctx.Body(), req)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	err = api.scUc.Create(ctx.UserContext(), req)
	if err != nil {
		return customErrorResponse(ctx, err)
	}

	return ctx.JSON(JsonResponse[any]{Success: true})
}

// EditSpaceCraft godoc
// @Summary 	EditSpaceCraft
// @Description	EditSpaceCraft
// @Tags		space-craft
// @Accept		json
// @Param 		id		 path 		string 	true	"space-craft id"
// @Param		space-craft		body		request.SpaceShipEditRequest	true	"Edit payload"
// @Produce		json
// @Success		200	{object}	JsonResponse[bool]{}
// @Router		/v1/space-craft/{id}	[put]
func (api *API) EditSpaceCraft(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return customErrorResponse(ctx, err)
	}

	req := &request.SpaceShipEditRequest{}
	err = json.Unmarshal(ctx.Body(), req)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	err = api.scUc.Edit(ctx.UserContext(), id, req)
	if err != nil {
		return customErrorResponse(ctx, err)
	}

	return ctx.JSON(JsonResponse[any]{Success: true})
}

// GetSpaceCraft godoc
// @Summary 	GetSpaceCraft
// @Description	GetSpaceCraft
// @Tags		space-craft
// @Accept		json
// @Param 		id		 path 		string 	true	"space-craft id"
// @Produce		json
// @Success		200	{object}	JsonResponse[response.SpaceCraftResponse]{}
// @Router		/v1/space-craft/{id}	[get]
func (api *API) GetSpaceCraft(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	res, err := api.scUc.Get(ctx.UserContext(), id)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	return ctx.JSON(JsonResponse[response.SpaceCraftResponse]{
		Data: *res,
	})
}

// FetchSpaceCraft godoc
// @Summary 	FetchSpaceCraft
// @Description	FetchSpaceCraft
// @Tags		space-craft
// @Accept		json
// @Param 		name		query 		string 	false 	"space-craft name"
// @Param 		status		query 		string 	false 	"space-craft status"
// @Param 		class		query 		string 	false 	"space-craft class"
// @Produce		json
// @Success		200	{object}	JsonResponse[bool]{}
// @Router		/v1/space-craft	[get]
func (api *API) FetchSpaceCraft(ctx *fiber.Ctx) error {
	return nil
}

// DeleteSpaceCraft godoc
// @Summary 	DeleteSpaceCraft
// @Description	DeleteSpaceCraft
// @Tags		space-craft
// @Accept		json
// @Param 		id		 path 		string 	true	"space-craft id"
// @Produce		json
// @Success		200	{object}	JsonResponse[bool]{}
// @Router		/v1/space-craft/{id}	[delete]
func (api *API) DeleteSpaceCraft(ctx *fiber.Ctx) error {
	idParam := ctx.Params("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	err = api.scUc.Delete(ctx.UserContext(), id)
	if err != nil {
		return customErrorResponse(ctx, err)
	}
	return ctx.JSON(JsonResponse[any]{Success: true})
}
