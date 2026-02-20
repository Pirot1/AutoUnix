package parser

import (
	"fmt"

	"github.com/go-rod/rod"
)

func GetAvailableLessons(page *rod.Page) {
	chapters := page.MustElements(".mt-5 rounded-[28px] bg-white dark:bg-[#1a1a1a] p-6 flex-0 overflow-y-auto") // Замени на реальный класс заголовка
	fmt.Println(chapters)
	for _, chapter := range chapters {
		chapter.MustClick()
		// 2. Теперь ищем уроки, которые стали видимыми
		lessons := page.MustElements(".flex flex-col a") // Пример класса урока
		fmt.Println(lessons)
		for _, lesson := range lessons {
			fmt.Printf("Успешно зашёл на видео!")
			// 3. Проверяем иконку (кружок слева)
			// Если кружок пустой — значит урок не пройден
			// Если там SVG с галочкой — пройден
			isDone, _, _ := lesson.Has("svg") // Если галочка — это SVG внутри круга

			if isDone {
				fmt.Println("✅ Пропускаем:", lesson.MustText())
				continue
			}

			// 4. Забираем ссылку и название
			url := lesson.MustElement("a").MustAttribute("href")
			fmt.Printf("🎯 Нашелся целевой урок: %s (URL: %s)\n", lesson.MustText(), *url)

			// Здесь можно либо сразу кликать, либо сохранить в список
		}
	}
}
