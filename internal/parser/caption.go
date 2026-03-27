package parser

import (
	"AutoUnix/internal/ai"
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"github.com/dslipak/pdf"
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
		fmt.Println("Субтитры не найдены на этой странице")
		return
	}
	re := regexp.MustCompile(`[<>:"/\|?*]`)
	name := re.ReplaceAllString(lessonName, "")
	folderPath := filepath.Join("lessons", strings.TrimSpace(name))
	err := os.MkdirAll(folderPath, 0755)
	if err != nil {
		log.Fatal("Ошибка создания папки:", err)
	}
	filePath := filepath.Join(folderPath, "lesson_summary.txt")
	subtitleURL := el.MustAttribute("src")
	if subtitleURL == nil || *subtitleURL == "" {
		fmt.Println("У тега track нет ссылки src")
		return
	}
	fmt.Println("Начал запись субтитров...")
	content := page.MustEval(`(url) => fetch(url).then(res => res.text())`, *subtitleURL).String()
	err = os.WriteFile(filePath, []byte(Clean_subs(content)), 0644)
	if err != nil {
		log.Fatalf("Ошибка при сохранении субтитров: %v\n", err)
	} else {
		log.Printf("Файл сохранен: %s\n", filePath)
	}
	// Ai-power caption
	aiText := ai.Make_AI_conspect(fullText)
	filePath = filepath.Join(folderPath, "lesson_AI_summary.txt")
	err = os.WriteFile(filePath, []byte(Clean_subs(aiText)), 0644)
	if err != nil {
		log.Fatalf("Ошибка при сохранении субтитров: %v\n", err)
	} else {
		log.Printf("Файл сохранен: %s\n", filePath)
	}
}
func ReadPDF(url string, lessonName string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Не удалось скачать PDF: %v", err)
		return
	}
	defer resp.Body.Close()
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("Не удалось прочитать тело ответа: %v", err)
		return
	}
	r, err := pdf.NewReader(bytes.NewReader(data), int64(len(data)))
	if err != nil {
		fmt.Printf("Ошибка инициализации PDF ридера: %v", err)
		return
	}
	textReader, err := r.GetPlainText()
	if err != nil {
		return
	}
	var buf bytes.Buffer
	_, err = buf.ReadFrom(textReader)
	if err != nil {
		return
	}
	Make_conspect(lessonName, buf.String())
}
func Check_materials(page *rod.Page, lessonName string) {
	exists, el, err := page.Has("span[title='Materials']")
	if err != nil {
		fmt.Printf("Ошибка при поиске: %v\n", err)
		return
	}
	if !exists {
		fmt.Println("Материалов на этом уроке нет.")
		return
	}
	el.MustClick()
	pdfEl, err := page.Timeout(2 * time.Second).Element("a[href$='.pdf']")
	if err == nil {
		pdfURL := pdfEl.MustAttribute("href")
		fmt.Printf("Нашел PDF-конспект: %s\n", *pdfURL)
		ReadPDF(*pdfURL, lessonName)
		return
	}
}
func Caption_recorder(page *rod.Page, lessonName string) {
	exists, el, err := page.Has("track[kind='captions']")
	if err != nil {
		fmt.Printf("Ошибка при поиске: %v\n", err)
		return
	}
	if !exists {
		fmt.Println("Субтитров на этом уроке нет.")
		Check_materials(page, lessonName)
		return
	}
	subtitleURL := el.MustAttribute("src")
	if subtitleURL == nil || *subtitleURL == "" {
		fmt.Println("У тега track нет ссылки src")
		return
	}
	fmt.Println("Начал запись субтитров...")
	content := page.MustEval(`(url) => fetch(url).then(res => res.text())`, *subtitleURL).String()
	Make_conspect(lessonName, content)
}
