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
		fmt.Println("Ошибка создания лог-файла:", err)
		return
	}
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ltime)

	err = godotenv.Load(".env") //Загрузка .env
	if err != nil {
		log.Print("Ошибка при загрузке .env файла")
		parser.NewEnvFile()
		log.Println("Перзапустите программу")
		time.Sleep(3 * time.Second)
		return
	}
	login := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASS")
	lessonName := os.Getenv("LESSON_NAME")
	if login == "" || password == "" || lessonName == "" { //Проверка, пустые ли данные
		log.Print("Ошибка: переменные USER_EMAIL или USER_PASS не найдены в .env или пусты")
		parser.NewEnvFile()
		log.Println("Перзапустите программу")
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
		log.Println("Все уроки выполнены, завершаю сессию")
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
	log.Println("Все уроки выполнены, завершаю сессию")
	time.Sleep(1 * time.Second)
}
