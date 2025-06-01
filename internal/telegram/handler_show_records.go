package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func (h *Handler) HandleShowRecords(chatID int64) {

	topExercises := []struct {
		ID   int
		Name string
	}{
		{5, "Подтягивания"},
		{10, "Отжимания от пола"},
		{11, "Жим штанги лёжа"},
	}

	msg := "🏆 Твои достижения:\n\n"

	for _, ex := range topExercises {
		rec, err := h.recordSVC.GetBestResult(context.Background(), chatID, ex.ID)
		if err != nil {
			h.HandleSendMessage(chatID, "Ошибка получения результатов по упражнению", nil)
			return
		}

		m := fmt.Sprintf("%s: нет данных\n", ex.Name)
		if rec != nil {
			m = rec.FormatMsg()
		}

		msg += m
	}

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Тренироваться", "training"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Вернуться в главное меню", "start"),
		),
	)

	h.HandleSendMessage(chatID, msg, keyboard)

}
