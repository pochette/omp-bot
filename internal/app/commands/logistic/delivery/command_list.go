package delivery

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

const (
	initialOffset  = 0
	pageLimit      = 2
	nextButtonText = "Next page"
)

func (c *commander) List(inputMessage *tgbotapi.Message) {

	outputMessage := "Here all the products: \n\n"

	products, _ := c.service.List(initialOffset, pageLimit)
	for _, p := range products {
		outputMessage += p.String()
		outputMessage += "\n"
	}

	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, outputMessage)
	serializedData, _ := json.Marshal(CallbackListData{
		Offset: initialOffset + pageLimit,
		Limit:  pageLimit,
	})

	callbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "delivery",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(nextButtonText, callbackPath.String())))
	_, err := c.bot.Send(msg)
	if err != nil {
		log.Printf("DeliveryCommander.List: error sending reply message to chat - %v", err)
	}

}
