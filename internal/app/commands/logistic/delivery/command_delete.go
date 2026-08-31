package delivery

import (
	"log"
	"strconv"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Delete(inputMsg *tgbotapi.Message) {
	args := inputMsg.CommandArguments()

	idx, err := strconv.Atoi(args)
	if err != nil {
		log.Println("wrong args", args)
		return
	}
	isRemoved, err := c.service.Remove(uint64(idx))

	if err != nil {
		log.Printf("DeliveryCommander.Delete: error delete the entity with id: %d, Error: %v", idx, err)
	}
	if !isRemoved {
		log.Printf("Entity with idx: %d was not deleted", idx)
	}

}
