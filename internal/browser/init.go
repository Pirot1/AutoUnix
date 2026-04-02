package browser

import (
	"log"

	"github.com/go-rod/rod"
	"github.com/go-rod/rod/lib/launcher"
	"github.com/go-rod/rod/lib/proto"
	"github.com/go-rod/stealth"
)

func Init(site string, head bool) (*rod.Browser, *rod.Page) {
	//Параметры инициализации
	chromePath := `C:\Program Files\Google\Chrome\Application\chrome.exe`
	l := launcher.New().Bin(chromePath).Headless(head).Devtools(false).Leakless(false)
	l.Set("autoplay-policy", "no-user-gesture-required")
	l.Set("disable-gpu")
	l.Set("no-sandbox")
	l.Set("disable-dev-shm-usage")
	l.Set("disable-extensions")
	l.Set("disable-dev-shm-usage")
	l.Set("disable-setuid-sandbox")
	l.Set("proxy-server", "direct://")
	l.Set("disable-software-rasterizer")
	url, err := l.Launch()
	if err != nil {
		log.Panicf("Не удалось запустить браузер: %v", err)
	}
	browser := rod.New().ControlURL(url).MustConnect().NoDefaultDevice()
	page := stealth.MustPage(browser)
	page.MustSetViewport(1920, 1080, 1, false)
	page.MustSetUserAgent(&proto.NetworkSetUserAgentOverride{
		UserAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/122.0.0.0 Safari/537.36",
	})
	_ = proto.NetworkEnable{}.Call(page)
	err = proto.NetworkSetBlockedURLs{
		Urls: []string{
			"*.png", "*.jpg", "*.jpeg", "*.gif", "*.svg",
			"*.woff", "*.woff2", "*.ttf",
			"*analytics*", "*metrika*", "*doubleclick*",
		},
	}.Call(page)
	if err != nil {
		log.Println("Ошибка настройки блокировки:", err)
	}
	page.MustNavigate(site)
	log.Println("Начинаем работу...")
	return browser, page
}
