package main

import (
	"strings"

	"github.com/bwmarrin/discordgo"
)

func create(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.Bot {
		return
	}

	if !strings.HasPrefix(m.Content, Prefix) {
		return
	}

	command := strings.TrimPrefix(m.Content, Prefix)
	arg := strings.Fields(command)
	if len(arg) == 0 {
		return
	}
	if arg[0] == "role" && len(arg) == 3 {
		createRole(s, m, arg)
	}

	if arg[0] == "prawin" && len(arg) == 1 {
		createExcuse(s, m)

	}
	if arg[0] == "waka" && len(arg) <= 3 {
		wakatime(s, m, arg)
	}

}
