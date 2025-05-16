package telegram

import (
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"github.com/osagolang/SteelNoteBot/internal/models"
)

func GenerateExerciseButtons(exercises []models.Exercise) tgbotapi.InlineKeyboardMarkup {

	var rows [][]tgbotapi.InlineKeyboardButton
	for _, ex := range exercises {
		button := tgbotapi.NewInlineKeyboardButtonData(ex.Name, fmt.Sprintf("exercise_%d", ex.ID))
		rows = append(rows, tgbotapi.NewInlineKeyboardRow(button))
	}

	return tgbotapi.NewInlineKeyboardMarkup(rows...)
}
