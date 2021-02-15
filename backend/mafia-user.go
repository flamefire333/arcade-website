package main

type MafiaUser struct {
	Name      string
	Alive     bool
	Role      Role
	Character Character
	Will      string
	Traits    []Trait
}

func generateMafiaUser(name string, roleToUse Role, character Character) MafiaUser {
	return MafiaUser{Name: name, Role: roleToUse, Character: character, Alive: true, Will: "", Traits: make([]Trait, 0)}
}
