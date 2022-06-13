package discord

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}
func main() {
	dg, err := discordgo.New(" Bot " + Token)
	if err != nil {
		fmt.Println("Error creating Discord Session", err)
		return
	}
	dg.AddHandler(messageCreate)
	dg.Identify.Intents = discordgo.IntentGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection", err)
		return
	}
	fmt.Println("Bot is now running Press CTRL+C to exit  ")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc

	dg.Close()
}

type Announcements struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	URL         string `json:"url"`
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "duyuru" {
		response, err := http.Get("Api Fırat Bot") //You can add api
		if err != nil {
			fmt.Println("Error occurred while invoking command ", err)
		}
		defer response.Body.Close()
		if response.StatusCode == 200 {
			_, err = s.ChannelFileSend(m.ChannelID, "announcements.txt", response.Body)
			if err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Error while send announcements")
		}
	}
}
