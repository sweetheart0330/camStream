package tgBot

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type TgBot struct {
	bot *tgbotapi.BotAPI
}

func NewTgBot() (tg *TgBot, err error) {
	tg.bot, err = tgbotapi.NewBotAPI("7235772396:AAFJ1WXs9qZ_n5OvrZUD9jsASGP723n6ZcY")
	if err != nil {
		return nil, err
	}

	// Включаем дебаг-режим, если необходимо
	tg.bot.Debug = true

	log.Printf("Authorized on account %s", tg.bot.Self.UserName)

	return tg, err
}
