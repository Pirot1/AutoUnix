package main

import (
	"fmt"
	"log"
	"os"

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
	//time.Sleep(100 * time.Second) //для замедление кода
	processLesson(page) //переход на следующую стадию
}
func processLesson(page *rod.Page) {
	var lessonName string
	fmt.Print("Введите назвение урока: ")
	fmt.Scanln(&lessonName)
	fmt.Println("Ищу урок на странице...")
	page.MustElement("input[placeholder=\"Courses search\"]").MustWaitVisible().MustInput(lessonName)
	page.MustElement("div[class=\"h-full flex\"]").MustWaitVisible().MustClick()
	fmt.Println("Урок найден успешно!")

	var videoURLs []string
	elements := page.MustActivate().MustElementsX("//div[@class=\"overflow-y-auto bg-[#F0F3FA] dark:bg-black\"]/div[1]/div")
	for i := range elements {
		xpath := fmt.Sprintf("//div[@class=\"overflow-y-auto bg-[#F0F3FA] dark:bg-black\"]/div[1]/div[%d]", i+1)
		page.MustElementX(xpath).MustClick()
		xpath_links := fmt.Sprintf("%s//a", xpath)
		links := page.MustElementsX(xpath_links)
		fmt.Printf("В категории %d найдено ссылок: [%d]\nx", i+1, len(links))
		for _, url := range links {
			attr := url.MustAttribute("href")
			if attr != nil {
				videoURLs = append(videoURLs, *attr)
			}
		}
		fmt.Println("urls:", videoURLs)
		for _, url := range videoURLs {
			cur_cours := fmt.Sprintf("https://uni-x.almv.kz%s", url)
			fmt.Printf("Перехожу на курс: %s\n", cur_cours)
			page.MustNavigate(cur_cours)
			page.MustWaitLoad()
			fmt.Println("Успешно загрузил страницу!")
			//time.Sleep(100 * time.Second) //для замедление кода
			//UNDER THE DEVELOPMETN
			//PLS CHECK VIDEO IS DONE
			isDone, _, _ := .Has("img[src*='check']")

			if isDone {
				fmt.Println("⏩ Урок уже отмечен как выполненный.")
				continue
			}
		}
	}

}

func solveQuiz(page *rod.Page) {
	fmt.Println("📝 Перехожу к тесту...")
	// Тут будет поиск текста вопроса и отправка в ИИ
}
