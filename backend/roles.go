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
}

var allRoles []Role

func initRoles() {
	allRoles = nil
	allRoles = append(allRoles, VillagerRole{Dummy: "Test"}, MafiaRole{})
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

const ROLE_VILLAGER = 1
const ROLE_MAFIA = 2
const ROLE_DETECTIVE = 3
