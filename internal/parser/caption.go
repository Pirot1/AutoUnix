package parser

import (
	"AutoUnix/internal/ai"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/go-rod/rod"
)

func Clean_subs(subtitle string) string {
	lines := strings.Split(subtitle, "\n")
	var cleanLines []string
	timeRegex := regexp.MustCompile(`\d{2}:\d{2}:\d{2}\.\d{3} --> \d{2}:\d{2}:\d{2}\.\d{3}`)
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" || strings.HasPrefix(line, "WEBVTT") || timeRegex.MatchString(line) {
			continue
		}
		if _, err := fmt.Sscanf(line, "%d", new(int)); err == nil && !strings.Contains(line, " ") {
			continue
		}
		cleanLines = append(cleanLines, line)
	}
	return strings.Join(cleanLines, "\n")
}

func Caption_recorder(page *rod.Page, lessonName string) {
	el, err := page.Element("track[kind='captions']")
	if err != nil {
		log.Println("Субтитры не найдены на этой странице")
		return
	}
	re := regexp.MustCompile(`[<>:"/\|?*]`)
	name := re.ReplaceAllString(lessonName, "")
	folderPath := filepath.Join("lessons", strings.TrimSpace(name))
	err = os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatal("Ошибка создания папки:", err)
	}
	filePath := filepath.Join(folderPath, "lesson_summary.txt")
	subtitleURL := el.MustAttribute("src")
	if subtitleURL == nil || *subtitleURL == "" {
		log.Fatal("У тега track нет ссылки src")
	}
	log.Println("Начал запись субтитров...")
	content := page.MustEval(`(url) => fetch(url).then(res => res.text())`, *subtitleURL).String()
	err = os.WriteFile(filePath, []byte(Clean_subs(content)), 0644)
	if err != nil {
		log.Fatalf("Ошибка при сохранении субтитров: %v\n", err)
	} else {
		log.Printf("Файл сохранен: %s\n", filePath)
	}

	// Ai-power caption
	aiText := ai.Make_AI_conspect(content)
	filePath = filepath.Join(folderPath, "lesson_AI_summary.txt")
	err = os.WriteFile(filePath, []byte(Clean_subs(aiText)), 0644)
	if err != nil {
		log.Fatalf("Ошибка при сохранении субтитров: %v\n", err)
	} else {
		log.Printf("Файл сохранен: %s\n", filePath)
	}
}
