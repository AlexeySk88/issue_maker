package main

import (
	log "github.com/sirupsen/logrus"
	"issue_maker/managers"
	"os"
	"runtime/debug"
	"time"
)

func main() {
	logFile, err := managers.GetFileLog()
	if err != nil {
		managers.ErrorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	initLogs(logFile)
	req, err := managers.FileRead()
	if err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "File read error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		time.Sleep(time.Second * 5)
		return
	}

	milestones, err := managers.GetMilestones(req)
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

	if err = managers.CheckRequest(req); err != nil {
		managers.ErrorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	req, err = managers.Send(req)
	if err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Gitlab Send error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = managers.FileWrite(req); err != nil {
		managers.ErrorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "File write error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	time.Sleep(time.Second * 5)
}

func initLogs(file *os.File) {
	log.SetOutput(file)
	log.SetLevel(log.ErrorLevel)
}
