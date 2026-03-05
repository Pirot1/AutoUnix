package ai

import (
	"fmt"
	"strconv"

	"AutoUnix/internal/browser"
)

func AskGemini_Web(question string, options []string) int {
	b, page := browser.Init("https://gemini.google.com/app", false) // потом поставить false
	defer b.MustClose()
	fmt.Println("Успешно запустил Gemini")

	optionsList := ""
	for i, opt := range options {
		optionsList += fmt.Sprintf("%d. %s\n", i+1, opt)
	}
	prompt := fmt.Sprintf(
		"Ты — помощник в обучении. Твоя задача: выбрать один правильный ответ из предложенного списка. "+
			"Ответь ТОЛЬКО номером от 1 до 4 выбранного варианта, без объяснений, без знаков препинания в конце и без текста.\n\n"+
			"Вопрос: %s\nВарианты:\n%v ",
		question, optionsList,
	)
	page.MustElementX(`//div[@role="textbox"]`).MustInput(prompt)
	fmt.Println("Ввёл вопрос. Жду ответ...")
	result := page.MustWaitLoad().MustElement(`//p[@data-path-to-node="0"]`).MustText()
	res, _ := strconv.Atoi(result)
	return res
}
