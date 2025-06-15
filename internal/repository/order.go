package repository

import (
	"catalog-bot-api/internal/models"
	"catalog-bot-api/pkg/postgres"
	"fmt"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func CreateOrder(body models.CreateOrderRequest) error {
	order := models.Order{
		UserId: body.UserId,
	}

	if err := postgres.DB.Create(&order).Error; err != nil {
		return err
	}

	var units []models.Unit
	for _, data := range body.DataList {
		units = append(units, models.Unit{
			Price:   data.Price,
			Title:   data.Title,
			OrderId: int(order.ID),
		})
	}
	postgres.DB.Create(&units)

	if err := sendMessagesToTelegram("7384309625:AAGQhUgjJ_caTtyfbiB88kDAquGJWPZrOuw", int64(body.UserId), body.DataList, int(order.ID)); err != nil {
		return err
	}

	return nil
}

func GetOrder(id int) models.OrderResponse {
	var order models.Order
	postgres.DB.Raw("SELECT * FROM orders WHERE id = ?", id).Scan(&order)

	var units []models.Unit
	postgres.DB.Raw("SELECT title FROM units WHERE order_id = ?", order.ID).Scan(&units)

	response := models.OrderResponse{
		Order: order,
		Units: units,
	}

	return response
}

func sendMessagesToTelegram(botToken string, chatID int64, messages []models.UnitStruct, orderId int) error {
	bot, err := tgbotapi.NewBotAPI(botToken)
	if err != nil {
		return err
	}

	prefix := fmt.Sprintf("Ваш заказ #%d успешно принят! Вы заказали:\n", orderId)
	// messageText := prefix + strings.Join(messages, "\n")
	messageText := prefix // + strings.Join(messages, "\n")
	for _, message := range messages {
		messageText += fmt.Sprintf("%s | %d\n", message.Title, message.Price)
	}

	msg := tgbotapi.NewMessage(chatID, messageText)

	_, err = bot.Send(msg)
	if err != nil {
		return err
	}

	return nil
}
