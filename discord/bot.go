package discord_bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
)

var BotToken string = os.Getenv("DISCORD_TOKEN")

func checkNilErr(e error) {
 if e != nil {
  log.Fatal("Error message")
 }
}

func Send(message string) {
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)
	discord.Open()
	fmt.Println(message)
	discord.ChannelMessageSend("1383135547597258935", message)
	fmt.Println("Message sent")
	// create a session
	

	// add a event handler

	// open session
	
	

	
	discord.Close()

}