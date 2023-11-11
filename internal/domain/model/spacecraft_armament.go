package model

type SpaceCraftArmament struct {
	BaseModel
	Id           int `db:"id"`
	SpaceCraftId int `db:"space_craft_id"`
	ArmamentId   int `db:"armament_id"`
	Quantity     int `db:"qty"`
}
