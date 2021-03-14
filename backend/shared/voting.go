package shared

import (
	"math/rand"
	"strings"
)

type VotingBarrierInterface interface {
	GetVoters() []string
	GetOptions() []string
	ExecuteOption(option []string)
	GetBase() *VotingBarrierBase
	GetTitle() string
}

type VotingBarrierBase struct {
	Votes  map[int]map[string]string
	Fields []VoteField
}

type VoteField struct {
	Type      string   `json:"type"`
	Options   []string `json:"options"`
	BarrierID int      `json:"barrierID"`
}

type VotingData struct {
	Title  string      `json:"title"`
	Fields []VoteField `json:"fields"`
	List   []Vote      `json:"list"`
}

type Vote struct {
	Voter string `json:"name"`
	Voted string `json:"vote"`
}

func (voteBarrierBase VotingBarrierBase) GetVoteList(barrier VotingBarrierInterface, barrierID int) []Vote {
	votes := make([]Vote, 0, len(voteBarrierBase.Votes))
	voters := barrier.GetVoters()
	for _, voter := range voters {
		v, exists := voteBarrierBase.Votes[barrierID][voter]
		if exists {
			votes = append(votes, Vote{Voter: voter, Voted: v})
		} else {
			votes = append(votes, Vote{Voter: voter, Voted: "Has Not Voted"})
		}
	}
	return votes
}

func (voteBarrierBase VotingBarrierBase) GetTotalVoteList(barrier VotingBarrierInterface) []Vote {
	votes := make([]Vote, 0, len(voteBarrierBase.Votes))
	voters := barrier.GetVoters()
	for _, voter := range voters {
		items := make([]string, 0)
		for i := range barrier.GetBase().Fields {
			barrierID := barrier.GetBase().Fields[i].BarrierID
			v, exists := voteBarrierBase.Votes[barrierID][voter]
			if exists {
				items = append(items, v)
			} else {
				items = append(items, "Has Not Voted")
			}
		}
		votes = append(votes, Vote{Voter: voter, Voted: strings.Join(items, ", ")})
	}
	return votes
}

func (base VotingBarrierBase) GetWinningOption(barrier VotingBarrierInterface) []string {
	options := make([]string, 0)
	for i := range base.Fields {
		tallies := make(map[string]int, 0)
		voteList := base.GetVoteList(barrier, base.Fields[i].BarrierID)
		maxTally := 0
		for _, vote := range voteList {
			tally, exists := tallies[vote.Voted]
			if !exists {
				tally = 0

			}
			tally = tally + 1
			tallies[vote.Voted] = tally
			if tally > maxTally {
				maxTally = tally
			}
		}
		if maxTally == 0 {
			return make([]string, 0)
		}
		choices := make([]string, 0)
		hasNoOne := false
		for k, v := range tallies {
			if v == maxTally {
				if k == "No One" {
					hasNoOne = true
				}
				choices = append(choices, k)
			}
		}
		if hasNoOne {
			options = append(options, "No One")
		} else {
			rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })
			options = append(options, choices[0])
		}
	}
	return options
}

var nextVotingBarrierID = 0

func GetNextVotingBarrierID() int {
	id := nextVotingBarrierID
	nextVotingBarrierID = nextVotingBarrierID + 1
	return id
}

func GetStandardVotingOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for i := range MafiaUsers {
		if MafiaUsers[i].Alive {
			options = append(options, MafiaUsers[i].Character.Name)
		}
	}
	return options
}

func GetStandardBaseVotes(barrierIDs []int) map[int]map[string]string {
	outerMap := make(map[int]map[string]string, 0)
	for _, barrierID := range barrierIDs {
		innerMap := make(map[string]string, 0)
		outerMap[barrierID] = innerMap
	}
	return outerMap
}
