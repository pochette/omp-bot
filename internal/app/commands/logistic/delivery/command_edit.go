package delivery

import (
	"fmt"
	"log"
	"strconv"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
	"github.com/ozonmp/omp-bot/internal/model/logistic"
)

func (c *commander) Edit(inputMsg *tgbotapi.Message) {
	args := strings.SplitN(inputMsg.CommandArguments(), " | ", 3)
	id, err := strconv.ParseUint(strings.TrimSpace(args[0]), 10, 64)
	if err != nil {
		_, _ = c.bot.Send(tgbotapi.NewMessage(inputMsg.Chat.ID,
			fmt.Sprintf("Wrong args idx data: %v", err)))
		return
	}

	cityFrom := strings.TrimSpace(args[1])
	cityTo := strings.TrimSpace(args[2])

	existing, err := c.service.Describe(id)
	if err != nil {
		log.Printf("DeliveryCommander.Edit: delivery not found: %v", err)
		return
	}
	updated := *existing
	updated.AddressTo = cityTo
	updated.AddressFrom = cityFrom
	updated.Status = logistic.DeliveryStatusUpdated

	err = c.service.Update(id, updated)
	if err != nil {
		log.Printf("DeliveryCommander.Edit: update failed: %v", err)
		return
	}

	msg := tgbotapi.NewMessage(inputMsg.Chat.ID,
		fmt.Sprintf("Доставка с id: %d обновлена: %v", id, updated))

	_, err = c.bot.Send(msg)
	if err != nil {
		log.Printf("DeliveryCommander.Edit: error sending message: %v", err)
	}
}
