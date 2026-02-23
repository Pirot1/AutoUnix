package parser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func Proceed_lesson(page *rod.Page) {
	btn, err := page.ElementX("//button[@class=\"bg-[#FFDD33] text-black font-bold rounded mt-3 px-2 py-1\"]")
	if err == nil {

	}
	if err == nil && btn.MustVisible() {
		fmt.Println("Продолжаю смотреть урок")
		btn.MustClick()
	} else {
		fmt.Println("Начинаю смотреть урок")
		btn = page.MustActivate().MustElementX("//button[@class=\"plyr__control plyr__control--overlaid\"]")
		btn.MustClick()
	}
	waitForVideoEnd(page)
}
func waitForVideoEnd(page *rod.Page) {
	fmt.Println("Видео запущено. Жду окончания...")
	for {
		ended := page.MustEval(`() => {
			const video = document.querySelector('video');
			if (!video) return false;
			return video.ended;
		}`).Bool()
		if ended {
			fmt.Println("Видео досмотрено до конца!")
			break
		}
		currentTime := page.MustEval(`() => document.querySelector('video').currentTime`).Int()
		fmt.Printf("\rТекущее время: %d сек.", currentTime)
		time.Sleep(5 * time.Second)
	}
}
