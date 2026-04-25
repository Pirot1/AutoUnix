package parser

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/go-rod/rod"
)

func Autorisation(page *rod.Page, login string, password string) {
	page.MustElement("input[type=\"email\"]").MustWaitVisible().MustInput(login)
	page.MustElement("input[type=\"password\"]").MustWaitVisible().MustInput(password)
	page.MustElement("button[type=\"submit\"]").MustClick()
	page.MustWaitLoad()
	fmt.Println("Successfully autorisate!")
}

func NewEnvFile() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Input login: ")
	login, _ := reader.ReadString('\n')
	login = strings.TrimSpace(login)

	fmt.Println("Input password: ")
	password, _ := reader.ReadString('\n')
	password = strings.TrimSpace(password)

	fmt.Println("Input your course: ")
	course, _ := reader.ReadString('\n')
	course = strings.TrimSpace(course)

	envData := fmt.Sprintf("USER_EMAIL=%s\nUSER_PASS=%s\nLESSON_NAME=%s", login, password, course)
	err := os.WriteFile(".env", []byte(envData), 0644)
	if err != nil {
		fmt.Println("Writing error")
	} else {
		fmt.Println("Successfully create .env!")
	}
}
