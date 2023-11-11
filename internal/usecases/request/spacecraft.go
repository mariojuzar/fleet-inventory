package request

import (
	"github.com/mariojuzar/fleet-inventory/internal/domain/model"
	"time"
)

type SpaceShipCreateRequest struct {
	Name      string              `json:"name"`
	Class     string              `json:"class"`
	Crew      int                 `json:"crew"`
	Image     string              `json:"image"`
	Value     float64             `json:"value"`
	Status    string              `json:"status"`
	Armaments []SpaceShipArmament `json:"armaments"`
}

type SpaceShipArmament struct {
	ArmamentId int `json:"armament_id"`
	Quantity   int `json:"quantity"`
}

type SpaceShipEditRequest struct {
	Name   string  `json:"name"`
	Class  string  `json:"class"`
	Crew   int     `json:"crew"`
	Image  string  `json:"image"`
	Value  float64 `json:"value"`
	Status string  `json:"status"`
}

type SpaceShipFetchRequest struct {
	Name   string `json:"-"`
	Class  string `json:"-"`
	Status string `json:"-"`
}

func (c *SpaceShipCreateRequest) ToModel() *model.SpaceCraft {
	now := time.Now()

	craft := &model.SpaceCraft{
		BaseModel: model.BaseModel{
			CreatedAt: now,
			UpdatedAt: now,
		},
		Name:   c.Name,
		Class:  c.Class,
		Crew:   c.Crew,
		Image:  c.Image,
		Value:  c.Value,
		Status: c.Status,
	}

	for _, armament := range c.Armaments {
		craft.Armaments = append(craft.Armaments, model.SpaceCraftArmamentEmbed{
			ArmamentId: armament.ArmamentId,
			Quantity:   armament.Quantity,
		})
	}

	return craft
}

func (c *SpaceShipEditRequest) ToModelUpdate() *model.SpaceCraftUpdate {
	update := &model.SpaceCraftUpdate{
		Name:   &c.Name,
		Class:  &c.Class,
		Crew:   &c.Crew,
		Image:  &c.Image,
		Value:  &c.Value,
		Status: &c.Status,
	}

	return update
}
