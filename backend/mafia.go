package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
)

var mafiaIngoingChannel = make(chan interface{})
var mafiaOutgoingChannel = make(chan interface{})
var users []User = make([]User, 0)
var mafiaUsers []MafiaUser = make([]MafiaUser, 0)

type User struct {
	Name string `json:"name"`
}

type MafiaUser struct {
	Name      string
	Alive     bool
	Role      Role
	Character Character
	Will      string
}

type MafiaStatusRequest struct {
	Name string `json:"name"`
}

type MafiaStatus struct {
	LobbyData   []User            `json:"lobbyData"`
	GameInfo    GameStatus        `json:"gameStatus"`
	ActiveRoles []ActiveRole      `json:"activeRoles"`
	Characters  []ActiveCharacter `json:"characters"`
	MyCharacter string            `json:"myCharacter"`
	VotingData  []VotingData      `json:"votingData"`
}

type GameStatus struct {
	Started  bool `json:"started"`
	Phase    int  `json:"phase"`
	Day      bool `json:"day"`
	DayCount int  `json:"dayCount"`
}

type ActiveRole struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type ActiveCharacter struct {
	Name   string `json:"name"`
	Avatar string `json:"avatar"`
	Alive  bool   `json:"alive"`
	RoleID int    `json:"roleID"`
	ID     int    `json:"id"`
}

type VotingBarrierInterface interface {
	getVoters() []string
	getOptions() []string
	executeOption(option string)
	getBase() *VotingBarrierBase
}

func (voteBarrierBase VotingBarrierBase) getVoteList(barrier VotingBarrierInterface) []Vote {
	votes := make([]Vote, 0, len(voteBarrierBase.Votes))
	voters := barrier.getVoters()
	for _, voter := range voters {
		v, exists := voteBarrierBase.Votes[voter]
		if exists {
			votes = append(votes, Vote{Voter: voter, Voted: v})
		} else {
			votes = append(votes, Vote{Voter: voter, Voted: "Has Not Voted"})
		}
	}
	return votes
}

var nextVotingBarrierID = 0

func GetNextVotingBarrierID() int {
	id := nextVotingBarrierID
	nextVotingBarrierID = nextVotingBarrierID + 1
	return id
}

type DayVotingBarrier struct {
	Base VotingBarrierBase
}

func (dvb DayVotingBarrier) getVoters() []string {
	voters := make([]string, 0)
	for _, user := range mafiaUsers {
		if user.Alive {
			voters = append(voters, user.Character.Name)
		}
	}
	return voters
}

func (dvb DayVotingBarrier) getOptions() []string {
	options := make([]string, 0)
	options = append(options, "No One")
	for _, user := range mafiaUsers {
		if user.Alive {
			options = append(options, user.Character.Name)
		}
	}
	return options
}

func (dvb DayVotingBarrier) executeOption(option string) {
	killPlayerByNameFromVote(option, dvb)
}

func (dvb DayVotingBarrier) getBase() *VotingBarrierBase {
	return &dvb.Base
}

type VotingBarrierBase struct {
	Votes map[string]string
	ID    int
}

func (base VotingBarrierBase) getWinningOption(barrier VotingBarrierInterface) (string, int) {
	tallies := make(map[string]int, 0)
	voteList := base.getVoteList(barrier)
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
		return "", 0
	}
	choices := make([]string, 0)
	for k, v := range tallies {
		if v == maxTally {
			if k == "No One" {
				return k, maxTally
			}
			choices = append(choices, k)
		}
	}
	rand.Shuffle(len(choices), func(i, j int) { choices[i], choices[j] = choices[j], choices[i] })
	return choices[0], maxTally
}

type VotingData struct {
	Title string `json:"title"`
	ID    int    `json:"id"`
	List  []Vote `json:"list"`
}

type Vote struct {
	Voter string `json:"name"`
	Voted string `json:"vote"`
}

type GetUsersRequest struct{}

type MainRequest struct{}

type SetupRequest struct {
	Group       int
	Roles       []int
	ActiveRoles []ActiveRole
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

func getMafiaUserByName(name string) *MafiaUser {
	for i := range mafiaUsers {
		if mafiaUsers[i].Name == name {
			return &mafiaUsers[i]
		}
	}
	return nil
}

func getMafiaUserByCharacterName(name string) *MafiaUser {
	for i := range mafiaUsers {
		if mafiaUsers[i].Character.Name == name {
			return &mafiaUsers[i]
		}
	}
	return nil
}

func clearMafiaStatusCache() {
	mafiaStatusCache = make(map[string]MafiaStatus, 0)
}

func getMafiaStatus(name string) MafiaStatus {
	currentMafiaUser := getMafiaUserByName(name)
	myCharacterName := ""
	if currentMafiaUser != nil {
		myCharacterName = currentMafiaUser.Character.Name
	}
	gameInfo := CurrentGameInfo
	characters := make([]ActiveCharacter, 0, len(mafiaUsers))
	for _, user := range mafiaUsers {
		character := user.Character
		roleToShow := -1
		if !user.Alive || user.Name == name {
			roleToShow = user.Role.getRoleID()
		}
		characters = append(characters, ActiveCharacter{Name: character.Name, Avatar: character.Avatar, Alive: user.Alive, RoleID: roleToShow, ID: 0})
	}
	votingData := make([]VotingData, 0)
	for _, barrier := range activeVotingBarriers {
		voters := barrier.getVoters()
		for _, voter := range voters {
			if voter == myCharacterName {
				votingData = append(votingData, VotingData{Title: "Vote", ID: barrier.getBase().ID, List: barrier.getBase().getVoteList(barrier)})
				break
			}
		}
	}
	status := MafiaStatus{LobbyData: users, GameInfo: gameInfo, ActiveRoles: activeRoles, Characters: characters, MyCharacter: myCharacterName, VotingData: votingData}
	mafiaStatusCache[name] = status
	return status
}

var CurrentGameInfo = GameStatus{Started: false, Phase: 0, Day: true, DayCount: 0}
var activeRoles = make([]ActiveRole, 0)
var mafiaStatusCache = make(map[string]MafiaStatus, 0)
var activeVotingBarriers = make([]VotingBarrierInterface, 0)

var lastDayChange = time.Now()

func setGameToDay() {
	CurrentGameInfo.Phase = CurrentGameInfo.Phase + 1
	CurrentGameInfo.Day = true
	CurrentGameInfo.DayCount = CurrentGameInfo.DayCount + 1
	dayVotingBarrier := DayVotingBarrier{Base: VotingBarrierBase{Votes: make(map[string]string, 0), ID: GetNextVotingBarrierID()}}
	activeVotingBarriers = []VotingBarrierInterface{dayVotingBarrier}
	lastDayChange = time.Now()
}

func setGameToNight() {
	CurrentGameInfo.Phase = CurrentGameInfo.Phase + 1
	CurrentGameInfo.Day = false
	activeVotingBarriers = make([]VotingBarrierInterface, 0)
	for _, role := range allRoles {
		barriers := role.getVotingBarriers()
		for i := range barriers {
			activeVotingBarriers = append(activeVotingBarriers, barriers[i])
		}
	}
	lastDayChange = time.Now()
}

func setupGame(r SetupRequest) {
	log.Printf("Setting up Game")
	activeRoles = r.ActiveRoles
	CurrentGameInfo.Started = true
	CurrentGameInfo.DayCount = 0
	rolesToUse := r.Roles
	if len(rolesToUse) != len(users) {
		return
	}
	rand.Shuffle(len(rolesToUse), func(i, j int) { rolesToUse[i], rolesToUse[j] = rolesToUse[j], rolesToUse[i] })
	groupCharacters := make([]Character, 0)
	for _, character := range allCharacters {
		if character.GroupID == r.Group {
			groupCharacters = append(groupCharacters, character)
		}
	}
	rand.Shuffle(len(groupCharacters), func(i, j int) { groupCharacters[i], groupCharacters[j] = groupCharacters[j], groupCharacters[i] })
	mafiaUsers = make([]MafiaUser, 0, len(rolesToUse))
	for i := 0; i < len(rolesToUse); i = i + 1 {
		roleToUse := getRole(rolesToUse[i])
		mafiaUsers = append(mafiaUsers, MafiaUser{Name: users[i].Name, Role: roleToUse, Character: groupCharacters[i], Alive: true, Will: ""})
	}
	rand.Shuffle(len(mafiaUsers), func(i, j int) { mafiaUsers[i], mafiaUsers[j] = mafiaUsers[j], mafiaUsers[i] })
	setGameToDay()
}

//TODO maybe add canVote check?
func handleVote(request VoteRequest) {
	mafiaUser := getMafiaUserByName(request.Name)
	if mafiaUser != nil {
		for _, barrier := range activeVotingBarriers {
			if barrier.getBase().ID == request.ContainerID {
				barrier.getBase().Votes[mafiaUser.Character.Name] = request.Vote
			}
		}
	}
}

func endGame(message string) {
	sendInfoMessage(message, CHAT_ALL, 0)
	CurrentGameInfo.Started = false
	CurrentGameInfo.Day = true
	clearMafiaStatusCache()
}

func mainMafiaLogic() {
	if CurrentGameInfo.Started {
		teamCounts := make(map[int]int, 0)
		teamCounts[TEAM_VILLAGER] = 0
		teamCounts[TEAM_MAFIA] = 0
		for i := range mafiaUsers {
			if mafiaUsers[i].Alive {
				teamCounts[mafiaUsers[i].Role.getTeam()]++
			}
		}
		if teamCounts[TEAM_MAFIA] == 0 {
			endGame("Villagers Win: Mafia is all dead")
			return
		}
		if teamCounts[TEAM_MAFIA] >= teamCounts[TEAM_VILLAGER] {
			endGame("Mafia Wins: Mafia outnumbers Villagers")
			return
		}
		allBarriesDone := true
		for _, barrier := range activeVotingBarriers {
			votingList := barrier.getVoters()
			votes := barrier.getBase().Votes
			if len(votingList) != len(votes) {
				allBarriesDone = false
			}
		}
		if allBarriesDone && time.Now().Sub(lastDayChange) > 30*time.Second {
			//TODO do barrier actions
			for _, barrier := range activeVotingBarriers {
				base := barrier.getBase()
				option, count := base.getWinningOption(barrier)
				if count > 0 {
					barrier.executeOption(option)
				}
			}
			if CurrentGameInfo.Day {
				setGameToNight()
			} else {
				setGameToDay()
			}
			clearMafiaStatusCache()
		}
	}
}

func sendInfoMessage(message string, chatID int, phaseMod int) {
	log.Printf("INFO MESSAGE: " + message)
	chatIngoingChannel <- ChatSendMessageRequest{userID: "", displayName: "Info", message: message, phase: CurrentGameInfo.Phase + phaseMod, startIndex: 0, avatar: "https://cdn.discordapp.com/emojis/759196861927260171.png?v=1", chatID: chatID}
	//Clear up the response so chat continues to work
	<-chatOutgoingChannel
}

func killPlayerByNameFromVote(name string, vote VotingBarrierInterface) {
	user := getMafiaUserByCharacterName(name)
	if user == nil {
		log.Printf("Could Not Find %s to Kill", name)
		return
	}
	if user.Alive {
		user.Alive = false
		will := user.Will
		sendInfoMessage(user.Character.Name+" has been killed, they left the following will \""+will+"\"", CHAT_ALL, 1)
	}
}

func mafiaRequestHandler() {
	for 1 == 1 {
		request := <-mafiaIngoingChannel
		switch r := request.(type) {
		case MafiaStatusRequest:
			status, exists := mafiaStatusCache[r.Name]
			if exists {
				mafiaOutgoingChannel <- status
			} else {
				mafiaOutgoingChannel <- getMafiaStatus(r.Name)
			}
		case User:
			users = append(users, r)
			clearMafiaStatusCache()
		case KickUserRequest:
			for i := range users {
				if users[i].Name == r.Name {
					users = append(users[:i], users[i+1:]...)
					break
				}
			}
			clearMafiaStatusCache()
		case SetupRequest:
			setupGame(r)
			clearMafiaStatusCache()
		case VoteRequest:
			handleVote(r)
			clearMafiaStatusCache()
		case MainRequest:
			mainMafiaLogic()
		case WillRequest:
			user := getMafiaUserByName(r.Name)
			if user != nil {
				user.Will = r.Will
			}
		case GetUsersRequest:
			mafiaOutgoingChannel <- users
		case GetChatIDRequest:
			found := false
			for _, user := range mafiaUsers {
				if user.Name == r.Name && !found {
					found = true
					if user.Alive {
						if CurrentGameInfo.Day {
							mafiaOutgoingChannel <- CHAT_ALL
						} else {
							mafiaOutgoingChannel <- user.Role.getNightChatGroup()
						}
					} else {
						mafiaOutgoingChannel <- CHAT_DEAD
					}
				}
			}
			if !found {
				mafiaOutgoingChannel <- CHAT_DEAD
			}
		default:
			log.Printf("Unexpected mafia request %+v", request)
		}
	}
}

func GameRunner() {
	for 1 == 1 {
		mafiaIngoingChannel <- MainRequest{}
		time.Sleep(time.Second)
	}
}

func addMafiaHandlers(router *mux.Router) {

	router.HandleFunc("/api/login/mafia/user/{user}/create", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		mafiaIngoingChannel <- User{Name: username}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/mafia/status/{user}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		mafiaIngoingChannel <- MafiaStatusRequest{Name: username}
		status := (<-mafiaOutgoingChannel).(MafiaStatus)
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: status})
	})
	router.HandleFunc("/api/mafia/kick/{user}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		mafiaIngoingChannel <- KickUserRequest{Name: username}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/login/mafia/user", func(w http.ResponseWriter, r *http.Request) {
		mafiaIngoingChannel <- GetUsersRequest{}
		users := <-mafiaOutgoingChannel
		json.NewEncoder(w).Encode(users)
	})
	router.HandleFunc("/api/mafia/roles", func(w http.ResponseWriter, r *http.Request) {
		roles := make([]interface{}, 0)
		for _, role := range allRoles {
			roles = append(roles, FrontEndRoleData{ID: role.getRoleID(), Name: role.getName(), Team: role.getTeam(), Icon: role.getIcon(), Description: role.getDescription()})
		}
		json.NewEncoder(w).Encode(roles)
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
		mafiaIngoingChannel <- data
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
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
		mafiaIngoingChannel <- VoteRequest{Name: username, ContainerID: containerID, Vote: vote}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
	router.HandleFunc("/api/mafia/setup", func(w http.ResponseWriter, r *http.Request) {
		params := r.URL.Query()
		group, err := strconv.Atoi(params.Get("group"))
		if err != nil {
			return
		}
		rolesToUse := make([]int, 0)
		activeRolesToUse := make([]ActiveRole, 0)
		for _, role := range allRoles {
			count, err := strconv.Atoi(params.Get(role.getName()))
			if err == nil {
				activeRolesToUse = append(activeRolesToUse, ActiveRole{Name: role.getName(), Amount: count})
				for i := 0; i < count; i = i + 1 {
					rolesToUse = append(rolesToUse, role.getRoleID())
				}
			}
		}
		mafiaIngoingChannel <- SetupRequest{Group: group, Roles: rolesToUse, ActiveRoles: activeRolesToUse}
		json.NewEncoder(w).Encode(RequestResponse{Status: 0, Info: nil})
	})
}
