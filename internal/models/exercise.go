package models

import "fmt"

type Exercise struct {
	ID          int
	Name        string
	MuscleGroup string
	HasWeight   bool
}

func (e Exercise) FormatMsgHasWeight() string {
	if e.HasWeight {
		return fmt.Sprintf("\n---\nВведи данные текущей тренировки\nв формате: вес повторения\n(например: 90 10)")
	} else {
		return fmt.Sprintf("\n---\nВведи данные текущей тренировки\nв формате: количество повторений\n(например: 15)")
	}
}
