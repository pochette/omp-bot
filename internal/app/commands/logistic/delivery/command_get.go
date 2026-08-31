package delivery

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Get(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Printf("wrong args: %v", args)
	}
	product, err := c.service.Get(uint64(idx))
	if err != nil {
		log.Printf("DeliveryCommander.Get: fail to get product with idx %d: %v", idx, err)
	}
	msg := tgbotapi.NewMessage(
		inputMsg.Chat.ID,
		product.String())
	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DeliveryCommander.Get: error sending reply message to chat - %v", err)
	}

}
