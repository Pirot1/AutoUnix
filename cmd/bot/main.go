package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"time"

	"github.com/joho/godotenv"

	"AutoUnix/internal/browser"
	"AutoUnix/internal/parser"
)

func main() {
	file, err := os.OpenFile("bot.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666) // Запись log
	if err != nil {
		fmt.Println("Error while creating .log file:", err)
		return
	}
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ltime)

	err = godotenv.Load(".env") //Загрузка .env
	if err != nil {
		log.Print("Error while loading .env file")
		parser.NewEnvFile()
		log.Println("Reloading program...")
		time.Sleep(3 * time.Second)
		return
	}
	login := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASS")
	lessonName := os.Getenv("LESSON_NAME")
	if login == "" || password == "" || lessonName == "" { //Проверка, пустые ли данные
		log.Print("Error: values USER_EMAIL and USER_PASS are not founded in .env or empty")
		parser.NewEnvFile()
		log.Println("Reloading program...")
		time.Sleep(3 * time.Second)
		return
	}
	// 1. Запуск браузера
	b, page := browser.Init("https://uni-x.almv.kz/platform/login", true) // потом поставить false
	defer b.MustClose()
	// 2. Авторизация
	parser.Autorisation(page, login, password)
	// 3. Поиск урока
	parser.FirstlLesson(page, lessonName)
	// Проверка доступен ли урок
	page.MustActivate().MustElementX("//a[@class=\"flex flex-row items-center cursor-pointer\"][1]")
	// 4. Собираем url видео которые нужно посмотреть
	urls := parser.GetAvailableLessons(page)
	if len(urls) == 0 {
		log.Println("All lessons are done, finishing up...")
		time.Sleep(1 * time.Second)
		return
	}
	// 5. Смотрим уроки и решаем тесты
	for _, url := range urls {
		log.Printf("Открываю: %s\n", url)
		page.MustNavigate(url)
		page.MustWaitLoad()

		parser.Proceed_lesson(page)
	}
	log.Println("All lessons are done, finishing up...")
	time.Sleep(1 * time.Second)
}
