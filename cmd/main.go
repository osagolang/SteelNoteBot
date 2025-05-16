package main

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/osagolang/SteelNoteBot/internal/config"
	"github.com/osagolang/SteelNoteBot/internal/repositories"
	"github.com/osagolang/SteelNoteBot/internal/services"
	"github.com/osagolang/SteelNoteBot/internal/telegram"
	"log"
)

func main() {
	// Подключение к БД
	configPostgres := config.GetPostgres()
	dbPool, err := pgxpool.New(context.Background(), configPostgres)
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}
	defer dbPool.Close()

	fmt.Println("Успешное подключение к БД")

	// Инициализация слоёв
	userRepo := repositories.NewUserRepo(dbPool)
	exerciseRepo := repositories.NewExerciseRepo(dbPool)
	recordRepo := repositories.NewRecordRepo(dbPool)

	userSVC := services.NewUserService(userRepo)
	exerciseSVC := services.NewExerciseService(exerciseRepo)
	recordSVC := services.NewRecordService(recordRepo)

	// Запуск бота
	telegram.StartBot(userSVC, exerciseSVC, recordSVC)
}
