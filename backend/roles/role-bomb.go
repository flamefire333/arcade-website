package roles

import (
	"arcade-website/shared"
	"arcade-website/traits"
)

type BombRole struct {
}

func (bomb BombRole) GetRoleID() int {
	return shared.ROLE_BOMB
}

func (bomb BombRole) GetName() string {
	return "bomb"
}

func (bomb BombRole) GetDescription() string {
	return "The bomb is a villager who will kill one of the people responsible for killing them"
}

func (bomb BombRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/662793370317619226.png?v=1"
}

func (bomb BombRole) GetTeam() int {
	return shared.TEAM_VILLAGER
}

func (bomb BombRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	return make([]shared.VotingBarrierInterface, 0)
}

func (bomb BombRole) GetNightChatGroup() int {
	return shared.CHAT_NOT_ALLOWED
}

func (bomb BombRole) Initialize() {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetRoleID() == shared.ROLE_BOMB {
			shared.MafiaUsers[i].Traits = append(shared.MafiaUsers[i].Traits, traits.BombTrait{})
		}
	}
}

func (bomb BombRole) GetSelfShowRoleID() int {
	return shared.ROLE_BOMB
}
