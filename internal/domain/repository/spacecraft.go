package repository

import (
	"context"
	"github.com/mariojuzar/fleet-inventory/internal/domain/model"
)

type SpaceCraftRepository interface {
	Insert(ctx context.Context, model *model.SpaceCraft) error
	Update(ctx context.Context, model *model.SpaceCraft) error
	Get(ctx context.Context, id int) (*model.SpaceCraft, error)
	Delete(ctx context.Context, id int) error
	Fetch(ctx context.Context, filter *model.SpaceCraftFetchFilter) ([]model.SpaceCraft, error)
}
