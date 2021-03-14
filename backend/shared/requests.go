package shared

type ChatSendMessageRequest struct {
	UserID        string
	DisplayName   string
	Message       string
	Phase         int
	StartIndex    int
	Avatar        string
	ChatID        int
	ToDisplayName string
}

type ChatReadMessageRequest struct {
	UserID     string
	Phase      int
	StartIndex int
	ChatID     int
	ChatName   string
}

type SendMessageRequestData struct {
	StartID       int    `json:"startID"`
	Phase         int    `json:"phase"`
	Avatar        string `json:"avatar"`
	UserName      string `json:"user_name"`
	CharacterName string `json:"character_name"`
	Message       string `json:"message"`
}

type GetChatNameRequest struct {
	Name string
}

type GetChatIDRequest struct {
	Name string
}
