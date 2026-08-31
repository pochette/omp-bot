package delivery

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/app/path"
	deliveryservice "github.com/ozonmp/omp-bot/internal/service/logistic/delivery"
)

type DeliveryCommander interface {
	Help(inputMsg *tgbotapi.Message)
	Get(inputMsg *tgbotapi.Message)
	List(inputMsg *tgbotapi.Message)
	Delete(inputMsg *tgbotapi.Message)
	New(inputMsg *tgbotapi.Message)
	Edit(inputMsg *tgbotapi.Message)

	HandleCommand(
		message *tgbotapi.Message,
		commandPath path.CommandPath,
	)

	HandleCallback(
		callback *tgbotapi.CallbackQuery,
		callbackPath path.CallbackPath,
	)
}

type commander struct {
	bot     *tgbotapi.BotAPI
	service deliveryservice.Service
}

func NewDeliveryCommander(
	bot *tgbotapi.BotAPI,
	service deliveryservice.Service,
) DeliveryCommander {
	return &commander{
		bot:     bot,
		service: service,
	}
}

func (c *commander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.CallbackName {
	case "list":
		c.CallBackList(callback, callbackPath)
	default:
		log.Printf("commander.HandleCommand: unknown callback name - %s", callbackPath.CallbackName)
	}

}

func (c *commander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.CommandName {
	case "help":
		c.Help(message)
	case "list":
		c.List(message)
	case "get":
		c.Get(message)
	default:
		c.Default(message)

	}
}
