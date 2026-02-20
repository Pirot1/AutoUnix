package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/go-rod/rod"
	"github.com/joho/godotenv"

	"AutoUnix/internal/browser"
	"AutoUnix/internal/parser"
)

func main() {
	err := godotenv.Load("../../.env") //Загрузка .env
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	login := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASS")
	if login == "" || password == "" { //Проверка, пустые ли данные
		log.Fatal("Ошибка: переменные USER_EMAIL или USER_PASS не найдены в .env или пусты")
	}
	// 1. Запуск браузера
	b, page := browser.Init(login, password)
	defer b.MustClose()
	// 2. Авторизация
	parser.Autorisation(page, login, password)
	// 3. Поиск урока
	lessonName := os.Getenv("LESSON_NAME")
	parser.FirstlLesson(page, lessonName)
	//проверка доступен ли урок
	first_url := page.MustActivate().MustElementX("//a[@class=\"flex flex-row items-center cursor-pointer\"][1]").MustAttribute("href")
	id := path.Base(*first_url)
	fmt.Printf("Текущий id урока: %s", id)
	// 4. Начинаем смотреть видео
	parser.GetAvailableLessons(page)
	time.Sleep(100 * time.Second)
}

func solveQuiz(page *rod.Page) {
	fmt.Println("📝 Перехожу к тесту...")
	// Тут будет поиск текста вопроса и отправка в ИИ
}
