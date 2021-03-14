package roles

import (
	"arcade-website/actions"
	"arcade-website/shared"
	"arcade-website/traits"
	"math/rand"
)

type CultistRole struct {
}

func (cultist CultistRole) GetRoleID() int {
	return shared.ROLE_CULTIST
}

func (cultist CultistRole) GetName() string {
	return "cultist"
}

func (cultist CultistRole) GetDescription() string {
	return "The cultist can vote to convert someone into a cultist. All cultist die if the cultist leader dies."
}

func (cultist CultistRole) GetIcon() string {
	return "https://cdn.discordapp.com/emojis/796875834832322600.png?v=1"
}

func (cultist CultistRole) GetTeam() int {
	return shared.TEAM_CULTIST
}

func (cultist CultistRole) GetNightChatGroup() int {
	return shared.CHAT_CULTIST
}

func (cultist CultistRole) GetVotingBarriers() []shared.VotingBarrierInterface {
	barriers := make([]shared.VotingBarrierInterface, 0, 1)
	barrierIDs := make([]int, 0)
	barrierID := shared.GetNextVotingBarrierID()
	barrierIDs = append(barrierIDs, barrierID)
	fields := make([]shared.VoteField, 0)
	fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
	barriers = append(barriers, CultistVotingBarrier{Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}})
	return barriers
}

func (cultist CultistRole) Initialize() {
	cultistIDs := make([]int, 0)
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Role.GetTeam() == shared.TEAM_CULTIST {
			cultistIDs = append(cultistIDs, i)
		}
	}
	rand.Shuffle(len(cultistIDs), func(i, j int) { cultistIDs[i], cultistIDs[j] = cultistIDs[j], cultistIDs[i] })
	shared.MafiaUsers[cultistIDs[0]].Traits = append(shared.MafiaUsers[cultistIDs[0]].Traits, traits.CultistLeaderTrait{})
	actions.SendPrivateInfoMessage("You are the cultist leader!", shared.CHAT_ALL, 1, shared.MafiaUsers[cultistIDs[0]].Character.Name)
}

func (cultist CultistRole) GetSelfShowRoleID() int {
	return shared.ROLE_CULTIST
}

type CultistVotingBarrier struct {
	Base shared.VotingBarrierBase
}

func (cvb CultistVotingBarrier) GetVoters() []string {
	voters := make([]string, 0)
	for _, user := range shared.MafiaUsers {
		if user.Role.GetTeam() == shared.TEAM_CULTIST && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (cvb CultistVotingBarrier) GetTitle() string {
	return "Cultist Vote"
}

func (cvb CultistVotingBarrier) GetOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range shared.MafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (cvb CultistVotingBarrier) ExecuteOption(option []string) {
	actions.SetPlayerRoleByNameFromVote(option[0], shared.ROLE_CULTIST, cvb)
}

func (cvb CultistVotingBarrier) GetBase() *shared.VotingBarrierBase {
	return &cvb.Base
}
