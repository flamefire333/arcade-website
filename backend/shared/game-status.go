package shared

type GameStatus struct {
	Started  bool `json:"started"`
	Phase    int  `json:"phase"`
	Day      bool `json:"day"`
	DayCount int  `json:"dayCount"`
}

var CurrentGameInfo = GameStatus{Started: false, Phase: 0, Day: true, DayCount: 0}

type MafiaStatus struct {
	LobbyData   []User            `json:"lobbyData"`
	GameInfo    GameStatus        `json:"gameStatus"`
	ActiveRoles []ActiveRole      `json:"activeRoles"`
	Characters  []ActiveCharacter `json:"characters"`
	MyCharacter string            `json:"myCharacter"`
	VotingData  []VotingData      `json:"votingData"`
}

var MafiaStatusCache = make(map[string]MafiaStatus, 0)

func ClearMafiaStatusCache() {
	MafiaStatusCache = make(map[string]MafiaStatus, 0)
}

type User struct {
	Name string `json:"name"`
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
