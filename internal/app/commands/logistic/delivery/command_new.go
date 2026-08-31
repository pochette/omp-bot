package delivery

import (
	"fmt"
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func (c *commander) New(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	cities := strings.SplitN(args, " | ", 2)
	cityFrom := strings.TrimSpace(cities[0])
	cityTo := strings.TrimSpace(cities[1])

	delivery := logistic.NewDelivery(cityFrom, cityTo, logistic.DeliveryStatus(1))
	_, err := c.service.Create(delivery)
	if err != nil {
		log.Printf("DeliveryCommander.New: Delivery was not created: %v", err)
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("Доставка создана: %v", delivery))

	_, e := c.bot.Send(msg)
	if err != nil {
		log.Printf("DeliveryCommander.New: Delivery was not created: %v", e)
	}

}
