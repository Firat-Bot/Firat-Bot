package main

import (
	"bufio"
	"flag"
	"fmt"
	controller "gophers/api/controllers"
	routers "gophers/api/routers"
	"io"
	"log"
	"os"
	"strings"

	"github.com/bwmarrin/discordgo"
	"github.com/gin-gonic/gin"
)

var (
	Token string
)

func init() {
	flag.StringVar(&Token, "t", "", "Bot Token")
	flag.Parse()
}

func main() {

	someString := "hello world\nand hello go and more"
	myReader := strings.NewReader(someString)

	fmt.Printf("%T", myReader) // *strings.Reader

	buffer := make([]byte, 10)
	for {
		count, err := myReader.Read(buffer)
		if err != nil {
			if err != io.EOF {
				fmt.Println(err)
			}
			break
		}
		fmt.Printf("Count: %v\n", count)
		fmt.Printf("Data: %v\n", string(buffer))
	}

	/*dg, err := discordgo.New("Bot " + Token)
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
	dg.Close()*/

}
func router() gin.IRoutes {
	port := os.Getenv("PORT")
	if port == "" {
		port = ":8080"
	}

	router := gin.New()
	gin.SetMode(gin.ReleaseMode)
	routers.AnnoRoutes(router)

	fmt.Println(port, ": connected")
	router.Run(port)
	return router.GET("/annos", controller.GetAnnos())

}

func ReadFile(file string) []string {
	f, err := os.Open(file)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	scanner := bufio.NewScanner(f)
	var reader []string
	for scanner.Scan() {
		reader = append(reader, "%s\n", scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return reader

}
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {

	if m.Author.ID == s.State.User.ID {
		return
	}

	if m.Content == "duyuru" {

		router()
		read := ReadFile("announcements.txt")
		fmt.Println(read)
		_, err := s.ChannelFileSend(m.ChannelID, "duyuru", "")
		if err != nil {
			fmt.Println("Dosya gönderilemedi", err)
		}
	}

}
