package ai

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func SolvingQuiz(page *rod.Page) {
	btn := page.MustActivate().MustElementX("//button[@class=\"text-white text-[15px] font-bold whitespace-nowrap\"]").MustWaitVisible()
	btn.MustClick()
	fmt.Println("Начинаю тест...")
	for i := 0; i < 5; i++ {
		fmt.Printf("Вопрос №%d\n", i+1)
		quest := fmt.Sprintf(`//div[@class="flex flex-row overflow-x-auto"]/div[%d]`, i+1)
		page.MustElementX(quest).MustClick()
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
			answer = rand.Intn(3) + 1
			fmt.Printf("ИИ не смог выбрать ответ, поэтому ответ был выбран случайно: %d\n", answer)
		}
		ans := fmt.Sprintf(`//div[@class="flex flex-col mt-5"]/div[%d]`, answer)
		page.MustActivate().MustElementX(ans).MustClick()
		time.Sleep(500 * time.Millisecond)
	}
	fmt.Println("Все вопросы пройдены. Завершаю тест...")
	page.MustElementX(`//button[@class="flex items-center text-white text-[15px] font-bold whitespace-nowrap max-sm:whitespace-normal"]`).MustClick()
	page.MustElementX(`//div[@class="flex-1 py-3 mr-1 border border-[#5F6B88] text-[#5F6B88] rounded-lg cursor-pointer text-[15px] font-bold text-center"]`).MustClick()
	grade := page.MustElementX(`//div[@class="flex items-center max-sm:justify-center"]/p/span[1]`).MustText()
	fmt.Printf("Оценка за тест: %s\n", grade)
}
