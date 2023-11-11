package model

type SpaceCraft struct {
	BaseModel
	Id        int     `db:"id"`
	Name      string  `db:"name"`
	Class     string  `db:"class"`
	Crew      int     `db:"crew"`
	Image     string  `db:"image"`
	Value     float64 `db:"value"`
	Status    string  `db:"status"`
	Armaments []SpaceCraftArmamentEmbed
}

type SpaceCraftArmamentEmbed struct {
	ArmamentId int    `db:"armament_id"`
	Title      string `db:"title"`
	Quantity   int    `db:"qty"`
}

type SpaceCraftUpdate struct {
	Name   *string
	Class  *string
	Crew   *int
	Image  *string
	Value  *float64
	Status *string
}

type SpaceCraftFetchFilter struct {
}
