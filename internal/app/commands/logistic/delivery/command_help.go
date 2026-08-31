package delivery

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func (c *commander) Help(inputMessage *tgbotapi.Message) {
	outputMessage := tgbotapi.NewMessage(inputMessage.Chat.ID,
		"/help - help\n"+
			"/list - list products\n"+
			"/get{id} - get product by id \n"+
			"/delete{id} - to delete delivery by id ")
	_, err := c.bot.Send(outputMessage)
	if err != nil {
		log.Printf("commander.Help: error sending reply message to chat - %v", err)
	}

}
