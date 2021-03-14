package chat

import (
	"arcade-website/shared"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var messageHolder MessageHolder

type Message struct {
	userID      string
	Avatar      string `json:"avatar"`
	DisplayName string `json:"name"`
	Message     string `json:"message"`
	Phase       int    `json:"phase"`
	Type        int    `json:"messageType"`
	ID          int    `json:"id"`
	recipient   string
}

type PhaseMessageHolder struct {
	messages []Message
}

type MessageHolder struct {
	holders map[int]PhaseMessageHolder
}

type ChatReadResponse struct {
	Status int       `json:"status"`
	Chat   []Message `json:"chat"`
}

func getChatID(name string) int {
	shared.MafiaIngoingChannel <- shared.GetChatIDRequest{Name: name}
	return (<-shared.MafiaOutgoingChannel).(int)
}

func getChatName(name string) string {
	shared.MafiaIngoingChannel <- shared.GetChatNameRequest{Name: name}
	return (<-shared.MafiaOutgoingChannel).(string)
}

func sendMessage(request shared.ChatSendMessageRequest) {
	phaseHolder, ok := messageHolder.holders[request.Phase]
	if !ok {
		phaseHolder = PhaseMessageHolder{messages: make([]Message, 0)}
	}
	message := Message{userID: request.UserID, DisplayName: request.DisplayName, Message: request.Message, Avatar: request.Avatar,
		Phase: request.Phase, Type: request.ChatID, ID: len(phaseHolder.messages), recipient: request.ToDisplayName}
	phaseHolder.messages = append(phaseHolder.messages, message)
	messageHolder.holders[request.Phase] = phaseHolder
}

func readMessages(request shared.ChatReadMessageRequest) []Message {
	phaseHolder, ok := messageHolder.holders[request.Phase]
	if ok && len(phaseHolder.messages) >= request.StartIndex {
		base := phaseHolder.messages[request.StartIndex:]
		ret := make([]Message, 0)
		for _, message := range base {
			if (message.Type == request.ChatID || request.ChatID == shared.CHAT_DEAD || message.Type == shared.CHAT_ALL) && (message.recipient == "" || message.recipient == request.ChatName) {
				ret = append(ret, message)
			}
		}
		return ret
	}
	return make([]Message, 0)
}

func ChatRequestHandler() {
	messageHolder.holders = make(map[int]PhaseMessageHolder)
	for 1 == 1 {
		request := <-shared.ChatIngoingChannel
		switch r := request.(type) {
		case shared.ChatSendMessageRequest:
			sendMessage(r)
			readRequest := shared.ChatReadMessageRequest{UserID: r.UserID, Phase: r.Phase, StartIndex: r.StartIndex, ChatID: r.ChatID, ChatName: r.DisplayName}
			shared.ChatOutgoingChannel <- readMessages(readRequest)
		case shared.ChatReadMessageRequest:
			shared.ChatOutgoingChannel <- readMessages(r)
		default:
			log.Printf("Unexpected chat request %+v", request)
		}

	}
}

func AddChatHandlers(router *mux.Router) {

	router.HandleFunc("/api/mafia/chat/read/{user}/{phase:[0-9]+}/{startID:[0-9]+}", func(w http.ResponseWriter, r *http.Request) {
		username := mux.Vars(r)["user"]
		phase, err := strconv.Atoi(mux.Vars(r)["phase"])
		if err != nil {
			log.Printf("Chat Read Failure 1: " + err.Error())
			return
		}
		startID, err := strconv.Atoi(mux.Vars(r)["startID"])
		if err != nil {
			log.Printf("Chat Read Failure 2: " + err.Error())
			return
		}
		chatName := getChatName(username)
		chatID := getChatID(username)
		shared.ChatIngoingChannel <- shared.ChatReadMessageRequest{UserID: username, Phase: phase, StartIndex: startID, ChatID: chatID, ChatName: chatName}
		messages := (<-shared.ChatOutgoingChannel).([]Message)
		json.NewEncoder(w).Encode(ChatReadResponse{Status: 0, Chat: messages})
	})

	router.HandleFunc("/api/mafia/chat/send", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		var data shared.SendMessageRequestData
		err := decoder.Decode(&data)
		if err != nil {
			log.Printf("Failed to decode in chat send: " + err.Error())
			return
		}
		chatID := getChatID(data.UserName)
		shared.MafiaIngoingChannel <- data
		data = (<-shared.MafiaOutgoingChannel).(shared.SendMessageRequestData)
		shared.ChatIngoingChannel <- shared.ChatSendMessageRequest{
			UserID:        data.UserName,
			DisplayName:   data.CharacterName,
			Message:       data.Message,
			Phase:         data.Phase,
			StartIndex:    data.StartID,
			Avatar:        data.Avatar,
			ChatID:        chatID,
			ToDisplayName: ""}
		messages := (<-shared.ChatOutgoingChannel).([]Message)
		json.NewEncoder(w).Encode(ChatReadResponse{Status: 0, Chat: messages})
	})
}
