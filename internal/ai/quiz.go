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
		answer := AskGemini(question, quiz)
		fmt.Printf("ИИ выбрал: %d\n", answer)
		if answer == 0 {
			fmt.Printf("Ошибка, ИИ не смог найти ответ\n")
			break
		}
		ans := fmt.Sprintf(`//div[@class="flex flex-col mt-5"]/div[%d]`, answer)
		page.MustElementX(ans).MustClick()
		time.Sleep(500 * time.Millisecond)
		page.MustElementX(`//button[@class="max-lg:px-[36px] max-sm:px-[40px] px-[72px]  py-[12px] gap-1 rounded-[24px] bg-[#0068FF] flex justify-center items-center"]`).MustClick()
	}
	fmt.Println("Все вопросы пройдены. Завершаю тест...")
}
