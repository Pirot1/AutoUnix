package main

import (
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
		log.Println("Error with creating .log file:", err)
		return
	}
	multi := io.MultiWriter(file, os.Stdout)
	log.SetOutput(multi)
	log.SetFlags(log.Ltime)

	err = godotenv.Load(".env") //loading .env
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
	if login == "" || password == "" || lessonName == "" { //Cheking data
		log.Print("Error: variable USER_EMAIL or USER_PASS are not founded in .env or they empty")
		parser.NewEnvFile()
		log.Println("Reload programm")
		time.Sleep(3 * time.Second)
		return
	}
	// 1. Browser initialisation
	b, page := browser.Init("https://uni-x.almv.kz/platform/login", true)
	defer b.MustClose()
	// 2. Autorisation
	parser.Autorisation(page, login, password)
	// 3. Lesson finder
	parser.FirstlLesson(page, lessonName)
	// Checking wether lesson is available
	page.MustActivate().MustElementX("//a[@class=\"flex flex-row items-center cursor-pointer\"][1]")
	// 4. Collecting video URLs
	urls := parser.GetAvailableLessons(page)
	if len(urls) == 0 {
		log.Println("All lessons are done, finishing up...")
		time.Sleep(1 * time.Second)
		return
	}
	// 5. Wathing lessons and solving test
	for _, url := range urls {
		log.Printf("Открываю: %s\n", url)
		page.MustNavigate(url)
		page.MustWaitLoad()

		parser.Proceed_lesson(page)
	}
	log.Println("All lessons are done, finishing up...")
	time.Sleep(1 * time.Second)
}
