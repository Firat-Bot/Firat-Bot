package discord

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gophers/api/routers"
	"io/ioutil"
	"os"
	"os/signal"
	"syscall"
	"time"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func DiscordRun() {

	dg, err := discordgo.New("Bot " + Token)
	if err != nil {
		fmt.Println("error creating Discord session,", err)
		return
	}

	dg.AddHandler(messageCreate)

	dg.Identify.Intents = discordgo.IntentsGuildMessages

	err = dg.Open()
	if err != nil {
		fmt.Println("error opening connection,", err)
		return
	}

	fmt.Println("Bot is now running Press CTRL+C to exit  ")
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	<-sc
	dg.Close()
}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "/duyuru" {

		go routers.Router()
		time.Sleep(2 * time.Second)

		_, err := s.ChannelMessageSend(m.ChannelID, readFile())
		if err != nil {
			fmt.Println("Dosya gönderilemedi", err)
		}
		defer os.Remove("announcements.txt")
	}
}

func readFile() string {
	content, err := ioutil.ReadFile("announcements.txt")
	if err != nil {
		fmt.Println(err)
	}
	return string(content)
}
