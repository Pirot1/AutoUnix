package browser

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
)

func Init(login string, password string) (*rod.Browser, *rod.Page) {
	l := launcher.New().Headless(false).Devtools(false).Leakless(false).Set("autoplay-policy", "no-user-gesture-required")
	url, err := l.Launch()
	if err != nil {
		panic(fmt.Sprintf("Не удалось запустить браузер: %v", err))
	}
	browser := rod.New().ControlURL(url).MustConnect().NoDefaultDevice()
	//defer browser.MustClose()

	page := browser.MustPage("https://uni-x.almv.kz/platform/login")
	fmt.Println("Начинаем работу...")
	return browser, page
}
