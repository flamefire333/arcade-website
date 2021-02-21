package main

type BombRole struct {
}

func (bomb BombRole) getRoleID() int {
	return ROLE_BOMB
}

func (bomb BombRole) getName() string {
	return "bomb"
}

func (bomb BombRole) getDescription() string {
	return "The bomb is a villager who will kill one of the people responsible for killing them"
}

func (bomb BombRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/662793370317619226.png?v=1"
}

func (bomb BombRole) getTeam() int {
	return TEAM_VILLAGER
}

func (bomb BombRole) getVotingBarriers() []VotingBarrierInterface {
	return make([]VotingBarrierInterface, 0)
}

func (bomb BombRole) getNightChatGroup() int {
	return CHAT_NOT_ALLOWED
}

func (bomb BombRole) initialize() {
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getRoleID() == ROLE_BOMB {
			mafiaUsers[i].Traits = append(mafiaUsers[i].Traits, BombTrait{})
		}
	}
}

func (bomb BombRole) getSelfShowRoleID() int {
	return ROLE_BOMB
}
