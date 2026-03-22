package main

import (
	"fmt"
	"log"
	"os"
	"path"
	"time"

	"github.com/joho/godotenv"

	"AutoUnix/internal/browser"
	"AutoUnix/internal/parser"
)

func main() {
	err := godotenv.Load(".env") //Загрузка .env
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	login := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASS")
	if login == "" || password == "" { //Проверка, пустые ли данные
		log.Fatal("Ошибка: переменные USER_EMAIL или USER_PASS не найдены в .env или пусты")
	}
	// 1. Запуск браузера
	b, page := browser.Init("https://uni-x.almv.kz/platform/login", true) // потом поставить false
	defer b.MustClose()
	// 2. Авторизация
	parser.Autorisation(page, login, password)
	// 3. Поиск урока
	lessonName := os.Getenv("LESSON_NAME")
	parser.FirstlLesson(page, lessonName)
	// Проверка доступен ли урок
	first_url := page.MustActivate().MustElementX("//a[@class=\"flex flex-row items-center cursor-pointer\"][1]").MustAttribute("href")
	id := path.Base(*first_url)
	fmt.Printf("Текущий id урока: %s\n", id)
	// 4. Собираем url видео которые нужно посмотреть
	urls := parser.GetAvailableLessons(page)
	if len(urls) == 0 {
		fmt.Println("Все уроки выполнены, завершаю сессию")
		time.Sleep(1 * time.Second)
		return
	}
	// 5. Смотрим уроки и решаем тесты
	for _, url := range urls {
		fmt.Printf("Открываю: %s\n", url)
		page.MustNavigate(url)
		page.MustWaitLoad()

		parser.Proceed_lesson(page)
	}
	fmt.Println("Все уроки выполнены, завершаю сессию")
	time.Sleep(1 * time.Second)
}
