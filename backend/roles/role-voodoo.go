package roles

import (
	"arcade-website/actions"
	"arcade-website/shared"
	"arcade-website/traits"
)

type VoodooRole struct{}

func (voodoo VoodooRole) GetRoleID() int {
	return shared.ROLE_VOODOO
}

func (voodoo VoodooRole) GetName() string {
	return "voodoo lady"
}

func (voodoo VoodooRole) GetDescription() string {
	return "The Voodoo Lady can choose a word that if a target says they die. Voodoo Ladies are sided with the Mafia."
}

func (voodoo VoodooRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/663189977177260082.png?v=1"
}

func (voodoo VoodooRole) GetTeam() int {
	return shared.TEAM_MAFIA
}

func (voodoo VoodooRole) GetNightChatGroup() int {
	return shared.CHAT_MAFIA
}

func (voodoo VoodooRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetRoleID() == shared.ROLE_VOODOO {
			barrierIDs := make([]int, 0)
			barrierID := shared.GetNextVotingBarrierID()
			barrierID2 := shared.GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID, barrierID2)
			fields := make([]shared.VoteField, 0)
			fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
			fields = append(fields, shared.VoteField{Type: "text", Options: make([]string, 0), BarrierID: barrierID2})
			barriers = append(barriers, VoodooVotingBarrier{Name: shared.MafiaUsers[i].Character.Name, Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}})
		}
	}
	return barriers
}

func (voodoo VoodooRole) Initialize() {}

func (voodoo VoodooRole) GetSelfShowRoleID() int {
	return shared.ROLE_VOODOO
}

type VoodooVotingBarrier struct {
	Base shared.VotingBarrierBase
	Name string
}

func (vvb VoodooVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Character.Name == vvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (vvb VoodooVotingBarrier) GetTitle() string {
	return "Voodoo Vote"
}

func (vvb VoodooVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (vvb VoodooVotingBarrier) ExecuteOption(option []string) {
	actions.AddTraitToPlayerByName(option[0], traits.VoodooCursedTrait{Word: option[1], VoodooPlayer: actions.GetMafiaUserByCharacterName(vvb.Name)})
}

func (vvb VoodooVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &vvb.Base
}
