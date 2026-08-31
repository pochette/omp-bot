package delivery

import (
	"encoding/json"
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
)

type CallbackListData struct {
	Offset uint64 `json:"offset"`
	Limit  uint64 `json:"limit"`
}

func (c *commander) CallBackList(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {

	var data = CallbackListData{}

	err := json.Unmarshal([]byte(callbackPath.CallbackData), &data)
	if err != nil {
		log.Printf("DeliveryCommander.CallbackList: "+
			"error reading json data for type CallbackListData from "+
			"input string %v - %v", callbackPath.CallbackData, err)
		return
	}

	delivers, err := c.service.List(data.Offset, data.Limit)
	if err != nil {
		log.Printf("CallBackList: error listing deliveries: %v", err)
		return
	}

	outputMessage := "Доставки: \n\n"

	for _, d := range delivers {
		outputMessage += d.String() + "\n"
	}

	nextCallbackListData := CallbackListData{
		Offset: data.Offset + data.Limit,
		Limit:  data.Limit,
	}

	serializedData, err := json.Marshal(nextCallbackListData)
	if err != nil {
		log.Printf("CallBackList: error encoding callback: %v", err)
		return
	}

	nextCallbackPath := path.CallbackPath{
		Domain:       "logistic",
		Subdomain:    "delivery",
		CallbackName: "list",
		CallbackData: string(serializedData),
	}

	msg := tgbotapi.NewMessage(
		callback.Message.Chat.ID,
		outputMessage)

	msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData(nextButtonText, nextCallbackPath.String())))

	_, err = c.bot.Send(msg)

	if err != nil {
		log.Printf("DeliveryCommander.CallbackList: error sending reply message to chat - %v", err)
	}

}
