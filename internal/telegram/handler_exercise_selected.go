package telegram

import (
	"context"
	"fmt"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
	"strconv"
	"strings"
)

func (h *Handler) HandleExerciseSelected(chatID int64, exerciseID int) {

	records, err := h.recordSVC.GetRecords(context.Background(), chatID, exerciseID, 5)
	if err != nil {
		h.HandleSendMessage(chatID, "Ошибка получения тренировок", nil)
		return
	}

	msg := ""

	if len(records) == 0 {
		msg = "Упс! Ты ещё ничего не записывал\n"
	} else {
		msg = "Последние тренировки:\n\n"

		for _, r := range records {
			msg += r.FormatLastMsg()
		}
	}

	exercise, err := h.exerciseSVC.GetExerciseByID(context.Background(), exerciseID)
	if err != nil {
		log.Printf("Ошибка получения данных по упражнению: %v", err)
		return
	}

	if exercise.HasWeight {
		msg += "\n---\nВведи данные текущей тренировки\nв формате: вес повторения\n(например: 90 10)"
	} else {
		msg += "\n---\nВведи данные текущей тренировки\nв формате: количество повторений\n(например: 15)"
	}

	h.tempInput[chatID] = exerciseID

	h.HandleSendMessage(chatID, msg, nil)

}

func (h *Handler) HandleTrainingMessage(msg *tgbotapi.Message) {
	chatID := msg.Chat.ID
	text := msg.Text

	exerciseID, ok := h.tempInput[chatID]
	if !ok {
		h.HandleSendMessage(chatID, "Ты ещё ничего не выбрал. Нажми на кнопку /Start", nil)
		return
	}

	exercise, err := h.exerciseSVC.GetExerciseByID(context.Background(), exerciseID)
	if err != nil {
		log.Printf("Ошибка получения данных по упражнению: %v", err)
		return
	}

	var weight *float64
	var reps int

	parts := strings.Fields(text)

	if exercise.HasWeight {
		if len(parts) != 2 {
			h.HandleSendMessage(chatID, "Неверный формат.\nВведи: вес повторения\nНапример: 100 5", nil)
			return
		}

		w, err1 := strconv.ParseFloat(parts[0], 64)
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			h.HandleSendMessage(chatID, "Неверный формат. Введи: вес повторения\nНапример: 100 5", nil)
			return
		}
		weight = &w
		reps = r
	} else {
		if len(parts) != 1 {
			h.HandleSendMessage(chatID, "Неверный формат.\nВведи: количество повторений\nНапример: 15", nil)
			return
		}

		r, err3 := strconv.Atoi(parts[0])
		if err3 != nil {
			h.HandleSendMessage(chatID, "Неверный формат.\nВведи: количество повторений\nНапример: 15", nil)
		}
		reps = r
	}

	err = h.recordSVC.AddRecord(context.Background(), chatID, exerciseID, weight, reps)
	if err != nil {
		log.Printf("Ошибка при сохранении тренировки: %v", err)
		h.HandleSendMessage(chatID, "Ошибка сохранения тренировки", nil)
		return
	}

	delete(h.tempInput, chatID)

	if exercise.HasWeight {
		h.HandleSendMessage(chatID, fmt.Sprintf("✅ Тренировка сохранена: %.1f кг. × %d раз", *weight, reps), nil)
	} else {
		h.HandleSendMessage(chatID, fmt.Sprintf("✅ Тренировка сохранена: %d раз", reps), nil)
	}
}
