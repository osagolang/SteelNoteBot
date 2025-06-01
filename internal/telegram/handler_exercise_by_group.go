package telegram

import (
	"context"
	"fmt"
)

func muscleGroupName(eng string) string {
	muscleNames := map[string]string{
		"legs":      "Ноги",
		"back":      "Спина",
		"chest":     "Грудные",
		"shoulders": "Плечи",
		"biceps":    "Бицепс",
		"triceps":   "Трицепс",
		"calves":    "Икры",
		"press":     "Пресс",
	}

	if name, ok := muscleNames[eng]; ok {
		return name
	}
	return eng
}

func (h *Handler) HandleExerciseByGroup(chatID int64, muscleGroup string) {

	exercise, err := h.exerciseSVC.GetExerciseByGroup(context.Background(), muscleGroup)
	if err != nil {
		h.HandleSendMessage(chatID, "Ошибка получения упражнений", nil)
		return
	}
	if len(exercise) == 0 {
		h.HandleSendMessage(chatID, "Для этой группы мышц нет упражнений", nil)
	}

	btn := GenerateExerciseButtons(exercise)
	muscleName := muscleGroupName(muscleGroup)
	txt := fmt.Sprintf("Выбрана группа мышц: %s\n\nВыбирай упражнение 👇", muscleName)

	h.HandleSendMessage(chatID, txt, btn)

}
