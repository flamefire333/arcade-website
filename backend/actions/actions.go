package actions

import (
	"arcade-website/shared"
	"log"
)

func KillPlayerByNameFromVote(name string, vote shared.VotingBarrierInterface, phaseMod int) {
	user := GetMafiaUserByCharacterName(name)
	if user == nil {
		log.Printf("Could Not Find %s to Kill", name)
		return
	}
	if user.Alive {
		user.Alive = false
		will := user.Will
		SendInfoMessage(user.Character.Name+" has been killed, they left the following will \""+will+"\"", shared.CHAT_ALL, phaseMod)
		for i := range user.Traits {
			user.Traits[i].OnDeathByVote(vote, user, phaseMod)
		}
	}
}

func KillPlayerByNameFromPlayerAction(name string, killer *shared.MafiaUser, phaseMod int) {
	user := GetMafiaUserByCharacterName(name)
	if user == nil {
		log.Printf("Could Not Find %s to Kill", name)
		return
	}
	if user.Alive {
		user.Alive = false
		will := user.Will
		SendInfoMessage(user.Character.Name+" has been killed, they left the following will \""+will+"\"", shared.CHAT_ALL, phaseMod)
		for i := range user.Traits {
			user.Traits[i].OnDeathByPlayerAction(killer, user, phaseMod)
		}
	}
}

func SendInfoMessage(message string, chatID int, phaseMod int) {
	log.Printf("INFO MESSAGE: " + message)
	shared.ChatIngoingChannel <- shared.ChatSendMessageRequest{UserID: "", DisplayName: "Info", Message: message, Phase: shared.CurrentGameInfo.Phase + phaseMod, StartIndex: 0, Avatar: "https://cdn.discordapp.com/emojis/759196861927260171.png?v=1", ChatID: chatID, ToDisplayName: ""}
	//Clear up the response so chat continues to work
	<-shared.ChatOutgoingChannel
}

func SendPrivateInfoMessage(message string, chatID int, phaseMod int, recipient string) {
	log.Printf("PRIVATE INFO MESSAGE: " + message)
	shared.ChatIngoingChannel <- shared.ChatSendMessageRequest{UserID: "", DisplayName: "Info", Message: message, Phase: shared.CurrentGameInfo.Phase + phaseMod, StartIndex: 0, Avatar: "https://cdn.discordapp.com/emojis/759196861927260171.png?v=1", ChatID: chatID, ToDisplayName: recipient}
	//Clear up the response so chat continues to work
	<-shared.ChatOutgoingChannel
}

func GetMafiaUserByCharacterName(name string) *shared.MafiaUser {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Character.Name == name {
			return &shared.MafiaUsers[i]
		}
	}
	return nil
}

func SetPlayerRoleByNameFromVote(name string, roleID int, vote shared.VotingBarrierInterface) {
	user := GetMafiaUserByCharacterName(name)
	if user == nil {
		log.Printf("Could Not Find %s to Kill", name)
		return
	}
	if user.Alive {
		for i := range shared.AllRoles {
			if shared.AllRoles[i].GetRoleID() == roleID {
				user.Role = shared.AllRoles[i]
				SendPrivateInfoMessage("You have been converted into a "+user.Role.GetName(), shared.CHAT_ALL, 1, name)
				break
			}
		}
	}
}

func EndGame(message string) {
	SendInfoMessage(message, shared.CHAT_ALL, 0)
	shared.CurrentGameInfo.Started = false
	shared.CurrentGameInfo.Day = true
	clearMafiaStatusCache()
}

func clearMafiaStatusCache() {
	shared.MafiaStatusCache = make(map[string]shared.MafiaStatus, 0)
}

func AddTraitToPlayerByName(name string, trait shared.Trait) {
	for i := range shared.MafiaUsers {
		if shared.MafiaUsers[i].Alive && shared.MafiaUsers[i].Character.Name == name {
			shared.MafiaUsers[i].Traits = append(shared.MafiaUsers[i].Traits, trait)
		}
	}
}
