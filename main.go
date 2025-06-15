package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"
	"strconv"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"

	discord_bot "bot/discord"
)
func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}
	b.Start(ctx)
}

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error loading .env file:", err)
	}
}
func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	if update.Message == nil {
		return // ignore updates that aren’t a message
	}

	var channel_ids = []int64{
		-1002158048191,
		-1002351566952,
	}

	chatID, err := strconv.Atoi(os.Getenv("CHAT_ID"))
	if err != nil {
		fmt.Println("Invalid CHAT_ID:", err)
		return
	}

	if int(update.Message.Chat.ID) == chatID {
		for _, channel_id := range channel_ids {
			b.SendMessage(ctx, &bot.SendMessageParams{
				ChatID:    channel_id,
				// ParseMode: "MarkdownV2",
				Text:      update.Message.Text,
			})
		}
		fmt.Println(update.Message.Text, "user verified")
		discord_bot.Send(update.Message.Text)
	} else {
		fmt.Println(update.Message.Text, "user not verified")
	}
}