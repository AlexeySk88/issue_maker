package main

import (
	log "github.com/sirupsen/logrus"
	"os"
	"runtime/debug"
	"time"
)

func main() {
	logFile, err := getFileLog()
	if err != nil {
		errorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	initLogs(logFile)
	req, err := fileRead()
	if err != nil {
		errorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "File read error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	milestones, err := getMilestones(req)
	if err != nil {
		errorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Milestones get error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = req.updateMilestone(milestones); err != nil {
		errorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Update issues error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = checkRequest(req); err != nil {
		errorConsole.Println(err)
		time.Sleep(time.Second * 5)
		return
	}

	req, err = send(req)
	if err != nil {
		errorConsole.Println(err)
		log.WithFields(log.Fields{
			"title":       "Gitlab send error",
			"stack_trace": string(debug.Stack()),
		}).Error(err)
		return
	}

	if err = fileWrite(req); err != nil {
		errorConsole.Println(err)
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
