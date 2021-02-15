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
			emptyMap := make(map[string]string, 0)
			barriers = append(barriers, DetectiveVotingBarrier{Base: VotingBarrierBase{Votes: emptyMap, ID: GetNextVotingBarrierID()}, Name: mafiaUsers[i].Character.Name})
		}
	}
	return barriers
}

func (detective DetectiveRole) initialize() {}

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

func (dvb DetectiveVotingBarrier) executeOption(option string) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Character.Name == option {
			state := "weird"
			if mafiaUsers[i].Role.getTeam() == TEAM_MAFIA {
				state = "shady"
			} else if mafiaUsers[i].Role.getTeam() == TEAM_VILLAGER {
				state = "good"
			}
			sendPrivateInfoMessage(option+" seems to be "+state+"!", CHAT_ALL, 1, dvb.Name)
		}
	}
}

func (dvb DetectiveVotingBarrier) getBase() *VotingBarrierBase {
	return &dvb.Base
}
