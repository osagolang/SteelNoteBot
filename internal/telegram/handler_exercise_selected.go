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

	// Тут получаем англ. название группы мышц по ID
	muscleNameEng, err := h.exerciseSVC.GetExerciseByID(context.Background(), exerciseID)
	if err != nil {
		h.HandleSendMessage(chatID, "Ошибка получения группы мышц", nil)
		return
	}

	// Тут muscleGroupName из англ. названия вкидывает русское
	msg := fmt.Sprintf("Упражнение: %s\n---\n\n", muscleGroupName(muscleNameEng.Name))

	if len(records) == 0 {
		msg += "Упс! Ты ещё ничего не записывал\n"
	} else {
		msg += "Последние тренировки:\n\n"

		for _, r := range records {
			msg += r.FormatLastMsg()
		}
	}

	exercise, err := h.exerciseSVC.GetExerciseByID(context.Background(), exerciseID)
	if err != nil {
		log.Printf("Ошибка получения данных по упражнению: %v", err)
		return
	}

	msg += exercise.FormatMsgHasWeight()

	h.tempInput[chatID] = exerciseID
	msgID := h.HandleSendMessage(chatID, msg, nil)
	h.lastMsgID[chatID] = msgID

}

func (h *Handler) HandleTrainingMessage(msg *tgbotapi.Message) {

	chatID := msg.Chat.ID
	text := msg.Text
	messageID := msg.MessageID

	if lastMsgID, ok := h.lastMsgID[chatID]; ok {
		h.HandleDeleteMessage(chatID, lastMsgID)
		delete(h.lastMsgID, chatID)
	}

	exerciseID, ok := h.tempInput[chatID]
	if !ok {
		h.HandleSendMessage(chatID, "Вы ещё ничего не выбрали. Нажми на кнопку /Start", nil)
		return
	}

	exercise, err := h.exerciseSVC.GetExerciseByID(context.Background(), exerciseID)
	if err != nil {
		log.Printf("Ошибка получения данных по упражнению: %v", err)
		return
	}

	if exercise == nil {
		log.Printf("Нет упражнений...: %v", err)
		return
	}

	var weight *float64
	var reps int

	parts := strings.Fields(text)

	if exercise.HasWeight {
		if len(parts) != 2 {
			h.HandleDeleteMessage(chatID, messageID)
			badMsg := fmt.Sprintf("Неверный формат ввода для этого упражнения.\n\nВы ввели: \"%s\"\n\nВведите: вес повторения\nНапример: 100 5", text)
			mmm := h.HandleSendMessage(chatID, badMsg, nil)
			h.lastMsgID[chatID] = mmm
			return
		}

		w, err1 := strconv.ParseFloat(parts[0], 64)
		r, err2 := strconv.Atoi(parts[1])
		if err1 != nil || err2 != nil {
			h.HandleDeleteMessage(chatID, messageID)
			badMsg := fmt.Sprintf("Неверный формат ввода для этого упражнения.\n\nВы ввели: \"%s\"\n\nВведите: вес повторения\nНапример: 100 5", text)
			mmm := h.HandleSendMessage(chatID, badMsg, nil)
			h.lastMsgID[chatID] = mmm
			return
		}
		weight = &w
		reps = r
	} else {
		if len(parts) != 1 {
			h.HandleDeleteMessage(chatID, messageID)
			badMsg := fmt.Sprintf("Неверный формат ввода для этого упражнения.\n\nВы ввели: \"%s\"\n\nВведите: количество повторений\nНапример: 15", text)
			mmm := h.HandleSendMessage(chatID, badMsg, nil)
			h.lastMsgID[chatID] = mmm
			return
		}

		r, err3 := strconv.Atoi(parts[0])
		if err3 != nil {
			h.HandleDeleteMessage(chatID, messageID)
			badMsg := fmt.Sprintf("Неверный формат ввода для этого упражнения.\n\nВы ввели: \"%s\"\n\nВведите: количество повторений\nНапример: 15", text)
			mmm := h.HandleSendMessage(chatID, badMsg, nil)
			h.lastMsgID[chatID] = mmm
			return
		}
		reps = r
	}

	err = h.recordSVC.AddRecord(context.Background(), chatID, exerciseID, weight, reps)
	if err != nil {
		log.Printf("Ошибка при сохранении тренировки: %v", err)
		h.HandleSendMessage(chatID, "Ошибка сохранения тренировки", nil)
		return
	}

	h.HandleDeleteMessage(chatID, messageID)

	delete(h.tempInput, chatID)

	keyboard := tgbotapi.NewInlineKeyboardMarkup(
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Следующее упражнение", "training"),
		),
		tgbotapi.NewInlineKeyboardRow(
			tgbotapi.NewInlineKeyboardButtonData("Закончить тренировку", "start"),
		),
	)

	if exercise.HasWeight {
		h.HandleSendMessage(chatID, fmt.Sprintf(
			"✅ Тренировка сохранена!\n\n%s:\n%.1f кг. × %d раз\n\n---\nЧто делаем дальше? 👇", exercise.Name, *weight, reps), keyboard)
	} else {
		h.HandleSendMessage(chatID, fmt.Sprintf(
			"✅ Тренировка сохранена!\n\n%s:\n%d раз\n\n---\nЧто делаем дальше? 👇", exercise.Name, reps), keyboard)
	}
}
