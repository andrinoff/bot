package discord_bot

import (
	"fmt"
	"log"
	"os"

	"github.com/bwmarrin/discordgo"
	"github.com/joho/godotenv"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
}



func checkNilErr(e error) {
 if e != nil {
  log.Fatal("Error message")
 }
}

func Send(message string) {
	var BotToken string = os.Getenv("DISCORD_TOKEN")
	discord, err := discordgo.New("Bot " + BotToken)
	checkNilErr(err)
	discord.Open()
	fmt.Println(message)
	discord.ChannelMessageSend("1383135547597258935", message)
	fmt.Println("Message sent")	
	discord.Close()
}