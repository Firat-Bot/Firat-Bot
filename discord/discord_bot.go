package discord

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"os"
	"os/signal"
	"syscall"
)

var (
	Token string
)

type Announcements struct {
	Title       string `json:"title"`
	Description string `json:"description"`
	Url         string `json:"url"`
}
type Lecturer struct {
	Name      string `json:"name"`
	Email     string `json:"email"`
	Phone     string `json:"phone"`
	Workspace string `json:"workspace"`
}

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
		url := "http://localhost:8080/annos"
		Announce(url, s, m)
	}

	if m.Content == "/prof" {
		url := "http://localhost:8080/lectures"
		Lectures(url, s, m)
	}

	if m.Content == "dilekce" || m.Content == "dilekçe" || m.Content == "dilekçeler" {
		s.ChannelMessageSend(m.ChannelID, "http://yaz.tf.firat.edu.tr/tr/page/560")
	}

	if m.Content == "e-posta" || m.Content == "ogrenci e-posta" || m.Content == "firat mail" {
		s.ChannelMessageSend(m.ChannelID, "https://posta.firat.edu.tr/")
	}

	if m.Content == "ders-icerikleri" || m.Content == "Ders İçerikleri" || m.Content == "ders icerikleri" {
		s.ChannelMessageSend(m.ChannelID, "http://yaz.tf.firat.edu.tr/subdomain_files/yaz.tf.firat.edu.tr/files/27/30.08.2021%20Guncel%20Ders%20I%CC%87c%CC%A7erikleri%20Genis%CC%A7letilmis%CC%A7%20Hali%20(1).pdf")
	}

	if m.Content == "ders listesi" || m.Content == "Ders Listesi" || m.Content == "müfredat" {
		s.ChannelMessageSend(m.ChannelID, "http://yaz.tf.firat.edu.tr/subdomain_files/yaz.tf.firat.edu.tr/files/27/YAZILIM%202016%20Mufredati%20(1).pdf\n")
	}
}
