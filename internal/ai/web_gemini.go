package ai

import (
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"AutoUnix/internal/browser"

	"github.com/go-rod/rod/lib/input"
)

func AskGemini_Web(question string, options []string) int {
	b, page := browser.Init("https://gemini.google.com/app", true) // потом поставить false
	defer b.MustClose()
	log.Println("Successfully init Gemini")

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
	deadline := time.Now().Add(60 * time.Second)
	var result string
	for time.Now().Before(deadline) {
		found, el, _ := page.Has(".markdown-main-panel")
		if found {
			currentText := el.MustProperty("innerText").String()
			currentText = strings.TrimSpace(currentText)
			if len(currentText) == 1 {
				result = currentText
			} else {
				continue
			}
			busy, _ := el.Attribute("aria-busy")
			if len(currentText) == 1 {
				log.Println("Delay was founded. Take the raw answer.")
				result = currentText
				res, _ := strconv.Atoi(result)
				return res
			}
			if busy != nil && *busy == "false" {
				result = currentText
				res, _ := strconv.Atoi(result)
				return res
			}
		}
		time.Sleep(1 * time.Second)
	}
	return 0
}

func Make_AI_conspect(txt string) string {
	b, page := browser.Init("https://gemini.google.com/app", true) // потом поставить false
	defer b.MustClose()
	log.Println("Successfully init Gemini")

	prompt := fmt.Sprintf(
		"Ты — помощник в обучении. Твоя задача: написать конспект на английском языке изходя из \"сырого\" файла конспекта:\n%s", txt,
	)
	page.MustElementX(`//div[@role="textbox"]`).MustInput(prompt)
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Ask question. waiting for an answer...")
	deadline := time.Now().Add(60 * time.Second)
	var result string
	var lastText string
	var sameCount int
	for time.Now().Before(deadline) {
		found, el, _ := page.Has(".markdown-main-panel")
		if found {
			currentText := el.MustProperty("innerText").String()
			currentText = strings.TrimSpace(currentText)
			if len(currentText) > 0 && currentText == lastText {
				sameCount++
			} else {
				sameCount = 0
			}
			lastText = currentText
			busy, _ := el.Attribute("aria-busy")
			if (busy != nil && *busy == "false") || (sameCount >= 7 && len(currentText) > 50) {
				if sameCount >= 7 {
					log.Println("Delay was founded. Take the raw answer.")
				}
				result = currentText
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
	return result
}
