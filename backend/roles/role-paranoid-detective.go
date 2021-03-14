package roles

import (
	"arcade-website/actions"
	"arcade-website/shared"
	"math/rand"
)

type ParanoidDetectiveRole struct{}

func (detective ParanoidDetectiveRole) GetRoleID() int {
	return shared.ROLE_PARANOID_DETECTIVE
}

func (detective ParanoidDetectiveRole) GetName() string {
	return "paranoid detective"
}

func (detective ParanoidDetectiveRole) GetDescription() string {
	return "Paranoid Detective can check the alignment of a player as a night action, but will receive a random result. Paranoid Detectives are sided with the Village."
}

func (detective ParanoidDetectiveRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/796870998778445844.png?v=1"
}

func (detective ParanoidDetectiveRole) GetTeam() int {
	return shared.TEAM_VILLAGER
}

func (detective ParanoidDetectiveRole) GetNightChatGroup() int {
	return shared.CHAT_NOT_ALLOWED
}

func (detective ParanoidDetectiveRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetRoleID() == shared.ROLE_PARANOID_DETECTIVE {
			barrierIDs := make([]int, 0)
			barrierID := shared.GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID)
			fields := make([]shared.VoteField, 0)
			fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
			barriers = append(barriers, ParanoidDetectiveVotingBarrier{Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}, Name: shared.MafiaUsers[i].Character.Name})
		}
	}
	return barriers
}

func (detective ParanoidDetectiveRole) Initialize() {}

func (detective ParanoidDetectiveRole) GetSelfShowRoleID() int {
	return shared.ROLE_DETECTIVE
}

type ParanoidDetectiveVotingBarrier struct {
	Base shared.VotingBarrierBase
	Name string
}

func (dvb ParanoidDetectiveVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Character.Name == dvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (dvb ParanoidDetectiveVotingBarrier) GetTitle() string {
	return "Detective Investigation"
}

func (dvb ParanoidDetectiveVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb ParanoidDetectiveVotingBarrier) ExecuteOption(option []string) {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Character.Name == option[0] {
			options := make([]string, 0)
			options = append(options, "weird", "shady", "good")
			rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
			state := options[0]
			actions.SendPrivateInfoMessage(option[0]+" seems to be "+state+"!", shared.CHAT_ALL, 1, dvb.Name)
		}
	}
}

func (dvb ParanoidDetectiveVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &dvb.Base
}
