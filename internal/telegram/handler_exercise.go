package telegram

import "context"

func (h *Handler) HandleExercise(chatID int64, muscleGroup string) {

	exercise, err := h.exerciseSVC.GetExerciseByGroup(context.Background(), muscleGroup)
	if err != nil {
		h.HandleSendMessage(chatID, "Ошибка получения упражнений", nil)
		return
	}
	if len(exercise) == 0 {
		h.HandleSendMessage(chatID, "Для этой группы мышц нет упражнений", nil)
	}

	btn := GenerateExerciseButtons(exercise)

	h.HandleSendMessage(chatID, "Выбирай упражнение 👇", btn)
}
