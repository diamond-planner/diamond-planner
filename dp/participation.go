package dp

import "github.com/pocketbase/pocketbase/core"

type ParticipationsByType struct {
	In          []*core.Record         `json:"in"`
	Out         []*core.Record         `json:"out"`
	Maybe       []*core.Record         `json:"maybe"`
	Unspecified []UserParticipationDTO `json:"unspecified"`
}

type UserParticipationDTO struct {
	ID        string `json:"id" db:"id"`
	FirstName string `json:"first_name" db:"first_name"`
	LastName  string `json:"last_name" db:"last_name"`
}
