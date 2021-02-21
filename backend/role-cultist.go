package main

import "math/rand"

type CultistRole struct {
}

func (cultist CultistRole) getRoleID() int {
	return ROLE_CULTIST
}

func (cultist CultistRole) getName() string {
	return "cultist"
}

func (cultist CultistRole) getDescription() string {
	return "The cultist can vote to convert someone into a cultist. All cultist die if the cultist leader dies."
}

func (cultist CultistRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/796875834832322600.png?v=1"
}

func (cultist CultistRole) getTeam() int {
	return TEAM_CULTIST
}

func (cultist CultistRole) getNightChatGroup() int {
	return CHAT_CULTIST
}

func (cultist CultistRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	barrierIDs := make([]int, 0)
	barrierID := GetNextVotingBarrierID()
	barrierIDs = append(barrierIDs, barrierID)
	fields := make([]VoteField, 0)
	fields = append(fields, VoteField{Type: "option", Options: getStandardVotingOptions(), BarrierID: barrierID})
	barriers = append(barriers, CultistVotingBarrier{Base: VotingBarrierBase{Votes: getStandardBaseVotes(barrierIDs), Fields: fields}})
	return barriers
}

func (cultist CultistRole) initialize() {
	cultistIDs := make([]int, 0)
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getTeam() == TEAM_CULTIST {
			cultistIDs = append(cultistIDs, i)
		}
	}
	rand.Shuffle(len(cultistIDs), func(i, j int) { cultistIDs[i], cultistIDs[j] = cultistIDs[j], cultistIDs[i] })
	mafiaUsers[cultistIDs[0]].Traits = append(mafiaUsers[cultistIDs[0]].Traits, CultistLeaderTrait{})
	sendPrivateInfoMessage("You are the cultist leader!", CHAT_ALL, 1, mafiaUsers[cultistIDs[0]].Character.Name)
}

func (cultist CultistRole) getSelfShowRoleID() int {
	return ROLE_CULTIST
}

type CultistVotingBarrier struct {
	Base VotingBarrierBase
}

func (cvb CultistVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Role.getTeam() == TEAM_CULTIST && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (cvb CultistVotingBarrier) getTitle() string {
	return "Cultist Vote"
}

func (cvb CultistVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (cvb CultistVotingBarrier) executeOption(option []string) {
	setPlayerRoleByNameFromVote(option[0], ROLE_CULTIST, cvb)
}

func (cvb CultistVotingBarrier) getBase() *VotingBarrierBase {
	return &cvb.Base
}
