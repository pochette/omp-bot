package logistic

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	deliverycommander "github.com/ozonmp/omp-bot/internal/app/commands/logistic/delivery"
	"github.com/ozonmp/omp-bot/internal/app/path"
	deliveryservice "github.com/ozonmp/omp-bot/internal/service/logistic/delivery"
)

type Commander interface {
	HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath)
	HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath)
}
type LogisticCommander struct {
	deliveryCommander Commander
}

func NewLogisticCommander(bot *tgbotapi.BotAPI) *LogisticCommander {
	service := deliveryservice.NewDummyDeliveryService()

	return &LogisticCommander{
		deliveryCommander: deliverycommander.NewDeliveryCommander(
			bot,
			service),
	}
}

func (c *LogisticCommander) HandleCallback(callback *tgbotapi.CallbackQuery, callbackPath path.CallbackPath) {
	switch callbackPath.Subdomain {
	case "delivery":
		c.deliveryCommander.HandleCallback(callback, callbackPath)
	default:
		log.Printf("LogisticCommander.HandleCallback: unknown subdomain - %s", callbackPath.Subdomain)
	}
}

func (c *LogisticCommander) HandleCommand(message *tgbotapi.Message, commandPath path.CommandPath) {
	switch commandPath.Subdomain {
	case "delivery":
		c.deliveryCommander.HandleCommand(message, commandPath)
	default:
		log.Printf("LogisticCommander.HandleCommand: unknown subdomain - %s",
			commandPath.Subdomain)

	}
}
