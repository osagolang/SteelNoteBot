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
		return fmt.Sprintf("Неверный формат.\nВведи: вес повторения\nНапример: 100 5")
	} else {
		return fmt.Sprintf("Неверный формат.\nВведи: количество повторений\nНапример: 15")
	}
}
