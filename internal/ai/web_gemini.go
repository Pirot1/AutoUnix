package ai

import (
	"fmt"
	"strconv"
	"time"

	"AutoUnix/internal/browser"

	"github.com/go-rod/rod/lib/input"
)

func AskGemini_Web(question string, options []string) int {
	b, page := browser.Init("https://gemini.google.com/app", true) // потом поставить false
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
	page.KeyActions().Press(input.Enter).MustDo()
	fmt.Println("Ввёл вопрос. Жду ответ...")
	page.MustElementX(`//div[@class="container"]`).MustVisible()
	time.Sleep(500 * time.Millisecond)
	result := page.MustElementX(`//div[@class="container"]`).MustText()
	time.Sleep(500 * time.Millisecond)
	res, _ := strconv.Atoi(result)
	return res
}

func Make_AI_conspect(txt string) string {
	b, page := browser.Init("https://gemini.google.com/app", false) // потом поставить false
	defer b.MustClose()
	fmt.Println("Успешно запустил Gemini")

	prompt := fmt.Sprintf(
		"Ты — помощник в обучении. Твоя задача: написать конспект на английском языке изходя из \"сырого\" файла конспекта:\n%s", txt,
	)
	page.MustElementX(`//div[@role="textbox"]`).MustInput(prompt)
	page.KeyActions().Press(input.Enter).MustDo()
	fmt.Println("Ввёл вопрос. Жду ответ...")
	page.MustElementX(`//div[@class="container"]`).MustVisible()
	time.Sleep(500 * time.Millisecond)
	result := page.MustElementX(`//div[@class="markdown markdown-main-panel stronger enable-updated-hr-color" ]`).MustText()
	time.Sleep(500 * time.Millisecond)
	return result
}
