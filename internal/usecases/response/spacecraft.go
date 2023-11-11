package response

import "github.com/mariojuzar/fleet-inventory/internal/domain/model"

type SpaceCraftResponse struct {
	Id        int                          `json:"id,omitempty"`
	Name      string                       `json:"name,omitempty"`
	Class     string                       `json:"class,omitempty"`
	Crew      int                          `json:"crew,omitempty"`
	Image     string                       `json:"image,omitempty"`
	Value     float64                      `json:"valued,omitempty"`
	Status    string                       `json:"status,omitempty"`
	Armaments []SpaceCraftArmamentResponse `json:"armament,omitempty"`
}

type SpaceCraftArmamentResponse struct {
	Title    string `json:"title,omitempty"`
	Quantity int    `json:"qty,omitempty"`
}

type SpaceCraftFetchResponse struct {
	Id     int    `json:"id,omitempty"`
	Name   string `json:"name,omitempty"`
	Status string `json:"status,omitempty"`
	Class  string `json:"class,omitempty"`
}

func (s *SpaceCraftResponse) FromModel(m *model.SpaceCraft) {
	s.Id = m.Id
	s.Name = m.Name
	s.Status = m.Status
	s.Value = m.Value
	s.Image = m.Image
	s.Crew = m.Crew
	s.Class = m.Class

	for _, armament := range m.Armaments {
		s.Armaments = append(s.Armaments, SpaceCraftArmamentResponse{
			Title:    armament.Title,
			Quantity: armament.Quantity,
		})
	}
}
