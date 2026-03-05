package ai

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func SolvingQuiz(page *rod.Page) {
	btn := page.MustActivate().MustElementX("//button[@class=\"text-white text-[15px] font-bold whitespace-nowrap\"]").MustWaitVisible()
	btn.MustClick()
	fmt.Println("Начинаю тест...")
	for i := 0; i < 5; i++ {
		fmt.Printf("Вопрос №%d\n", i+1)
		question := page.MustElementX("//span[@class=\"text-unix-text-black dark:text-[#AFB7CA] font-medium text-2xl\"]").MustText()
		options := page.MustElementsX("//p[@class=\"ml-4\"]")
		var quiz []string
		fmt.Printf("%s\n", question)
		for _, el := range options {
			txt := el.MustText()
			fmt.Println(txt)
			quiz = append(quiz, txt)
		}
		answer := AskGemini_Web(question, quiz[:]) //тут либо AskGemini_Web - без API или AskGemini - c API
		fmt.Printf("ИИ выбрал: %d\n", answer)
		if answer == 0 {
			fmt.Printf("Ошибка, ИИ не смог найти ответ\n")
			break
		}
		ans := fmt.Sprintf(`//div[@class="flex flex-col mt-5"]/div[%d]`, answer)
		page.MustActivate().MustElementX(ans).MustClick()
		time.Sleep(500 * time.Millisecond)
		page.MustActivate().MustElementX(`//div[@class="bg-[rounded-[24px] bg-white dark:bg-[#1A1A1A] py-[24px] px-[12px] mt-2 rounded-[24px]"]//button[2]`).MustClick()
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Все вопросы пройдены. Завершаю тест...")
}
