package managers

import (
	"fmt"
	"github.com/gookit/color"
	"issue_maker/entities"
	"issue_maker/helpers"
	"strings"
)

const createMessage = "\t* создана новая задача с заголовком:"
const updateMessage = "\t* обновлена задача с заголовком:"

var ErrorConsole = color.New(color.FgRed)
var InfoConsole = color.New(color.FgGreen)
var WarnConsole = color.New(color.FgYellow)

type ConsoleManager struct {
	scan Scan
}

type Scan func() (string, error)

func NewConsoleManager(s Scan) *ConsoleManager {
	return &ConsoleManager{scan: s}
}

func (cm *ConsoleManager) CheckRequest(request *entities.Request) error {
	cm.doPrint(request)
	fmt.Print("Начать запись в gitlab? (y/n): ")
	userResponse, err := cm.scan()
	if err != nil {
		return fmt.Errorf("ошибка ввода, запись в gitlab не совершена")
	}

	success := []string{"y", "yes", "д", "да"}
	if !helpers.ArrayContains(success, strings.TrimSpace(strings.ToLower(userResponse))) {
		return fmt.Errorf("прекращено пользователем, запись в gitlab не совершена")
	}
	return nil
}

func (cm *ConsoleManager) ReadConsole() string {
	input, _ := cm.scan()
	return strings.TrimSpace(input)
}

func  (cm *ConsoleManager) doPrint(request *entities.Request) {
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
