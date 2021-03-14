package shared

import (
	"arcade-website/character"
)

type MafiaUser struct {
	Name      string
	Alive     bool
	Role      Role
	Character character.Character
	Will      string
	Traits    []Trait
}

var MafiaUsers []MafiaUser = make([]MafiaUser, 0)
