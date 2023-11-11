package usecases

import (
	"context"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/request"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/response"
)

type SpaceCraftUseCases interface {
	Create(ctx context.Context, req *request.SpaceShipCreateRequest) error
	Edit(ctx context.Context, req *request.SpaceShipEditRequest) error
	Get(ctx context.Context, id int) (*response.SpaceCraftResponse, error)
	Fetch(ctx context.Context, req *request.SpaceShipFetchRequest) ([]response.SpaceCraftFetchResponse, error)
	Delete(ctx context.Context, id int) error
}
