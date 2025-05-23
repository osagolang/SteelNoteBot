package telegram

import (
	"context"
	"fmt"
)

func (h *Handler) HandleShowRecords(chatID int64) {

	topExercises := []struct {
		ID   int
		Name string
	}{
		{5, "Подтягивания"},
		{11, "Жим штанги лёжа"},
		{1, "Приседания со штангой"},
		{10, "Отжимания от пола"},
	}

	msg := "🏆 Твои достижения:\n\n"

	for _, ex := range topExercises {
		rec, err := h.recordSVC.GetBestResult(context.Background(), chatID, ex.ID)
		if err != nil {
			h.HandleSendMessage(chatID, "Ошибка получения результатов по упражнению", nil)
			return
		}

		if rec == nil {
			msg += fmt.Sprintf("%s: нет данных\n", ex.Name)
		} else {
			msg += fmt.Sprintf("%s: %.1f кг × %d раз\n", ex.Name, *rec.Weight, rec.Reps)
		}
	}

	h.HandleSendMessage(chatID, msg, nil)

}
