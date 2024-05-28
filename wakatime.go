package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/bwmarrin/discordgo"
)

type wakaResponses struct {
	Data struct {
		Range                      string `json:"range"`
		Average                    string `json:"human_readable_daily_average"`
		Total                      string `json:"human_readable_total"`
		Is_up_to_date              bool   `json:"is_up_to_date"`
		Is_coding_activity_visible bool   `json:"is_coding_activity_visible"`
		Is_other_usage_visible     bool   `json:"is_other_usage_visible"`
	} `json:"data"`
}

type languages struct {
	Data struct {
		Language []struct {
			Name    string  `json:"name"`
			Percent float64 `json:"Percent"`
		} `json:"languages"`
	} `json:"data"`
}

func wakatime(s *discordgo.Session, m *discordgo.MessageCreate, args []string) {
	username := args[1]
	url := fmt.Sprintf("https://api.wakatime.com/api/v1/users/%s/stats", username)
	client := &http.Client{
		Timeout: 20 * time.Second, // You can adjust the timeout duration as needed
	}
	response, err := client.Get(url)
	if err != nil {
		s.ChannelMessageSend(m.ChannelID, "Error fetching the data")
		log.Fatal(err)
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
		Range := strings.ReplaceAll(wakaResponse.Data.Range, "_", " ")

		if len(args) == 3 {
			if args[2] == "details" {
				if wakaResponse.Data.Is_other_usage_visible {
					var language languages
					err = json.Unmarshal(responseData, &language)
					fmt.Println(language)
					if err != nil {
						log.Fatal(err)
					}
					if len(language.Data.Language) == 0 {
						s.ChannelMessageSend(m.ChannelID, "No Data here to show")
						return
					}
					var message strings.Builder
					message.WriteString("**Languages Programmed in:**\n")
					for _, lang := range language.Data.Language {
						message.WriteString(fmt.Sprintf("```%s: %.2f%%```\n", lang.Name, lang.Percent))
					}
					time := fmt.Sprintf("**Data for %s**", Range)
					message.WriteString(time)
					s.ChannelMessageSend(m.ChannelID, message.String())

				} else {
					s.ChannelMessageSend(m.ChannelID, "Please enable **Display languages, editors, os, categories publicly** from WakaTime settings to access this command")
				}
			} else {
				s.ChannelMessageSend(m.ChannelID, "Wrong Command! **details** is the only command available")
			}
		} else {

			fmt.Println(wakaResponse)

			if wakaResponse.Data.Is_coding_activity_visible {
				var uptime string
				if wakaResponse.Data.Is_up_to_date {
					uptime = "is"
				} else {
					uptime = "isnot"
				}

				message := fmt.Sprintf("```The total time worked is %s with daily average of %s in range:%s.The data %s fresh```", wakaResponse.Data.Total, wakaResponse.Data.Average, Range, uptime)
				s.ChannelMessageSend(m.ChannelID, message)
			} else {
				s.ChannelMessageSend(m.ChannelID, "**Please go to wakatime settings** and check the **'Display code time publicly'** box to make this command work. Would be better if you keep the range to last 7 days")
			}
		}

	}
	if response.StatusCode == 404 {
		s.ChannelMessageSend(m.ChannelID, "```User not found```")

	}
}
