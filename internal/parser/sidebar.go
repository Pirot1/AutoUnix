package parser

import (
	"fmt"

	"github.com/go-rod/rod"
)

const (
	code = "data:image/svg+xml;base64,PHN2ZyB3aWR0aD0iMTYiIGhlaWdodD0iMTciIHZpZXdCb3g9IjAgMCAxNiAxNyIgZmlsbD0ibm9uZSIgeG1sbnM9Imh0dHA6Ly93d3cudzMub3JnLzIwMDAvc3ZnIj4KPHBhdGggZD0iTTE2IDguNUMxNiAxMi45MTgzIDEyLjQxODMgMTYuNSA4IDE2LjVDMy41ODE3MiAxNi41IDAgMTIuOTE4MyAwIDguNUMwIDQuMDgxNzIgMy41ODE3MiAwLjUgOCAwLjVDMTIuNDE4MyAwLjUgMTYgNC4wODE3MiAxNiA4LjVaTTEuOTkyIDguNUMxLjk5MiAxMS44MTgxIDQuNjgxODcgMTQuNTA4IDggMTQuNTA4QzExLjMxODEgMTQuNTA4IDE0LjAwOCAxMS44MTgxIDE0LjAwOCA4LjVDMTQuMDA4IDUuMTgxODcgMTEuMzE4MSAyLjQ5MiA4IDIuNDkyQzQuNjgxODcgMi40OTIgMS45OTIgNS4xODE4NyAxLjk5MiA4LjVaIiBmaWxsPSIjQUZCN0NBIiBmaWxsLW9wYWNpdHk9IjAuNSIvPgo8ZWxsaXBzZSBjeD0iOCIgY3k9IjguNSIgcng9IjgiIHJ5PSI4IiB0cmFuc2Zvcm09InJvdGF0ZSgtOTAgOCA4LjUpIiBmaWxsPSIjM0JCQzMwIi8+CjxwYXRoIGQ9Ik00LjIwODUgOC40OTk4NEw2LjkxNjgzIDExLjIwODJMMTIuMzMzNSA1Ljc5MTUiIHN0cm9rZT0id2hpdGUiIHN0cm9rZS13aWR0aD0iMS41IiBzdHJva2UtbGluZWNhcD0icm91bmQiIHN0cm9rZS1saW5lam9pbj0icm91bmQiLz4KPC9zdmc+Cg=="
)

func GetAvailableLessons(page *rod.Page) []string {
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
