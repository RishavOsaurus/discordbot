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
	if arg[0] == "role" {
		createRole(s, m, arg)
	}

	if arg[0] == "prawin" {
		createExcuse(s, m)
	}

}
