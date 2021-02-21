package main

import "math/rand"

type ParanoidDetectiveRole struct{}

func (detective ParanoidDetectiveRole) getRoleID() int {
	return ROLE_PARANOID_DETECTIVE
}

func (detective ParanoidDetectiveRole) getName() string {
	return "paranoid detective"
}

func (detective ParanoidDetectiveRole) getDescription() string {
	return "Paranoid Detective can check the alignment of a player as a night action, but will receive a random result. Paranoid Detectives are sided with the Village."
}

func (detective ParanoidDetectiveRole) getIcon() string {
	return "https://cdn.discordapp.com/emojis/796870998778445844.png?v=1"
}

func (detective ParanoidDetectiveRole) getTeam() int {
	return TEAM_VILLAGER
}

func (detective ParanoidDetectiveRole) getNightChatGroup() int {
	return CHAT_NOT_ALLOWED
}

func (detective ParanoidDetectiveRole) getVotingBarriers() []VotingBarrierInterface {
	barriers := make([]VotingBarrierInterface, 0, 1)
	for i := range mafiaUsers {
		if mafiaUsers[i].Role.getRoleID() == ROLE_PARANOID_DETECTIVE {
			barrierIDs := make([]int, 0)
			barrierID := GetNextVotingBarrierID()
			barrierIDs = append(barrierIDs, barrierID)
			fields := make([]VoteField, 0)
			fields = append(fields, VoteField{Type: "option", Options: getStandardVotingOptions(), BarrierID: barrierID})
			barriers = append(barriers, ParanoidDetectiveVotingBarrier{Base: VotingBarrierBase{Votes: getStandardBaseVotes(barrierIDs), Fields: fields}, Name: mafiaUsers[i].Character.Name})
		}
	}
	return barriers
}

func (detective ParanoidDetectiveRole) initialize() {}

func (detective ParanoidDetectiveRole) getSelfShowRoleID() int {
	return ROLE_DETECTIVE
}

type ParanoidDetectiveVotingBarrier struct {
	Base VotingBarrierBase
	Name string
}

func (dvb ParanoidDetectiveVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Character.Name == dvb.Name && user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (dvb ParanoidDetectiveVotingBarrier) getTitle() string {
	return "Detective Investigation"
}

func (dvb ParanoidDetectiveVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb ParanoidDetectiveVotingBarrier) executeOption(option []string) {
	for i := range mafiaUsers {
		if mafiaUsers[i].Character.Name == option[0] {
			options := make([]string, 0)
			options = append(options, "weird", "shady", "good")
			rand.Shuffle(len(options), func(i, j int) { options[i], options[j] = options[j], options[i] })
			state := options[0]
			sendPrivateInfoMessage(option[0]+" seems to be "+state+"!", CHAT_ALL, 1, dvb.Name)
		}
	}
}

func (dvb ParanoidDetectiveVotingBarrier) getBase() *VotingBarrierBase {
	return &dvb.Base
}
