package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"

	"github.com/bwmarrin/discordgo"
)

type wakaResponses struct {
	Data struct {
		Range   string `json:"range"`
		Average string `json:"human_readable_daily_average"`
		Total   string `json:"human_readable_total"`
	} `json:"data"`
}

func wakatime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	username := args[1]
	url := fmt.Sprintf("https://api.wakatime.com/api/v1/users/%s/stats", username)
	response, err := http.Get(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error fetching the data")
		log.Fatal("Error Fetching the data")
	}
	defer response.Body.Close()

	if response.StatusCode == http.StatusOK {
		responseData, err := io.ReadAll(response.Body)
		if err != nil {
			log.Fatal(err)
		}

		var wakaResponse wakaResponses
		err = json.Unmarshal(responseData, &wakaResponse)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(wakaResponse.Data.Average)
		fmt.Println(wakaResponse.Data.Total)
		fmt.Println(wakaResponse.Data.Range)
		Range := strings.ReplaceAll(wakaResponse.Data.Range, "_", " ")
		message := fmt.Sprintf("The Total time worked is %s with daily average of %s in range %s", wakaResponse.Data.Total, wakaResponse.Data.Average, Range)
		s.ChannelMessageSend(m.ChannelID, message)

	}
	if response.StatusCode == 404 {
		s.ChannelMessageSend(m.ChannelID, "User not found")
		log.Fatal("User Not Found")
	}
}
