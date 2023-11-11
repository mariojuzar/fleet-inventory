package spacecraft

import (
	"context"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/request"
	"github.com/mariojuzar/fleet-inventory/internal/usecases/response"
	"github.com/rs/zerolog/log"
)

func (m *Module) Create(ctx context.Context, req *request.SpaceShipCreateRequest) error {
	craft := req.ToModel()

	err := m.spaceCraftRepo.Insert(ctx, craft)
	if err != nil {
		log.Err(err).Msg("Failed to insert space ship")
		return err
	}

	return nil
}

func (m *Module) Edit(ctx context.Context, req *request.SpaceShipEditRequest) error {
	//TODO implement me
	panic("implement me")
}

func (m *Module) Get(ctx context.Context, id int) (*response.SpaceCraftResponse, error) {
	spaceCraft, err := m.spaceCraftRepo.Get(ctx, id)
	if err != nil {
		return nil, err
	}
	resp := &response.SpaceCraftResponse{}
	resp.FromModel(spaceCraft)
	return resp, nil
}

func (m *Module) Fetch(ctx context.Context, req *request.SpaceShipFetchRequest) ([]response.SpaceCraftFetchResponse, error) {
	//TODO implement me
	panic("implement me")
}

func (m *Module) Delete(ctx context.Context, id int) error {
	//TODO implement me
	panic("implement me")
}