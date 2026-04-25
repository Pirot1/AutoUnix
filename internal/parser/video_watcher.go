package parser

import (
	"fmt"
	"log"
	"time"

	"AutoUnix/internal/ai"

	"github.com/go-rod/rod"
)

const (
	test string = "//div[@class=\"absolute video-overlay w-full h-full flex-col items-center text-center justify-center flex\"]//button[@class=\"bg-[#FFDD33] text-black font-bold rounded mt-3 px-2 py-1\"]"
)

func Proceed_lesson(page *rod.Page) {
	log.Println("Player initialisation...")
	page.MustElement(".plyr").MustWaitVisible()
	btn, err := page.ElementX("//button[@class=\"bg-[#FFDD33] text-black font-bold rounded mt-3 px-2 py-1\"]")
	if err == nil && btn.MustVisible() {
		log.Println("Continue watching lesson")
		btn.MustClick()
	} else {
		log.Println("Start wathcing lesson")
		btn = page.MustActivate().MustElementX("//button[@class=\"plyr__control plyr__control--overlaid\"]")
		btn.MustClick()
	}

	time.Sleep(500 * time.Millisecond)
	page.MustEval(`() => {
		let v = document.getElementById('video');
		if (v && v.paused) {
			v.muted = true;
			v.play().catch(e => console.log("Browser blocked launch:", e));
		}
	}`)
	page.MustEval(`() => {
		const video = document.querySelector('video');
		if (video) {
			video.muted = true;
			video.volume = 0;
		}
	}`)
	waitForVideoEnd(page)
}

func waitForVideoEnd(page *rod.Page) {
	log.Println("Waiting for metadata...")
	page.MustElement("video").WaitStable(time.Second)
	lessonTitle := page.MustElementX(`//span[@class="text-unix-text-black font-semibold dark:text-white"]`).MustText()
	log.Println("Урок:", lessonTitle)
	Caption_recorder(page, lessonTitle)
	for {
		result, err := page.Eval(`() => {
			let v = document.getElementById('video');
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
			log.Println("\nConnection error")
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

		fmt.Printf("\rDuration: %d sec. out of %d sec. 			%.0f%%", current, total, (float32(current)/float32(total))*100)

		if ended || (total > 0 && current >= total) {
			log.Println("\nSuccessfully complete lesson!")
			HadlePostVideoActions(page)
			break
		}

		time.Sleep(1 * time.Second)
	}
}

func HadlePostVideoActions(page *rod.Page) {
	log.Println("Analizing...")
	res, err := page.Eval(`() => {
    const btn = Array.from(document.querySelectorAll('button')).find(b => b.innerText.includes('Go to test'));
    if (btn) {
        btn.click();
        return true;
    }
    return false;
	}`)

	if err == nil && res.Value.Bool() {
		log.Println("Found test")
		ai.SolvingQuiz(page)
	} else {
		log.Println("Test was not found")
		time.Sleep(1 * time.Second)
	}
}
