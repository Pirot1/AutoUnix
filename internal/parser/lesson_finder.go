package parser

import (
	"fmt"

	"github.com/go-rod/rod"
)

func FirstlLesson(page *rod.Page, lessonName string) {
	if lessonName == "" {
		fmt.Print("Введите назвение урока: ")
		fmt.Scanln(&lessonName)
	}
	fmt.Println("Ищу урок на странице...")
	page.MustElement("input[placeholder=\"Courses search\"]").MustWaitVisible().MustInput(lessonName)
	page.MustElement("div[class=\"h-full flex\"]").MustWaitVisible().MustClick()
	fmt.Println("Урок найден успешно!")

	page.MustActivate().MustElementX("//div[@class=\"overflow-y-auto bg-[#F0F3FA] dark:bg-black\"]/div[1]/div[1]").MustClick()
	first_url := page.MustActivate().MustElementX("//div[@class=\"overflow-y-auto bg-[#F0F3FA] dark:bg-black\"]/div[1]/div[1]/div[3]/a[1]").MustAttribute("href")
	page.MustNavigate("https://uni-x.almv.kz" + *first_url)
}
