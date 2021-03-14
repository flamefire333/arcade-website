package shared

type Trait interface {
	// Function that gets called when a user tries to send a message, returning modified message
	MessageConvert(message string, myself *MafiaUser) string

	//Function that gets called when user dies to a vote (Day vote, Mafia vote, etc.)
	OnDeathByVote(barrier VotingBarrierInterface, myself *MafiaUser, phaseMod int)

	//Function that gets called when user dies to a player action (e.g. The bomb)
	OnDeathByPlayerAction(killer *MafiaUser, myself *MafiaUser, phaseMod int)

	//Function that says this trait should be carried between nights
	ShouldKeepOnNightChange() bool
}
