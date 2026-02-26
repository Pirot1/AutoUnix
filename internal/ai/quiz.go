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
		question := page.MustElement("p.select-none.mb-5 span:nth-child(2)").MustText()
		options := page.MustElements("div.flex.flex-row.cursor-pointer")
		var quiz []string
		for _, el := range options {
			quiz = append(quiz, el.MustElement("p.ml-4").MustText())
		}
		fmt.Printf("%s\n", question)
		fmt.Printf("%v\n", options)
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
