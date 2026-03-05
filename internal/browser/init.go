package browser

import (
	"fmt"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/stealth"
)

func Init(site string, head bool) (*rod.Browser, *rod.Page) {
	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`
	l := launcher.New().Bin(chromePath).Headless(false).Devtools(false).Leakless(false).Set("autoplay-policy", "no-user-gesture-required")
	url, err := l.Launch()
	if err != nil {
		panic(fmt.Sprintf("Не удалось запустить браузер: %v", err))
	}
	browser := rod.New().ControlURL(url).MustConnect().NoDefaultDevice()
	//defer browser.MustClose()
	page := stealth.MustPage(browser)
	page = page.MustNavigate(site)
	fmt.Println("Начинаем работу...")
	return browser, page
}
