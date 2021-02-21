package main

type VoodooRole struct{}

func (voodoo VoodooRole) getRoleID() int {
	return ROLE_VOODOO
}

func (voodoo VoodooRole) getName() string {
	return "voodoo lady"
}

func (voodoo VoodooRole) getDescription() string {
	return "The Voodoo Lady can choose a word that if a target says they die. Voodoo Ladies are sided with the Mafia."
}

func (voodoo VoodooRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/663189977177260082.png?v=1"
}

func (voodoo VoodooRole) getTeam() int {
	return TEAM_MAFIA
}

func (voodoo VoodooRole) getNightChatGroup() int {
	return CHAT_MAFIA
}

func (voodoo VoodooRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getRoleID() == ROLE_VOODOO {
			barrierIDs := make([]int, 0)
			barrierID := GetNextVotingBarrierID()
			barrierID2 := GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID, barrierID2)
			fields := make([]VoteField, 0)
			fields = append(fields, VoteField{Type: "option", Options: getStandardVotingOptions(), BarrierID: barrierID})
			fields = append(fields, VoteField{Type: "text", Options: make([]string, 0), BarrierID: barrierID2})
			barriers = append(barriers, VoodooVotingBarrier{Name: mafiaUsers[i].Character.Name, Base: VotingBarrierBase{Votes: getStandardBaseVotes(barrierIDs), Fields: fields}})
		}
	}
	return barriers
}

func (voodoo VoodooRole) initialize() {}

type VoodooVotingBarrier struct {
	Base VotingBarrierBase
	Name string
}

func (vvb VoodooVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Character.Name == vvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (vvb VoodooVotingBarrier) getTitle() string {
	return "Voodoo Vote"
}

func (vvb VoodooVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (vvb VoodooVotingBarrier) executeOption(option []string) {
	addTraitToPlayerByName(option[0], VoodooCursedTrait{Word: option[1], VoodooPlayer: getMafiaUserByCharacterName(vvb.Name)})
}

func (vvb VoodooVotingBarrier) getBase() *VotingBarrierBase {
	return &vvb.Base
}
