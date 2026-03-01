package ai

import (
	"fmt"

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
		num_ans := 1
		for _, el := range options {
			if num_ans == answer {
				el.MustClick()
				break
			} else if answer == 0 {
				break
			} else {
				num_ans++
			}
		}
		page.MustElementR("button", "Next").MustClick()
	}
	fmt.Println("Все вопросы пройдены. Завершаю тест...")
}
