package mafia

import (
	"arcade-website/actions"
	"arcade-website/character"
	"arcade-website/mafiaUser"
	"arcade-website/roles"
	"arcade-website/shared"
	"arcade-website/voting"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var users []shared.User = make([]shared.User, 0)

type MafiaStatusRequest struct {
	Name string `json:"name"`
}

type GetUsersRequest struct{}

type MainRequest struct{}

type SetupRequest struct {
	Group       int
	Roles       []int
	ActiveRoles []shared.ActiveRole
}

type KickUserRequest struct {
	Name string
}

type VoteRequest struct {
	Name        string
	ContainerID int
	Vote        string
}

type RequestResponse struct {
	Status int         `json:"status"`
	Info   interface{} `json:"info"`
}

type WillRequest struct {
	Name string `json:"name"`
	Will string `json:"will"`
}

type ForceQuitRequest struct {
	Name string `json:"name"`
	Will string `json:"will"`
}

type FrontEndRoleData struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Team        int    `json:"team"`
	Icon        string `json:"icon"`
	Description string `json:"description"`
}

type CharacterGroup struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

func getMafiaUserByName(name string) *shared.MafiaUser {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Name == name {
			return &shared.MafiaUsers[i]
		}
	}
	return nil
}

func getMafiaStatus(name string) shared.MafiaStatus {
	currentMafiaUser := getMafiaUserByName(name)
	myCharacterName := ""
	if currentMafiaUser != nil {
		myCharacterName = currentMafiaUser.Character.Name
	}
	gameInfo := shared.CurrentGameInfo
	characters := make([]shared.ActiveCharacter, 0, len(shared.MafiaUsers))
	for _, user := range shared.MafiaUsers {
		character := user.Character
		roleToShow := -1
		if user.Name == name {
			roleToShow = user.Role.GetSelfShowRoleID()
		}
		if !user.Alive {
			roleToShow = user.Role.GetRoleID()
		}
		characters = append(characters, shared.ActiveCharacter{Name: character.Name, Avatar: character.Avatar, Alive: user.Alive, RoleID: roleToShow, ID: 0})
	}
	votingData := make([]shared.VotingData, 0)
	for _, barrier := range activeVotingBarriers {
		voters := barrier.GetVoters()
		for _, voter := range voters {
			if voter == myCharacterName {
				fields := barrier.GetBase().Fields
				votingData = append(votingData, shared.VotingData{Title: barrier.GetTitle(), Fields: fields, List: barrier.GetBase().GetTotalVoteList(barrier)})
				break
			}
		}
	}
	nonzeroActiveRoles := make([]shared.ActiveRole, 0)
	for i := range activeRoles {
		if activeRoles[i].Amount > 0 {
			nonzeroActiveRoles = append(nonzeroActiveRoles, activeRoles[i])
		}
	}
	status := shared.MafiaStatus{LobbyData: users, GameInfo: gameInfo, ActiveRoles: nonzeroActiveRoles, Characters: characters, MyCharacter: myCharacterName, VotingData: votingData}
	shared.MafiaStatusCache[name] = status
	return status
}

var activeRoles = make([]shared.ActiveRole, 0)
var activeVotingBarriers = make([]shared.VotingBarrierInterface, 0)

var lastDayChange = time.Now()

func setGameToDay() {
	shared.CurrentGameInfo.Phase = shared.CurrentGameInfo.Phase + 1
	shared.CurrentGameInfo.Day = true
	shared.CurrentGameInfo.DayCount = shared.CurrentGameInfo.DayCount + 1
	barrierIDs := make([]int, 0)
	barrierID := shared.GetNextVotingBarrierID()
	barrierIDs = append(barrierIDs, barrierID)
	fields := make([]shared.VoteField, 0)
	fields = append(fields, shared.VoteField{Type: "option", Options: shared.GetStandardVotingOptions(), BarrierID: barrierID})
	dayVotingBarrier := voting.DayVotingBarrier{Base: shared.VotingBarrierBase{Votes: shared.GetStandardBaseVotes(barrierIDs), Fields: fields}}
	activeVotingBarriers = []shared.VotingBarrierInterface{dayVotingBarrier}
	lastDayChange = time.Now()
}

func setGameToNight() {
	shared.CurrentGameInfo.Phase = shared.CurrentGameInfo.Phase + 1
	shared.CurrentGameInfo.Day = false
	activeVotingBarriers = make([]shared.VotingBarrierInterface, 0)
	for _, role := range shared.AllRoles {
		barriers := role.GetVotingBarriers()
		for i := range barriers {
			activeVotingBarriers = append(activeVotingBarriers, barriers[i])
		}
	}
	for i := range shared.MafiaUsers {
		keptTraits := make([]shared.Trait, 0)
		for j := range shared.MafiaUsers[i].Traits {
			if shared.MafiaUsers[i].Traits[j].ShouldKeepOnNightChange() {
				keptTraits = append(keptTraits, shared.MafiaUsers[i].Traits[j])
			}
		}
		shared.MafiaUsers[i].Traits = keptTraits
	}
	lastDayChange = time.Now()
}

func setupGame(r SetupRequest) {
	log.Printf("Setting up Game")
	activeRoles = r.ActiveRoles
	shared.CurrentGameInfo.Started = true
	shared.CurrentGameInfo.DayCount = 0
	rolesToUse := r.Roles
	if len(rolesToUse) != len(users) {
		return
	}
	rand.Shuffle(len(rolesToUse), func(i, j int) { rolesToUse[i], rolesToUse[j] = rolesToUse[j], rolesToUse[i] })
	groupCharacters := make([]character.Character, 0)
	for _, character := range character.AllCharacters {
		if character.GroupID == r.Group {
			groupCharacters = append(groupCharacters, character)
		}
	}
	rand.Shuffle(len(groupCharacters), func(i, j int) { groupCharacters[i], groupCharacters[j] = groupCharacters[j], groupCharacters[i] })
	shared.MafiaUsers = make([]shared.MafiaUser, 0, len(rolesToUse))
	for i := 0; i < len(rolesToUse); i = i + 1 {
		roleToUse := roles.GetRole(rolesToUse[i])
		shared.MafiaUsers = append(shared.MafiaUsers, mafiaUser.GenerateMafiaUser(users[i].Name, roleToUse, groupCharacters[i]))
	}
	//TODO this could be more efficient, but its a one time setup thing so :shrug:
	for i := range shared.AllRoles {
		for j := range rolesToUse {
			if shared.AllRoles[i].GetRoleID() == rolesToUse[j] {
				shared.AllRoles[i].Initialize()
				break
			}
		}
	}
	rand.Shuffle(len(shared.MafiaUsers), func(i, j int) {
		shared.MafiaUsers[i], shared.MafiaUsers[j] = shared.MafiaUsers[j], shared.MafiaUsers[i]
	})
	setGameToDay()
}

//TODO maybe add canVote check?
func handleVote(request VoteRequest) {
	mafiaUser := getMafiaUserByName(request.Name)
	if mafiaUser != nil {
		for _, barrier := range activeVotingBarriers {
			for _, field := range barrier.GetBase().Fields {
				if field.BarrierID == request.ContainerID {
					barrier.GetBase().Votes[request.ContainerID][mafiaUser.Character.Name] = request.Vote
				}
			}
		}
	}
}

func mainMafiaLogic() {
	if shared.CurrentGameInfo.Started {
		teamCounts := make(map[int]int, 0)
		teamCounts[shared.TEAM_VILLAGER] = 0
		teamCounts[shared.TEAM_MAFIA] = 0
		teamCounts[shared.TEAM_CULTIST] = 0
		totalAlive := 0
		for i := range shared.MafiaUsers {
			if shared.MafiaUsers[i].Alive {
				teamCounts[shared.MafiaUsers[i].Role.GetTeam()]++
				totalAlive++
			}
		}
		if teamCounts[shared.TEAM_MAFIA] == 0 && teamCounts[shared.TEAM_CULTIST] == 0 {
			actions.EndGame("Villagers Win: The evildoers are all dead")
			return
		}
		if 2*teamCounts[shared.TEAM_MAFIA] >= totalAlive {
			actions.EndGame("Mafia Wins: Mafia outnumber everyone else")
			return
		}
		if 2*teamCounts[shared.TEAM_CULTIST] >= totalAlive {
			actions.EndGame("Cultists Win: Cultists outnumber everyone else")
		}
		allBarriesDone := true
		for _, barrier := range activeVotingBarriers {
			votingList := barrier.GetVoters()
			for _, field := range barrier.GetBase().Fields {
				votes := barrier.GetBase().Votes[field.BarrierID]
				if len(votingList) != len(votes) {
					//log.Printf("Not done %+v %+v %+v", votingList, votes, barrier)
					allBarriesDone = false
				}
			}
		}
		if allBarriesDone && time.Now().Sub(lastDayChange) > 30*time.Second {
			//TODO do barrier actions
			for _, barrier := range activeVotingBarriers {
				base := barrier.GetBase()
				option := base.GetWinningOption(barrier)
				if len(option) > 0 {
					barrier.ExecuteOption(option)
				}
			}
			// In case one of the barrier's executing caused the game to end (e.g. Day Voting killing a jester for a jester win)
			if shared.CurrentGameInfo.Started {
				if shared.CurrentGameInfo.Day {
					setGameToNight()
				} else {
					setGameToDay()
				}
			}
			shared.ClearMafiaStatusCache()
		}
	}
}

func MafiaRequestHandler() {
	for 1 == 1 {
		request := <-shared.MafiaIngoingChannel
		switch r := request.(type) {
		case MafiaStatusRequest:
			status, exists := shared.MafiaStatusCache[r.Name]
			if exists {
				shared.MafiaOutgoingChannel <- status
			} else {
				shared.MafiaOutgoingChannel <- getMafiaStatus(r.Name)
			}
		case shared.User:
			users = append(users, r)
			shared.ClearMafiaStatusCache()
		case shared.SendMessageRequestData:
			message := r.Message
			for i := range shared.MafiaUsers {
				if shared.MafiaUsers[i].Alive && shared.MafiaUsers[i].Character.Name == r.CharacterName {
					for j := range shared.MafiaUsers[i].Traits {
						message = shared.MafiaUsers[i].Traits[j].MessageConvert(message, &shared.MafiaUsers[i])
					}
				}
			}
			r.Message = message
			shared.MafiaOutgoingChannel <- r
			shared.ClearMafiaStatusCache()
		case KickUserRequest:
			for i := range users {
				if users[i].Name == r.Name {
					users = append(users[:i], users[i+1:]...)
					break
				}
			}
			shared.ClearMafiaStatusCache()
		case SetupRequest:
			setupGame(r)
			shared.ClearMafiaStatusCache()
		case VoteRequest:
			handleVote(r)
			shared.ClearMafiaStatusCache()
		case MainRequest:
			mainMafiaLogic()
		case WillRequest:
			user := getMafiaUserByName(r.Name)
			if user != nil {
				user.Will = r.Will
			}
		case ForceQuitRequest:
			if shared.CurrentGameInfo.Started {
				actions.EndGame("Ended by Admin")
			}
			shared.ClearMafiaStatusCache()
		case GetUsersRequest:
			shared.MafiaOutgoingChannel <- users
		case shared.GetChatNameRequest:
			ret := r.Name
			if shared.CurrentGameInfo.Started {
				for i := range shared.MafiaUsers {
					if shared.MafiaUsers[i].Name == r.Name {
						ret = shared.MafiaUsers[i].Character.Name
					}
				}
			}
			shared.MafiaOutgoingChannel <- ret
		case shared.GetChatIDRequest:
			found := false
			if shared.CurrentGameInfo.Started {
				for _, user := range shared.MafiaUsers {
					if user.Name == r.Name && !found {
						found = true
						if user.Alive {
							if shared.CurrentGameInfo.Day {
								shared.MafiaOutgoingChannel <- shared.CHAT_ALL
							} else {
								shared.MafiaOutgoingChannel <- user.Role.GetNightChatGroup()
							}
						} else {
							shared.MafiaOutgoingChannel <- shared.CHAT_DEAD
						}
					}
				}
				if !found {
					shared.MafiaOutgoingChannel <- shared.CHAT_DEAD
				}
			} else {
				shared.MafiaOutgoingChannel <- shared.CHAT_ALL
			}
		default:
			log.Printf("Unexpected mafia request %+v", request)
		}
	}
}

func GameRunner() {
	for 1 == 1 {
		shared.MafiaIngoingChannel <- MainRequest{}
		time.Sleep(time.Second)
	}
}

func AddMafiaHandlers(router *mux.Router) {

	router.HandleFunc("/api/login/mafia/user/{user}/create", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		shared.MafiaIngoingChannel <- shared.User{Name: username}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/mafia/status/{user}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		shared.MafiaIngoingChannel <- MafiaStatusRequest{Name: username}
		status := (<-shared.MafiaOutgoingChannel).(shared.MafiaStatus)
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: status})
	})
	router.HandleFunc("/api/mafia/kick/{user}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		shared.MafiaIngoingChannel <- KickUserRequest{Name: username}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/login/mafia/user", func(w http.ResponseWriter, r *http.Request) {
		shared.MafiaIngoingChannel <- GetUsersRequest{}
		users := <-shared.MafiaOutgoingChannel
		json.NewEncoder(w).Encode(users)
	})
	router.HandleFunc("/api/mafia/roles", func(w http.ResponseWriter, r *http.Request) {
		rolesToReturn := make([]interface{}, 0)
		for _, role := range shared.AllRoles {
			rolesToReturn = append(rolesToReturn, FrontEndRoleData{ID: role.GetRoleID(), Name: role.GetName(), Team: role.GetTeam(), Icon: role.GetIcon(), Description: role.GetDescription()})
		}
		json.NewEncoder(w).Encode(rolesToReturn)
	})
	router.HandleFunc("/api/mafia/character/groups", func(w http.ResponseWriter, r *http.Request) {
		groups := []CharacterGroup{{Name: "Fire Emblem", ID: 1}, {Name: "Genshin Impact", ID: 2}}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: groups})
	})
	router.HandleFunc("/api/mafia/will", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var data WillRequest
		err := decoder.Decode(&data)
		if err != nil {
			log.Printf("Will Decode Failure: %+v\n", err)
			json.NewEncoder(w).Encode(RequestResponse{Status: 1, Info: nil})
			return
		}
		shared.MafiaIngoingChannel <- data
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/mafia/forcequit", func(w http.ResponseWriter, r *http.Request) {
		shared.MafiaIngoingChannel <- ForceQuitRequest{}
	})
	router.HandleFunc("/api/mafia/vote/{user}/{containerID:[0-9]+}/{vote}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		containerID, err := strconv.Atoi(mux.Vars(r)["containerID"])
		if err != nil {
			log.Printf("Vote Failure %s\n", err.Error())
			json.NewEncoder(w).Encode(RequestResponse{Status: 1, Info: nil})
			return
		}
		vote := mux.Vars(r)["vote"]
		shared.MafiaIngoingChannel <- VoteRequest{Name: username, ContainerID: containerID, Vote: vote}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/mafia/setup", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		group, err := strconv.Atoi(params.Get("group"))
		if err != nil {
			return
		}
		rolesToUse := make([]int, 0)
		activeRolesToUse := make([]shared.ActiveRole, 0)
		for _, role := range shared.AllRoles {
			count, err := strconv.Atoi(params.Get(role.GetName()))
			if err == nil {
				activeRolesToUse = append(activeRolesToUse, shared.ActiveRole{Name: role.GetName(), Amount: count})
				for i := 0; i < count; i = i + 1 {
					rolesToUse = append(rolesToUse, role.GetRoleID())
				}
			}
		}
		shared.MafiaIngoingChannel <- SetupRequest{Group: group, Roles: rolesToUse, ActiveRoles: activeRolesToUse}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
}
