package main

import "log"

type Role interface {
	getRoleID() int
	getName() string
	getDescription() string
	getIcon() string
	getTeam() int
	getVotingBarriers() []VotingBarrierInterface
	getNightChatGroup() int
	initialize()
	getSelfShowRoleID() int
}

var allRoles []Role

func initRoles() {
	allRoles = nil
	allRoles = append(allRoles, VillagerRole{}, MafiaRole{}, DetectiveRole{}, BombRole{}, CultistRole{}, VoodooRole{}, JesterRole{}, ParanoidDetectiveRole{})
}

func getRole(id int) Role {
	for _, role := range allRoles {
		if role.getRoleID() == id {
			return role
		}
	}
	log.Printf("Failed to find role id %d\n", id)
	return nil
}

const TEAM_VILLAGER = 1
const TEAM_MAFIA = 2
const TEAM_CULTIST = 3
const TEAM_JESTER = 4

const ROLE_VILLAGER = 1
const ROLE_MAFIA = 2
const ROLE_DETECTIVE = 3
const ROLE_BOMB = 4
const ROLE_CULTIST = 5
const ROLE_VOODOO = 6
const ROLE_JESTER = 7
const ROLE_PARANOID_DETECTIVE = 8
