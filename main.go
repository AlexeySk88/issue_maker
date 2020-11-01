package main

import "fmt"

func main() {
	req, err := fileRead()
	if err != nil {
		errorConsole.Println(err)
		return
	}

	milestones, err := getMilestones(req)
	if err != nil {
		errorConsole.Println(err)
		return
	}

	if err = req.updateMilestone(milestones); err != nil {
		errorConsole.Println(err)
		return
	}

	if err = checkRequest(req); err != nil {
		errorConsole.Println(err)
		return
	}

	req, err = send(req)
	if err != nil {
		errorConsole.Println(err)
		return
	}

	if err = fileWrite(req); err != nil {
		fmt.Println(err)
		return
	}
}
