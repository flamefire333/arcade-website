package main

type DetectiveRole struct{}

func (detective DetectiveRole) getRoleID() int {
	return ROLE_DETECTIVE
}

func (detective DetectiveRole) getName() string {
	return "detective"
}

func (detective DetectiveRole) getDescription() string {
	return "Detective can check the alignment of a player as a night action. Detectives are sided with the Village."
}

func (detective DetectiveRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/700748202931912814.png?v=1"
}

func (detective DetectiveRole) getTeam() int {
	return TEAM_VILLAGER
}

func (detective DetectiveRole) getNightChatGroup() int {
	return CHAT_NOT_ALLOWED
}

func (detective DetectiveRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getRoleID() == ROLE_DETECTIVE {
			barrierIDs := make([]int, 0)
			barrierID := GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID)
			fields := make([]VoteField, 0)
			fields = append(fields, VoteField{Type: "option", Options: getStandardVotingOptions(), BarrierID: barrierID})
			barriers = append(barriers, DetectiveVotingBarrier{Base: VotingBarrierBase{Votes: getStandardBaseVotes(barrierIDs), Fields: fields}, Name: mafiaUsers[i].Character.Name})
		}
	}
	return barriers
}

func (detective DetectiveRole) initialize() {}

func (detective DetectiveRole) getSelfShowRoleID() int {
	return ROLE_DETECTIVE
}

type DetectiveVotingBarrier struct {
	Base VotingBarrierBase
	Name string
}

func (dvb DetectiveVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Character.Name == dvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (mvb DetectiveVotingBarrier) getTitle() string {
	return "Detective Investigation"
}

func (dvb DetectiveVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb DetectiveVotingBarrier) executeOption(option []string) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Character.Name == option[0] {
			state := "weird"
			if mafiaUsers[i].Role.getTeam() == TEAM_MAFIA {
				state = "shady"
			} else if mafiaUsers[i].Role.getTeam() == TEAM_VILLAGER {
				state = "good"
			}
			sendPrivateInfoMessage(option[0]+" seems to be "+state+"!", CHAT_ALL, 1, dvb.Name)
		}
	}
}

func (dvb DetectiveVotingBarrier) getBase() *VotingBarrierBase {
	return &dvb.Base
}
