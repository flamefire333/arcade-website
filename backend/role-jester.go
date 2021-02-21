package main

type JesterRole struct {
}

func (jester JesterRole) getRoleID() int {
	return ROLE_JESTER
}

func (jester JesterRole) getName() string {
	return "jester"
}

func (jester JesterRole) getDescription() string {
	return "The jester is trying to get voted off in the day vote. The jester is on its own team."
}

func (jester JesterRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/656062178159493135.png?v=1"
}

func (jester JesterRole) getTeam() int {
	return TEAM_JESTER
}

func (jester JesterRole) getNightChatGroup() int {
	return CHAT_NOT_ALLOWED
}

func (jester JesterRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	return barriers
}

func (jester JesterRole) initialize() {
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getTeam() == TEAM_JESTER {
			mafiaUsers[i].Traits = append(mafiaUsers[i].Traits, JesterTrait{})
		}
	}
}

func (jester JesterRole) getSelfShowRoleID() int {
	return ROLE_JESTER
}
