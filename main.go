package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"issue_maker/entities"
	"issue_maker/managers"
	"net/http"
	"runtime/debug"
	"time"
)

func main() {
	cm := managers.NewConsoleManager(func() (string, error) {
		var str string
		_, err := fmt.Scan(&str)
		return str, err
	})

	fm := managers.NewFileManager(afero.NewOsFs())
	logFile, err := fm.GetFileLog()
	if err != nil {
		managers.ErrorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	initLogs(logFile)
	req, err := getRequest(fm, cm)
	if err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "File read error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		time.Sleep(time.Second * 5)
		return
	}
	rm := managers.NewRestManager(req.ProjectId, req.AccessToken, &http.Client{}, fm)
	im := managers.NewIssueManager(rm, fm, req)

	milestones, err := rm.GetMilestones()
	if err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Milestones get error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = req.UpdateMilestone(milestones); err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Update issues error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = cm.CheckRequest(req); err != nil {
		managers.ErrorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	req, err = im.Send()
	if err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Gitlab Send error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = fm.WriteDoneFile(req); err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "File write error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	time.Sleep(time.Second * 5)
}

func initLogs(file afero.File) {
	log.SetOutput(file)
	log.SetLevel(log.ErrorLevel)
}

func getRequest(fm *managers.FileManager, cm *managers.ConsoleManager) (*entities.Request, error) {
	req, err := fm.ReadIssuesFile()
	if err != nil {
		managers.ErrorConsole.Println(err)
	} else {
		return req, nil
	}
	managers.ErrorConsole.Printf("Введите путь к файлу с задачами: ")
	reqPath := cm.ReadConsole()
	return fm.ReadIssuesFileFromPath(reqPath)
}
