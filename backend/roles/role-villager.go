package roles

import "arcade-website/shared"

type VillagerRole struct {
}

func (villager VillagerRole) GetRoleID() int {
	return shared.ROLE_VILLAGER
}

func (villager VillagerRole) GetName() string {
	return "villager"
}

func (villager VillagerRole) GetDescription() string {
	return "Villagers are boring and sided with the Village"
}

func (villager VillagerRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/766499554157658122.png?v=1"
}

func (villager VillagerRole) GetTeam() int {
	return shared.TEAM_VILLAGER
}

func (villager VillagerRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	return make([]shared.VotingBarrierInterface, 0)
}

func (villager VillagerRole) GetNightChatGroup() int {
	return shared.CHAT_NOT_ALLOWED
}

func (villager VillagerRole) Initialize() {}

func (villager VillagerRole) GetSelfShowRoleID() int {
	return shared.ROLE_VILLAGER
}
