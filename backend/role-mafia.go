package main

type MafiaRole struct{}

func (mafia MafiaRole) getRoleID() int {
	return ROLE_MAFIA
}

func (mafia MafiaRole) getName() string {
	return "mafia"
}

func (mafia MafiaRole) getDescription() string {
	return "Mafia can vote to kill someone as a night action. Mafia are sided with the Mafia."
}

func (mafia MafiaRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/766499535899459584.png?v=1"
}

func (mafia MafiaRole) getTeam() int {
	return TEAM_MAFIA
}

func (mafia MafiaRole) getNightChatGroup() int {
	return CHAT_MAFIA
}

func (mafia MafiaRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	emptyMap := make(map[string]string, 0)
	barriers = append(barriers, MafiaVotingBarrier{Base: VotingBarrierBase{Votes: emptyMap, ID: GetNextVotingBarrierID()}})
	return barriers
}

type MafiaVotingBarrier struct {
	Base VotingBarrierBase
}

func (mvb MafiaVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Role.getTeam() == TEAM_MAFIA && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (mvb MafiaVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (mvb MafiaVotingBarrier) executeOption(option string) {
	killPlayerByNameFromVote(option, mvb)
}

func (mvb MafiaVotingBarrier) getBase() *VotingBarrierBase {
	return &mvb.Base
}
