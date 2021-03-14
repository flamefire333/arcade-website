package traits

import (
	"arcade-website/actions"
	"arcade-website/shared"
	"arcade-website/voting"
	"math/rand"
	"strings"
)

type BombTrait struct{}

func (bomb BombTrait) MessageConvert(message string, myself *shared.MafiaUser) string {
	return message
}

func (bomb BombTrait) OnDeathByVote(barrier shared.VotingBarrierInterface, myself *shared.MafiaUser, phaseMod int) {
	voters := barrier.GetVoters()
	if len(voters) > 0 {
		rand.Shuffle(len(voters), func(i, j int) { voters[i], voters[j] = voters[j], voters[i] })
		actions.KillPlayerByNameFromPlayerAction(voters[0], myself, phaseMod)
	}
}

func (bomb BombTrait) OnDeathByPlayerAction(killer *shared.MafiaUser, myself *shared.MafiaUser, phaseMod int) {
	actions.KillPlayerByNameFromPlayerAction(killer.Character.Name, myself, phaseMod)
}

func (bomb BombTrait) ShouldKeepOnNightChange() bool {
	return true
}

type CultistLeaderTrait struct{}

func (cultistLeader CultistLeaderTrait) MessageConvert(message string, myself *shared.MafiaUser) string {
	return message
}

func (cultistLeader CultistLeaderTrait) OnDeathByVote(barrier shared.VotingBarrierInterface, myself *shared.MafiaUser, phaseMod int) {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Alive && shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_CULTIST {
			actions.KillPlayerByNameFromPlayerAction(shared.MafiaUsers[i].Character.Name, myself, phaseMod)
		}
	}
}

func (cultistLeader CultistLeaderTrait) OnDeathByPlayerAction(killer *shared.MafiaUser, myself *shared.MafiaUser, phaseMod int) {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Alive && shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_CULTIST {
			actions.KillPlayerByNameFromPlayerAction(shared.MafiaUsers[i].Character.Name, myself, phaseMod)
		}
	}
}

func (cultistLeader CultistLeaderTrait) ShouldKeepOnNightChange() bool {
	return true
}

type VoodooCursedTrait struct {
	Word         string
	VoodooPlayer *shared.MafiaUser
}

func (curse VoodooCursedTrait) MessageConvert(message string, myself *shared.MafiaUser) string {
	if strings.Contains(strings.ToLower(message), strings.ToLower(curse.Word)) {
		actions.KillPlayerByNameFromPlayerAction(myself.Character.Name, curse.VoodooPlayer, 0)
	}
	return message
}

func (curse VoodooCursedTrait) OnDeathByVote(barrier shared.VotingBarrierInterface, myself *shared.MafiaUser, phaseMod int) {
}

func (curse VoodooCursedTrait) OnDeathByPlayerAction(killer *shared.MafiaUser, myself *shared.MafiaUser, phaseMod int) {
}

func (curse VoodooCursedTrait) ShouldKeepOnNightChange() bool {
	return false
}

type JesterTrait struct{}

func (jester JesterTrait) MessageConvert(message string, myself *shared.MafiaUser) string {
	return message
}

func (jester JesterTrait) OnDeathByVote(barrier shared.VotingBarrierInterface, myself *shared.MafiaUser, phaseMod int) {
	switch barrier.(type) {
	case voting.DayVotingBarrier:
		actions.EndGame("The jester " + myself.Character.Name + " has won!")
	}
}

func (jester JesterTrait) OnDeathByPlayerAction(killer *shared.MafiaUser, myself *shared.MafiaUser, phaseMod int) {
}

func (jester JesterTrait) ShouldKeepOnNightChange() bool {
	return true
}
