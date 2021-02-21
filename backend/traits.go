package main

import (
	"math/rand"
	"strings"
)

type Trait interface {
	// Function that gets called when a user tries to send a message, returning modified message
	messageConvert(message string, myself *MafiaUser) string

	//Function that gets called when user dies to a vote (Day vote, Mafia vote, etc.)
	onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser, phaseMod int)

	//Function that gets called when user dies to a player action (e.g. The bomb)
	onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser, phaseMod int)

	//Function that says this trait should be carried between nights
	shouldKeepOnNightChange() bool
}

type BombTrait struct{}

func (bomb BombTrait) messageConvert(message string, myself *MafiaUser) string {
	return message
}

func (bomb BombTrait) onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser, phaseMod int) {
	voters := barrier.getVoters()
	if len(voters) > 0 {
		rand.Shuffle(len(voters), func(i, j int) { voters[i], voters[j] = voters[j], voters[i] })
		killPlayerByNameFromPlayerAction(voters[0], myself, phaseMod)
	}
}

func (bomb BombTrait) onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser, phaseMod int) {
	killPlayerByNameFromPlayerAction(killer.Character.Name, myself, phaseMod)
}

func (bomb BombTrait) shouldKeepOnNightChange() bool {
	return true
}

type CultistLeaderTrait struct{}

func (cultistLeader CultistLeaderTrait) messageConvert(message string, myself *MafiaUser) string {
	return message
}

func (cultistLeader CultistLeaderTrait) onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser, phaseMod int) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Alive && mafiaUsers[i].Role.getTeam() == TEAM_CULTIST {
			killPlayerByNameFromPlayerAction(mafiaUsers[i].Character.Name, myself, phaseMod)
		}
	}
}

func (cultistLeader CultistLeaderTrait) onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser, phaseMod int) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Alive && mafiaUsers[i].Role.getTeam() == TEAM_CULTIST {
			killPlayerByNameFromPlayerAction(mafiaUsers[i].Character.Name, myself, phaseMod)
		}
	}
}

func (cultistLeader CultistLeaderTrait) shouldKeepOnNightChange() bool {
	return true
}

type VoodooCursedTrait struct {
	Word         string
	VoodooPlayer *MafiaUser
}

func (curse VoodooCursedTrait) messageConvert(message string, myself *MafiaUser) string {
	if strings.Contains(strings.ToLower(message), strings.ToLower(curse.Word)) {
		killPlayerByNameFromPlayerAction(myself.Character.Name, curse.VoodooPlayer, 0)
	}
	return message
}

func (curse VoodooCursedTrait) onDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser, phaseMod int) {
}

func (curse VoodooCursedTrait) onDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser, phaseMod int) {
}

func (curse VoodooCursedTrait) shouldKeepOnNightChange() bool {
	return false
}
