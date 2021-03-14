package character

import (
	"arcade-website/database"
	"log"
)

type Character struct {
	ID      int
	Name    string
	Avatar  string
	GroupID int
}

var AllCharacters []Character

func SetupCharacters() {
	rows, err := database.Conn.Query("SELECT id, name, avatar, group_id FROM characters")
	if err != nil {
		log.Printf("Character SELECT failed %+v\n", err)
		return
	}
	characters := make([]Character, 0)
	for rows.Next() {
		c := Character{}
		rows.Scan(&c.ID, &c.Name, &c.Avatar, &c.GroupID)
		characters = append(characters, c)
	}
	AllCharacters = characters
}
