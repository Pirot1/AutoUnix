package parser

import (
	"fmt"

	"github.com/go-rod/rod"
)

func GetAvailableLessons(page *rod.Page, code string) []string {
	chapters := page.MustActivate().MustElementsX("//div[@class=\"mt-5 rounded-[28px] bg-white dark:bg-[#1a1a1a] p-6 flex-0 overflow-y-auto\"]//div[@class=\"flex flex-row items-start cursor-pointer\"]")
	fmt.Println("Открываю главы...")
	if len(chapters) > 1 {
		for _, chapter := range chapters[1:] {
			chapter.MustClick()
		}
	}
	lessons := page.MustActivate().MustElementsX("//a[@class = \"flex flex-row items-center cursor-pointer\"]")
	fmt.Println("Ищу урок...")
	var urls []string
	for _, lesson := range lessons {
		img, err := lesson.Element(".w-4.h-4")
		if err != nil {
			continue
		}
		src := img.MustAttribute("src")
		if src == nil {
			continue
		}
		if *src == code {
			continue
		} else {
			fmt.Println("Урок найден!")
			urls = append(urls, "https://uni-x.almv.kz"+*lesson.MustAttribute("href"))
		}
	}
	return urls
}
