package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (h *Handler) HandleSendMessage(chatID int64, text string, replyMarkup interface{}) {
	msg := tgbotapi.NewMessage(chatID, text)
	if replyMarkup != nil {
		msg.ReplyMarkup = replyMarkup
	}
	h.bot.Send(msg)
}

func (h *Handler) HandleDeleteMessage(chatID int64, messageID int) {
	deleteMsg := tgbotapi.NewDeleteMessage(chatID, messageID)
	h.bot.Send(deleteMsg)
}

func (h *Handler) HandleTrainingMessage(msg *tgbotapi.Message) error {
	chatID := msg.Chat.ID
	text := msg.Text

	exerciseID, ok := h.tempInput[chatID]
	if !ok {
		h.HandleSendMessage(chatID, "Надо выбрать упражнение", nil)
		return nil
	}

	parts := strings.Fields(text)
	if len(parts) != 2 {
		h.HandleSendMessage(chatID, "Неверный формат.\nВведи: вес повторения\nНапример: 100 5", nil)
		return nil
	}

	weight, err1 := strconv.ParseFloat(parts[0], 64)
	reps, err2 := strconv.Atoi(parts[1])
	if err1 != nil || err2 != nil {
		h.HandleSendMessage(chatID, "Неверный формат. Введи: вес повторения\nНапример: 100 5", nil)
		return nil
	}

	err := h.recordSVC.AddRecord(context.Background(), chatID, exerciseID, weight, reps)
	if err != nil {
		log.Printf("Ошибка при сохранении тренировки: %v", err)
		h.HandleSendMessage(chatID, "Ошибка сохранения тренировки", nil)
		return nil
	}

	delete(h.tempInput, chatID)

	h.HandleSendMessage(chatID, fmt.Sprintf("✅ Тренировка сохранена: %.1f кг. × %d раз", weight, reps), nil)
	return nil
}
