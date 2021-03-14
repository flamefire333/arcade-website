package roles

import (
	"arcade-website/shared"
	"arcade-website/traits"
)

type JesterRole struct {
}

func (jester JesterRole) GetRoleID() int {
	return shared.ROLE_JESTER
}

func (jester JesterRole) GetName() string {
	return "jester"
}

func (jester JesterRole) GetDescription() string {
	return "The jester is trying to get voted off in the day vote. The jester is on its own team."
}

func (jester JesterRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/656062178159493135.png?v=1"
}

func (jester JesterRole) GetTeam() int {
	return shared.TEAM_JESTER
}

func (jester JesterRole) GetNightChatGroup() int {
	return shared.CHAT_NOT_ALLOWED
}

func (jester JesterRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	return barriers
}

func (jester JesterRole) Initialize() {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_JESTER {
			shared.MafiaUsers[i].Traits = append(shared.MafiaUsers[i].Traits, traits.JesterTrait{})
		}
	}
}

func (jester JesterRole) GetSelfShowRoleID() int {
	return shared.ROLE_JESTER
}
