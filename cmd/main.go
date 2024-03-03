package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"os/signal"

	"github.com/go-telegram/bot"
	"github.com/go-telegram/bot/models"
	"github.com/joho/godotenv"

	"github.com/prandogabriel/bot-ideas/internal/database"
)

func main() {
	loadEnv()
	database.Connect()
	ctx, cancel := signal.NotifyContext(context.Background(), os.Interrupt)
	defer cancel()

	opts := []bot.Option{
		bot.WithDefaultHandler(handler),
	}

	b, err := bot.New(os.Getenv("TELEGRAM_BOT_TOKEN"), opts...)
	if err != nil {
		panic(err)
	}

	b.RegisterHandler(bot.HandlerTypeMessageText, "ola", bot.MatchTypeExact, handler2)
	b.Start(ctx)
}

func loadEnv() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func handler2(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("receive message handler 2, ", update.Message.Text)
}

func handler(ctx context.Context, b *bot.Bot, update *models.Update) {
	fmt.Println("receive message, ", update.Message.Text)
	b.SendMessage(ctx, &bot.SendMessageParams{
		ChatID: update.Message.Chat.ID,
		Text:   update.Message.Text,
	})
}
