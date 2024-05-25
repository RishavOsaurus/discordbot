package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/bwmarrin/discordgo"
)

type excuse struct {
	Id       int    `json:"id"`
	Excuse   string `json:"excuse"`
	Category string `json:"category"`
}

func createExcuse(s *discordgo.Session, m *discordgo.MessageCreate) {
	response, err := http.Get("https://excuser-three.vercel.app/v1/excuse/college/1")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	responseData, err := io.ReadAll(response.Body)
	fmt.Println(string(responseData))
	if err != nil {
		log.Fatal(err)
	}

	var excuses []excuse
	err = json.Unmarshal(responseData, &excuses)
	if err != nil {
		log.Fatal(err)
	}

	if len(excuses) > 0 {
		fmt.Printf("Fetched Excuse: %+v\n", excuses[0])
		s.ChannelMessageSend(m.ChannelID, excuses[0].Excuse)
	} else {
		s.ChannelMessageSend(m.ChannelID, "No excuses found.")
	}

}
