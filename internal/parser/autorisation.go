package parser

import (
	"fmt"

	"github.com/go-rod/rod"
)

func Autorisation(page *rod.Page, login string, password string) {
	page.MustElement("input[type=\"email\"]").MustWaitVisible().MustInput(login)
	fmt.Println("Login succses")
	page.MustElement("input[type=\"password\"]").MustWaitVisible().MustInput(password)
	fmt.Println("Password succses")
	page.MustElement("button[type=\"submit\"]").MustClick()
	// Ждем загрузки личного кабинета
	page.MustWaitLoad()
	fmt.Println("Успешный вход!")
}
