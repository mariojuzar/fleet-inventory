package spacecraft

import (
	"github.com/mariojuzar/fleet-inventory/internal/domain/repository"
	"github.com/mariojuzar/fleet-inventory/internal/usecases"
)

type Module struct {
	spaceCraftRepo repository.SpaceCraftRepository
}

type Opts struct {
	SpaceCraftRepo repository.SpaceCraftRepository
}

func New(opts *Opts) usecases.SpaceCraftUseCases {
	return &Module{
		spaceCraftRepo: opts.SpaceCraftRepo,
	}
}
