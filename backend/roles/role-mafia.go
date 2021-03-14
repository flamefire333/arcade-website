package roles

import (
	"arcade-website/actions"
	"arcade-website/shared"
)

type MafiaRole struct{}

func (mafia MafiaRole) GetRoleID() int {
	return shared.ROLE_MAFIA
}

func (mafia MafiaRole) GetName() string {
	return "mafia"
}

func (mafia MafiaRole) GetDescription() string {
	return "Mafia can vote to kill someone as a night action. Mafia are sided with the Mafia."
}

func (mafia MafiaRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/766499535899459584.png?v=1"
}

func (mafia MafiaRole) GetTeam() int {
	return shared.TEAM_MAFIA
}

func (mafia MafiaRole) GetNightChatGroup() int {
	return shared.CHAT_MAFIA
}

func (mafia MafiaRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	barrierIDs := make([]int, 0)
	barrierID := shared.GetNextVotingBarrierID()
	barrierIDs = append(barrierIDs, barrierID)
	fields := make([]shared.VoteField, 0)
	fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
	barriers = append(barriers, MafiaVotingBarrier{Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}})
	return barriers
}

func (mafia MafiaRole) Initialize() {}

func (mafia MafiaRole) GetSelfShowRoleID() int {
	return shared.ROLE_MAFIA
}

type MafiaVotingBarrier struct {
	Base shared.VotingBarrierBase
}

func (mvb MafiaVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Role.GetTeam() == shared.TEAM_MAFIA && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (mvb MafiaVotingBarrier) GetTitle() string {
	return "Maifa Vote"
}

func (mvb MafiaVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (mvb MafiaVotingBarrier) ExecuteOption(option []string) {
	actions.KillPlayerByNameFromVote(option[0], mvb, 1)
}

func (mvb MafiaVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &mvb.Base
}
