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

func AskGemini_Web(test []string) []int {
	b, page := browser.Init("https://gemini.google.com/app", true) // потом поставить false
	defer b.MustClose()
	log.Println("Успешно запустил Gemini")
	var result []string
	for i := 0; i < len(test); i++ {
		prompt := fmt.Sprintf(
			"Act as a quiz solver. Output ONLY the answer key for the following questions.\n"+
				"Strictly NO text, NO explanations, NO markdown.\n\n"+
				"%s",
			test[i],
		)
		page.MustElementX(`//div[@role="textbox"]`).MustInput(prompt)
		page.KeyActions().Press(input.Enter).MustDo()
		log.Println("Успешно ввёл запрос...")
		deadline := time.Now().Add(60 * time.Second)
		for time.Now().Before(deadline) {
			found, el, _ := page.Has(".markdown-main-panel")
			if found {
				currentText := el.MustProperty("innerText").String()
				currentText = strings.TrimSpace(currentText)
				if len(currentText) == 1 {
					log.Printf("Ответ: %s\n", currentText)
					result = append(result, currentText)
					break
				}
			}
			time.Sleep(1 * time.Second)
		}
		page.Reload()
	}
	if len(result) == 0 {
		return []int{0, 0, 0, 0, 0}
	}
	var total []int
	for _, el := range result {
		num, err := strconv.Atoi(el)
		if err != nil {
			log.Println("Unable to convert")
			num = 0
		}
		total = append(total, num)
	}
	return total
}

func Make_AI_conspect(txt string) string {
	b, page := browser.Init("https://gemini.google.com/app", true) // потом поставить false
	defer b.MustClose()
	log.Println("Успешно запустил Gemini")

	prompt := fmt.Sprintf(
		"Ты — помощник в обучении. Твоя задача: написать конспект на английском языке изходя из \"сырого\" файла конспекта:\n%s", txt,
	)
	page.MustElementX(`//div[@role="textbox"]`).MustInput(prompt)
	page.KeyActions().Press(input.Enter).MustDo()
	log.Println("Ввёл вопрос. Жду ответ...")
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
					log.Println("Обнаружено зависание статуса, но текст готов. Забираю.")
				}
				result = currentText
				break
			}
		}
		time.Sleep(1 * time.Second)
	}
	return result
}
