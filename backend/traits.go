package main

import "math/rand"

type Trait interface {
	// Function that gets called when a user tries to send a message, returning modified message
	messageConvert(message string) string

	//Function that gets called when user dies to a vote (Day vote, Mafia vote, etc.)
	onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser)

	//Function that gets called when user dies to a player action (e.g. The bomb)
	onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser)
}

type BombTrait struct{}

func (bomb BombTrait) messageConvert(message string) string {
	return message
}

func (bomb BombTrait) onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser) {
	voters := barrier.getVoters()
	if len(voters) > 0 {
		rand.Shuffle(len(voters), func(i, j int) { voters[i], voters[j] = voters[j], voters[i] })
		killPlayerByNameFromPlayerAction(voters[0], myself)
	}
}

func (bomb BombTrait) onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser) {
	killPlayerByNameFromPlayerAction(killer.Character.Name, myself)
}

type CultistLeaderTrait struct{}

func (cultistLeader CultistLeaderTrait) messageConvert(message string) string {
	return message
}

func (cultistLeader CultistLeaderTrait) onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Alive && mafiaUsers[i].Role.getTeam() == TEAM_CULTIST {
			killPlayerByNameFromPlayerAction(mafiaUsers[i].Character.Name, myself)
		}
	}
}

func (cultistLeader CultistLeaderTrait) onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Alive && mafiaUsers[i].Role.getTeam() == TEAM_CULTIST {
			killPlayerByNameFromPlayerAction(mafiaUsers[i].Character.Name, myself)
		}
	}
}
