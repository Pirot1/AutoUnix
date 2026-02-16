package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/joho/godotenv"
)

func main() {
	// 1. Запуск браузера
	l := launcher.New().Headless(false).Devtools(false).Leakless(false)
	url, err := l.Launch()
	if err != nil {
		panic(fmt.Sprintf("Не удалось запустить браузер: %v", err))
	}
	err = godotenv.Load("../../.env") //Загрузка .env
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	login := os.Getenv("USER_EMAIL")
	password := os.Getenv("USER_PASS")
	if login == "" || password == "" { //Проверка, пустые ли данные
		log.Fatal("Ошибка: переменные USER_EMAIL или USER_PASS не найдены в .env или пусты")
	}
	browser := rod.New().ControlURL(url).MustConnect().NoDefaultDevice()
	defer browser.MustClose()

	page := browser.MustPage("https://uni-x.almv.kz/platform/login")

	fmt.Println("Начинаем работу...")

	// 2. Авторизация
	page.MustElement("input[type=\"email\"]").MustWaitVisible().MustInput(login)
	fmt.Println("Login succses")
	page.MustElement("input[type=\"password\"]").MustWaitVisible().MustInput(password)
	fmt.Println("Password succses")
	page.MustElement("button[type=\"submit\"]").MustClick()

	// Ждем загрузки личного кабинета
	page.MustWaitLoad()
	fmt.Println("Успешный вход!")
	time.Sleep(100 * time.Second)

}
func processLesson(page *rod.Page) {
	var lessonName string
	fmt.Print("Введите назвение урока: ")
	fmt.Scanln(&lessonName)
	fmt.Println("Ищу урок на странице...")
	page.MustElement("input[placeholder=\"Courses search\"]").MustWaitVisible().MustInput(lessonName)
	page.MustElement("div[class=\"h-full flex\"]").MustWaitVisible().MustClick()
	fmt.Print("Урок найден успешно!")

	//UNDER THE DEVELOPMENT
	// Ждем появления видеоплеера
	video := page.MustElement("video")

	// Запускаем видео, если оно не пошло само
	video.MustClick()

	// Используем JS, чтобы проверить окончание видео
	for {
		// Eval позволяет выполнить любой JS код прямо в консоли браузера через Go
		isEnded := page.MustEval(`() => {
			let v = document.querySelector('video');
			return v ? v.ended : false;
		}`).Bool()

		if isEnded {
			fmt.Println("✅ Видео досмотрено!")
			break
		}

		fmt.Println("⏳ Видео еще идет... жду 10 секунд")
		time.Sleep(10 * time.Second)
	}

	// 4. Логика теста (заглушка)
	solveQuiz(page)
}

func solveQuiz(page *rod.Page) {
	fmt.Println("📝 Перехожу к тесту...")
	// Тут будет поиск текста вопроса и отправка в ИИ
}
