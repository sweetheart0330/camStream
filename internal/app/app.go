package app

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"os"
)

func Run(ctx context.Context) {
	parseConfig()
	bot, err := tgbotapi.NewBotAPI("7235772396:AAFJ1WXs9qZ_n5OvrZUD9jsASGP723n6ZcY")
	if err != nil {
		log.Panic(err)
	}

	// Включаем дебаг-режим, если необходимо
	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)
}

func parseConfig() {
	path := os.Args[1]

	data, err := os.ReadFile("config.yaml")
	if err != nil {
		fmt.Println(err)
		return
	}

}
