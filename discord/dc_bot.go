package discord

import (
	"flag"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
	controller "gophers/api/controllers"
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
func AddDiscord() {
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

func router() gin.IRoutes {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	router.Run(port)
	return router.GET("/annos", controller.GetAnnos())
}

func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}
	if m.Content == "duyuru" {
		router()
	}
}
