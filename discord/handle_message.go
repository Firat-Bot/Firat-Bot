package discord

import (
	"encoding/json"
	"fmt"
	"github.com/bwmarrin/discordgo"
	"gophers/api/routers"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func Announce(url string, s *discordgo.Session, m *discordgo.MessageCreate) {

	go routers.RouterAnnounce()
	time.Sleep(5 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var announ []Announcements
	err = json.Unmarshal(body, &announ)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(announ)
	for i := range announ {
		_, err = s.ChannelMessageSend(m.ChannelID, announ[i].Title)
		_, err = s.ChannelMessageSend(m.ChannelID, announ[i].Description)
		_, err = s.ChannelMessageSend(m.ChannelID, announ[i].Url)

		if err != nil {
			fmt.Println("Don't Send File ", err)
		}
	}
}

func Lectures(url string, s *discordgo.Session, m *discordgo.MessageCreate) {

	go routers.RouterLectures()
	time.Sleep(5 * time.Second)

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}

	var lectures []Lecturer
	err = json.Unmarshal(body, &lectures)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(lectures)
	for i := range lectures {
		_, err = s.ChannelMessageSend(m.ChannelID, lectures[i].Name)
		_, err = s.ChannelMessageSend(m.ChannelID, lectures[i].Email)
		_, err = s.ChannelMessageSend(m.ChannelID, lectures[i].Phone)
		_, err = s.ChannelMessageSend(m.ChannelID, lectures[i].Workspace)

		if err != nil {
			fmt.Println("Don't Send File ", err)
		}
	}
}
