package shared

type Role interface {
	GetRoleID() int
	GetName() string
	GetDescription() string
	GetIcon() string
	GetTeam() int
	GetVotingBarriers() []VotingBarrierInterface
	GetNightChatGroup() int
	Initialize()
	GetSelfShowRoleID() int
}

var AllRoles []Role
