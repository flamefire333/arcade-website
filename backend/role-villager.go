package main

type VillagerRole struct {
	Dummy string
}

func (villager VillagerRole) getRoleID() int {
	return ROLE_VILLAGER
}

func (villager VillagerRole) getName() string {
	return "villager"
}

func (villager VillagerRole) getDescription() string {
	return "Villagers are boring and sided with the Village"
}

func (villager VillagerRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/766499554157658122.png?v=1"
}

func (villager VillagerRole) getTeam() int {
	return TEAM_VILLAGER
}

func (villager VillagerRole) getVotingBarriers() []VotingBarrierInterface {
	return make([]VotingBarrierInterface, 0)
}

func (villager VillagerRole) getNightChatGroup() int {
	return CHAT_NOT_ALLOWED
}
