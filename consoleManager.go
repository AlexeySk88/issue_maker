package main

import (
	"fmt"
	"github.com/gookit/color"
	"os"
	"strings"
)

const createMessage = "\t* создана новая задача с заголовком:"
const updateMessage = "\t* обновлена задача с заголовком:"

var errorConsole = color.New(color.FgRed)
var infoConsole = color.New(color.FgGreen)
var warnConsole = color.New(color.FgYellow)

func checkRequest(request *Request) error {
	doPrint(request)
	fmt.Print("Начать запись в gitlab? (y/n): ")
	var userResponse string
	if _, err := fmt.Fscan(os.Stdin, &userResponse); err != nil {
		return fmt.Errorf("ошибка ввода, запись в gitlab не совершена")
	}

	success := []string{"y", "yes", "д", "да"}
	if !contains(success, strings.TrimSpace(strings.ToLower(userResponse))) {
		return fmt.Errorf("прекращено пользователем, запись в gitlab не совершена")
	}
	return nil
}

func doPrint(request *Request) {
	var create []string
	var update []string
	var unknown []string
	for _, i := range request.Issues {
		if i.isCreate() {
			create = append(create, fmt.Sprintf("%s %s", createMessage, i.Title))
		} else if i.isUpdate() {
			update = append(update, fmt.Sprintf("%s %s", updateMessage, i.Title))
		} else {
			unknown = append(unknown, fmt.Sprintf("* задача с заголовком %s не может быть отправлена в gitlab", i.Title))
		}
	}

	fmt.Println("Будут выполнены следующие действия:")
	for _, c := range create {
		infoConsole.Println(c)
	}
	for _, u := range update {
		warnConsole.Println(u)
	}
	for _, u := range unknown {
		errorConsole.Println(u)
	}
}
