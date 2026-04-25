package ai

import (
	"context"
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	"google.golang.org/genai"
)

func AskGemini(question string, options []string) int {
	err := godotenv.Load(".env") //Loading .env
	if err != nil {
		log.Fatal("Ошибка при загрузке .env файла\n")
	}
	GeminiAPI := os.Getenv("AI_KEY")
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

	ctx := context.Background()
	client, err := genai.NewClient(ctx, &genai.ClientConfig{
		APIKey:  GeminiAPI,
		Backend: genai.BackendGeminiAPI,
	},
	)
	if err != nil {
		log.Printf("Ошибка создания клиента: %v\n", err)
		return 0
	}
	result, err := client.Models.GenerateContent(
		ctx,
		"gemma-3-12b-it",
		genai.Text(prompt),
		nil,
	)
	if err != nil {
		log.Printf("Ошибка при создании ответа: %v\n", err)
	}
	int_result, _ := strconv.Atoi(result.Text())
	return int_result
}
