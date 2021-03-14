package mafiaUser

import (
	"arcade-website/character"
	"arcade-website/shared"
)

func GenerateMafiaUser(name string, roleToUse shared.Role, character character.Character) shared.MafiaUser {
	return shared.MafiaUser{Name: name, Role: roleToUse, Character: character, Alive: true, Will: "", Traits: make([]shared.Trait, 0)}
}
