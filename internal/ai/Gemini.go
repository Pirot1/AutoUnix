package ai

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
)

const GeminiURL = "https://generativelanguage.googleapis.com/v1beta/models/gemini-1.5-flash:generateContent?key="

type GeminiResponse struct {
	Candidates []struct {
		Content struct {
			Parts []struct {
				Text string `json:"text"`
			} `json:"parts"`
		} `json:"content"`
	} `json:"candidates"`
}

func AskGemini(question string, options []string) (ans int) {
	err := godotenv.Load("../../.env") //Загрузка .env
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла")
	}
	GeminiAPI := os.Getenv("AI_KEY")
	optionsList := ""
	for i, opt := range options {
		optionsList += fmt.Sprintf("%d. %s\n", i+1, opt)
	}
	prompt := fmt.Sprintf(
		"Ты — помощник в обучении. Твоя задача: выбрать один правильный ответ из предложенного списка. "+
			"Ответь ТОЛЬКО номером от 1 до 4 выбранного варианта, без объяснений, без знаков препинания в конце и без текста.\n\n"+
			"Вопрос: %s\nВарианты:\n%v",
		question, optionsList,
	)
	requestBody, _ := json.Marshal(map[string]interface{}{
		"contents": []interface{}{
			map[string]interface{}{
				"parts": []interface{}{
					map[string]interface{}{"text": prompt},
				},
			},
		},
	})
	resp, err := http.Post(GeminiURL+GeminiAPI, "application/json", bytes.NewBuffer(requestBody))
	if err == nil {
		fmt.Println("Ошибка запроса к Gemini:", err)
		return 0
	}
	defer resp.Body.Close()
	body, _ := io.ReadAll(resp.Body)
	var geminiResp GeminiResponse
	json.Unmarshal(body, &geminiResp)

	rawResponse := geminiResp.Candidates[0].Content.Parts[0].Text
	var index int
	_, err = fmt.Sscanf(rawResponse, "%d", &index)
	if err != nil {
		fmt.Println("Ошибка: ИИ вернул не число:", rawResponse)
		return 0
	}
	return index
}
