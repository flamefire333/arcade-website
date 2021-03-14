package roles

import (
	"arcade-website/actions"
	"arcade-website/shared"
)

type DetectiveRole struct{}

func (detective DetectiveRole) GetRoleID() int {
	return shared.ROLE_DETECTIVE
}

func (detective DetectiveRole) GetName() string {
	return "detective"
}

func (detective DetectiveRole) GetDescription() string {
	return "Detective can check the alignment of a player as a night action. Detectives are sided with the Village."
}

func (detective DetectiveRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/700748202931912814.png?v=1"
}

func (detective DetectiveRole) GetTeam() int {
	return shared.TEAM_VILLAGER
}

func (detective DetectiveRole) GetNightChatGroup() int {
	return shared.CHAT_NOT_ALLOWED
}

func (detective DetectiveRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetRoleID() == shared.ROLE_DETECTIVE {
			barrierIDs := make([]int, 0)
			barrierID := shared.GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID)
			fields := make([]shared.VoteField, 0)
			fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
			barriers = append(barriers, DetectiveVotingBarrier{Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}, Name: shared.MafiaUsers[i].Character.Name})
		}
	}
	return barriers
}

func (detective DetectiveRole) Initialize() {}

func (detective DetectiveRole) GetSelfShowRoleID() int {
	return shared.ROLE_DETECTIVE
}

type DetectiveVotingBarrier struct {
	Base shared.VotingBarrierBase
	Name string
}

func (dvb DetectiveVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Character.Name == dvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (mvb DetectiveVotingBarrier) GetTitle() string {
	return "Detective Investigation"
}

func (dvb DetectiveVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb DetectiveVotingBarrier) ExecuteOption(option []string) {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Character.Name == option[0] {
			state := "weird"
			if shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_MAFIA {
				state = "shady"
			} else if shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_VILLAGER {
				state = "good"
			}
			actions.SendPrivateInfoMessage(option[0]+" seems to be "+state+"!", shared.CHAT_ALL, 1, dvb.Name)
		}
	}
}

func (dvb DetectiveVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &dvb.Base
}
