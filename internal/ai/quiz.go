package ai

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/go-rod/rod"
)

func SolvingQuiz(page *rod.Page) {
	btn := page.MustActivate().MustElementX("//button[@class=\"text-white text-[15px] font-bold whitespace-nowrap\"]").MustWaitVisible()
	btn.MustClick()
	log.Println("Начинаю тест...")
	var test []string
	for i := 0; i < 5; i++ {
		quest := fmt.Sprintf(`//div[@class="flex flex-row overflow-x-auto"]/div[%d]`, i+1)
		page.MustElementX(quest).MustClick()
		question := page.MustElementX("//span[@class=\"text-unix-text-black dark:text-[#AFB7CA] font-medium text-2xl\"]").MustText()
		options := page.MustElementsX("//p[@class=\"ml-4\"]")
		optionsList := ""
		for i, opt := range options {
			optionsList += fmt.Sprintf("%d. %s\n", i+1, opt.MustText())
		}
		test = append(test, fmt.Sprintf("Вопрос №%d: %s\nВарианты:\n%s\n", i+1, question, optionsList))
	}
	var answers []int
	answers = AskGemini_Web(test[:]) //тут либо AskGemini_Web - без API или AskGemini - c API
	for i := 0; i < 5; i++ {
		answer := answers[i]
		log.Printf("Вопрос №%d\n", i+1)
		quest := fmt.Sprintf(`//div[@class="flex flex-row overflow-x-auto"]/div[%d]`, i+1)
		page.MustElementX(quest).MustClick()
		time.Sleep(1 * time.Second)
		log.Printf("ИИ выбрал: %d", answer)
		if answer == 0 {
			answer = rand.Intn(3) + 1
			log.Printf("ИИ не смог выбрать ответ, поэтому ответ был выбран случайно: %d\n", answer)
		}
		ans := fmt.Sprintf(`//div[@class="flex flex-col mt-5"]/div[%d]`, answer)
		page.MustElementX(ans).MustClick()
		log.Println("Успешно ответил!")
	}
	log.Println("Все вопросы пройдены. Завершаю тест...")
	page.MustElementX(`//button[@class="flex items-center text-white text-[15px] font-bold whitespace-nowrap max-sm:whitespace-normal"]`).MustClick()
	page.MustElementX(`//div[@class="flex-1 py-3 mr-1 border border-[#5F6B88] text-[#5F6B88] rounded-lg cursor-pointer text-[15px] font-bold text-center"]`).MustClick()
	grade := page.MustElementX(`//div[@class="flex items-center max-sm:justify-center"]/p/span[1]`).MustText()
	fmt.Printf("Оценка за тест: %s\n", grade)
}
