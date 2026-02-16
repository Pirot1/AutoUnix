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
	// Launch() запускает браузер, Headless(false) позволяет видеть, что происходит.
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
	// Используем MustElement для поиска по селекторам
	page.MustElement("input[type=\"email\"]").MustWaitVisible().MustInput(login)
	//fmt.Println("Login succses")
	page.MustElement("input[type=\"password\"]").MustWaitVisible().MustInput(password)
	//fmt.Println("Password succses")
	page.MustElement("button[type=\"submit\"]").MustClick() //Нажатие по кнопке

	// Ждем загрузки личного кабинета
	page.MustWaitLoad()
	fmt.Println("Успешный вход!")
	time.Sleep(100 * time.Second) //test

}

// Заглушка для ИИ (сюда нужно вставить вызов OpenAI/Gemini)
func getAIAnswer(question string) string {
	fmt.Printf("Запрос к ИИ по вопросу: %s\n", question)
	// В реальности здесь будет HTTP запрос к API
	return "Вариант 2"
}
