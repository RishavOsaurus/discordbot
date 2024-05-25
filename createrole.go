package main

import (
	"fmt"
	"strings"

	"github.com/bwmarrin/discordgo"
)

func createRole(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	fmt.Println("Inside the case")
	if len(args) < 3 {
		fmt.Println("Insufficient arguments")
		return
	}

	guildID := m.GuildID
	fmt.Println(guildID)

	roles, err := s.GuildRoles(guildID)
	if err != nil {
		fmt.Println("Error getting roles:", err)
		return
	}

	allRoles := make(map[string]string)
	for _, role := range roles {
		allRoles[role.Name] = role.ID
	}
	new_comp_role := strings.Join(args[2:], " ")
	reqRole, ok := allRoles[new_comp_role]
	fmt.Println(new_comp_role)
	fmt.Println("Here:", reqRole)
	if !ok {
		fmt.Println("Role not found")
		return
	}
	fmt.Println("reached here")

	// userID, err := getUserIDByUsername(s, guildID, args[1])
	// if err != nil {
	// 	fmt.Println("Error:", err)
	// 	return
	// }
	// fmt.Println("reacded here")
	userID := strings.TrimPrefix(args[1], "<@")
	userID = strings.TrimSuffix(userID, ">")
	fmt.Println("Hmm")
	err = s.GuildMemberRoleAdd(guildID, userID, reqRole)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Check you Permissions or Role Hierarchy")
		return
	}
	s.ChannelMessageSend(m.ChannelID, "<@"+userID+"> has been assigned the role "+new_comp_role)
}
