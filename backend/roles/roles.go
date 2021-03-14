package roles

import (
	"arcade-website/shared"
	"log"
)

func InitRoles() {
	shared.AllRoles = nil
	shared.AllRoles = append(shared.AllRoles, VillagerRole{}, MafiaRole{}, DetectiveRole{}, BombRole{}, CultistRole{}, VoodooRole{}, JesterRole{}, ParanoidDetectiveRole{})
}

func GetRole(id int) shared.Role {
	for _, role := range shared.AllRoles {
		if role.GetRoleID() == id {
			return role
		}
	}
	log.Printf("Failed to find role id %d\n", id)
	return nil
}
