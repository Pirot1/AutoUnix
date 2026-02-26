package parser

import (
	"fmt"
	"time"

	"github.com/go-rod/rod"
)

func Proceed_lesson(page *rod.Page) {
	fmt.Println("Инициализация плеера...")
	page.MustElement(".plyr").MustWaitVisible()
	btn, err := page.ElementX("//button[@class=\"bg-[#FFDD33] text-black font-bold rounded mt-3 px-2 py-1\"]")
	if err == nil && btn.MustVisible() {
		fmt.Println("Продолжаю смотреть урок")
		btn.MustClick()
	} else {
		fmt.Println("Начинаю смотреть урок")
		btn = page.MustActivate().MustElementX("//button[@class=\"plyr__control plyr__control--overlaid\"]")
		btn.MustClick()
	}

	time.Sleep(500 * time.Millisecond)
	page.MustEval(`() => {
		let v = document.getElementById('video');
		if (v && v.paused) {
			v.muted = true; // Снимаем блокировку браузера
			v.play().catch(e => console.log("Браузер заблокировал запуск:", e));
		}
	}`)
	waitForVideoEnd(page)
}

func waitForVideoEnd(page *rod.Page) {
	fmt.Println("Жду загрузки метаданных видео...")
	page.MustElement("video").WaitStable(time.Second)
	for {
		result, err := page.Eval(`() => {
			let v = document.getElementById('video');
			
			// Если видео нет или readyState === 0 (метаданные не загружены)
			if (!v || v.readyState === 0) {
				return { ready: false, ended: false, current: 0, total: 0 };
			}
			
			return {
				ready: true,
				ended: v.ended,
				current: Math.floor(v.currentTime),
				total: Math.floor(v.duration)
			};
		}`)

		if err != nil {
			fmt.Println("\nОшибка связи с плеером. Возможно, страница перезагрузилась.")
			break
		}
		ready := result.Value.Get("ready").Bool()
		if !ready {
			time.Sleep(1 * time.Second)
			continue
		}

		ended := result.Value.Get("ended").Bool()
		current := result.Value.Get("current").Int()
		total := result.Value.Get("total").Int()

		fmt.Printf("\rВоспроизведение: %d сек. из %d сек.", current, total)

		if ended || (total > 0 && current >= total-1) {
			fmt.Println("\nВидео успешно досмотрено!")
			break
		}

		time.Sleep(3 * time.Second)
	}
}
