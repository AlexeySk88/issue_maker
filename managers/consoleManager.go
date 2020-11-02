package managers

import (
	"fmt"
	"github.com/gookit/color"
	"issue_maker/entities"
	"issue_maker/helpers"
	"os"
	"strings"
)

const createMessage = "\t* создана новая задача с заголовком:"
const updateMessage = "\t* обновлена задача с заголовком:"

var ErrorConsole = color.New(color.FgRed)
var InfoConsole = color.New(color.FgGreen)
var WarnConsole = color.New(color.FgYellow)

func CheckRequest(request *entities.Request) error {
	doPrint(request)
	fmt.Print("Начать запись в gitlab? (y/n): ")
	var userResponse string
	if _, err := fmt.Fscan(os.Stdin, &userResponse); err != nil {
		return fmt.Errorf("ошибка ввода, запись в gitlab не совершена")
	}

	success := []string{"y", "yes", "д", "да"}
	if !helpers.Contains(success, strings.TrimSpace(strings.ToLower(userResponse))) {
		return fmt.Errorf("прекращено пользователем, запись в gitlab не совершена")
	}
	return nil
}

func doPrint(request *entities.Request) {
	var create []string
	var update []string
	var unknown []string
	for _, i := range request.Issues {
		if i.IsCreate() {
			create = append(create, fmt.Sprintf("%s %s", createMessage, i.Title))
		} else if i.IsUpdate() {
			update = append(update, fmt.Sprintf("%s %s", updateMessage, i.Title))
		} else {
			unknown = append(unknown, fmt.Sprintf("* задача с заголовком %s не может быть отправлена в gitlab", i.Title))
		}
	}

	fmt.Println("Будут выполнены следующие действия:")
	for _, c := range create {
		InfoConsole.Println(c)
	}
	for _, u := range update {
		WarnConsole.Println(u)
	}
	for _, u := range unknown {
		ErrorConsole.Println(u)
	}
}
