package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (c *Commander) Help(inputMessage *tgbotapi.Message) {
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Бот помощник Яндекс еда\nСписок доступных команд:\n\n /list - list \n /help - help")
	c.bot.Send(msg)
}
