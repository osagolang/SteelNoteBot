package telegram

import (
	"context"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

func (h *Handler) HandleStart(msg *tgbotapi.Message) {

	user := &models.User{
		TelegramID: msg.From.ID,
		Username:   msg.From.UserName,
	}
	_ = h.userSVC.RegisterUser(context.Background(), user)

	/*

		btn1 := tgbotapi.NewInlineKeyboardButtonData("Тренировка1", "training")
		btn2 := tgbotapi.NewInlineKeyboardButtonData("Рекорды1", "records")
		keyboard := tgbotapi.NewInlineKeyboardMarkup(
			tgbotapi.NewInlineKeyboardRow(btn1, btn2),
		)

		msgConfig := tgbotapi.NewMessage(msg.Chat.ID, "Потренируемся или показать твои рекорды? Выбирай 👇")
		msgConfig.ReplyMarkup = keyboard

		h.bot.Send(msgConfig)

	*/

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тренировка", "training"),
			tgbotapi.NewInlineKeyboardButtonData("Рекорды", "records"),
		),
	)

	h.HandleSendMessage(msg.Chat.ID, "Потренируемся или показать твои рекорды? Выбирай 👇", keyboard)

}
