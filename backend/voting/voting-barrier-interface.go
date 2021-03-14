package voting

import (
	"arcade-website/actions"
	"arcade-website/shared"
)

type DayVotingBarrier struct {
	Base shared.VotingBarrierBase
}

func (dvb DayVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (dvb DayVotingBarrier) GetTitle() string {
	return "Day Vote"
}

func (dvb DayVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb DayVotingBarrier) ExecuteOption(option []string) {
	actions.KillPlayerByNameFromVote(option[0], dvb, 1)
}

func (dvb DayVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &dvb.Base
}
