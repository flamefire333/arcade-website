package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

var chatIngoingChannel = make(chan interface{})
var chatOutgoingChannel = make(chan interface{})
var messageHolder MessageHolder

type Message struct {
	userID      string `json:"username"`
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

type ChatSendMessageRequest struct {
	userID        string
	displayName   string
	message       string
	phase         int
	startIndex    int
	avatar        string
	chatID        int
	toDisplayName string
}

type ChatReadMessageRequest struct {
	userID     string
	phase      int
	startIndex int
	chatID     int
	chatName   string
}

type ChatReadResponse struct {
	Status int       `json:"status"`
	Chat   []Message `json:"chat"`
}

type GetChatIDRequest struct {
	Name string
}

func getChatID(name string) int {
	mafiaIngoingChannel <- GetChatIDRequest{Name: name}
	return (<-mafiaOutgoingChannel).(int)
}

type GetChatNameRequest struct {
	Name string
}

func getChatName(name string) string {
	mafiaIngoingChannel <- GetChatNameRequest{Name: name}
	return (<-mafiaOutgoingChannel).(string)
}

func sendMessage(request ChatSendMessageRequest) {
	phaseHolder, ok := messageHolder.holders[request.phase]
	if !ok {
		phaseHolder = PhaseMessageHolder{messages: make([]Message, 0)}
	}
	message := Message{userID: request.userID, DisplayName: request.displayName, Message: request.message, Avatar: request.avatar,
		Phase: request.phase, Type: request.chatID, ID: len(phaseHolder.messages), recipient: request.toDisplayName}
	phaseHolder.messages = append(phaseHolder.messages, message)
	messageHolder.holders[request.phase] = phaseHolder
}

func readMessages(request ChatReadMessageRequest) []Message {
	phaseHolder, ok := messageHolder.holders[request.phase]
	if ok && len(phaseHolder.messages) >= request.startIndex {
		base := phaseHolder.messages[request.startIndex:]
		ret := make([]Message, 0)
		for _, message := range base {
			if (message.Type == request.chatID || request.chatID == CHAT_DEAD || message.Type == CHAT_ALL) && (message.recipient == "" || message.recipient == request.chatName) {
				ret = append(ret, message)
			}
		}
		return ret
	}
	return make([]Message, 0)
}

func chatRequestHandler() {
	messageHolder.holders = make(map[int]PhaseMessageHolder)
	for 1 == 1 {
		request := <-chatIngoingChannel
		switch r := request.(type) {
		case ChatSendMessageRequest:
			sendMessage(r)
			readRequest := ChatReadMessageRequest{userID: r.userID, phase: r.phase, startIndex: r.startIndex, chatID: r.chatID, chatName: r.displayName}
			chatOutgoingChannel <- readMessages(readRequest)
		case ChatReadMessageRequest:
			chatOutgoingChannel <- readMessages(r)
		default:
			log.Printf("Unexpected chat request %+v", request)
		}

	}
}

func addChatHandlers(router *mux.Router) {

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
		chatIngoingChannel <- ChatReadMessageRequest{username, phase, startID, chatID, chatName}
		messages := (<-chatOutgoingChannel).([]Message)
		json.NewEncoder(w).Encode(ChatReadResponse{Status: 0, Chat: messages})
	})

	router.HandleFunc("/api/mafia/chat/send", func(w http.ResponseWriter, r *http.Request) {
		decoder := json.NewDecoder(r.Body)
		type sendMessageRequestData struct {
			StartID       int    `json:"startID"`
			Phase         int    `json:"phase"`
			Avatar        string `json:"avatar"`
			UserName      string `json:"user_name"`
			CharacterName string `json:"character_name"`
			Message       string `json:"message"`
		}
		var data sendMessageRequestData
		err := decoder.Decode(&data)
		if err != nil {
			log.Printf("Failed to decode in chat send: " + err.Error())
			return
		}
		chatID := getChatID(data.UserName)
		chatIngoingChannel <- ChatSendMessageRequest{data.UserName, data.CharacterName, data.Message, data.Phase, data.StartID, data.Avatar, chatID, ""}
		messages := (<-chatOutgoingChannel).([]Message)
		json.NewEncoder(w).Encode(ChatReadResponse{Status: 0, Chat: messages})
	})
}

const CHAT_DEAD = 0
const CHAT_ALL = 1
const CHAT_MAFIA = 2
const CHAT_NOT_ALLOWED = 3
const CHAT_CULTIST = 4
