package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateOrder(body models.CreateOrderRequest) error {
	if err := postgres.DB.Create(&body).Error; err != nil {
		return err
	}

	if err := sendMessagesToTelegram("7384309625:AAEafLrMt8MpZHAvfmvj6U3b0AZ8X3YTzs0", int64(body.UserId), body.Titles); err != nil {
		return err
	}

	return nil
}

func GetOrder(id int) models.Order {
	var order models.Order
	postgres.DB.Raw("SELECT * FROM orders WHERE id = ?", id).Scan(&order)

	return order
}

func sendMessagesToTelegram(botToken string, chatID int64, messages []string) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	prefix := "Ваш заказ успешно принят! Вы заказали:\n"
	messageText := prefix + strings.Join(messages, "\n")

	msg := tgbotapi.NewMessage(chatID, messageText)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}