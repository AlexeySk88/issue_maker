package managers

import (
	"encoding/json"
	"fmt"
	"issue_maker/entities"
	"net/http"
)

const baseUrl string = "https://gitlab.com/api/v4/projects"

func GetMilestones(request *entities.Request) (*[]entities.Milestone, error) {
	client := http.Client{}
	url := fmt.Sprintf(
		"%s/%d/milestones",
		baseUrl,
		request.ProjectId,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", request.AccessToken)
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ошибка при выполнении запроса %s, код %d", url, resp.StatusCode)
	}

	return responseMilestoneHandler(resp)
}

func Send(request *entities.Request) (*entities.Request, error) {
	newReq := entities.Request{
		AccessToken: request.AccessToken,
		ProjectId:   request.ProjectId,
		Milestone:   request.Milestone,
	}
	client := http.Client{}
	for index, issue := range request.Issues {
		var req *http.Request
		if issue.IsCreate() {
			req = create(request, index)
		} else if issue.IsUpdate() {
			req = update(request, index)
		} else {
			ErrorConsole.Printf("Задача с заголовком %s не будет отправлена в gitlab\n", issue.Title)
			continue
		}
		fmt.Printf("Записываю в gitlab задачу с заголовком: %s\n", issue.Title)

		req.Header.Add("PRIVATE-TOKEN", request.AccessToken)
		resp, err := client.Do(req)
		if err != nil {
			return nil, err
		}
		if resp.StatusCode > 300 {
			return nil, fmt.Errorf("ошибка при выполнении запроса %s, код %d", req.URL, resp.StatusCode)
		}

		newIssue := responseCreateHandler(resp, &issue)
		if newIssue != nil {
			newReq.Issues = append(newReq.Issues, *newIssue)
		}
	}

	fmt.Println("Задачи успешно записаны!")
	return &newReq, nil
}

func create(request *entities.Request, index int) *http.Request {
	url := fmt.Sprintf(
		"%s/%d/issues?%s",
		baseUrl,
		request.ProjectId,
		request.RequestParam(index),
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		ErrorConsole.Println(err)
	}
	return req
}

func update(request *entities.Request, index int) *http.Request {
	url := fmt.Sprintf(
		"%s/%d/issues/%d?%s",
		baseUrl,
		request.ProjectId,
		request.Issues[index].Id,
		request.RequestParam(index),
	)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		ErrorConsole.Println(err)
	}
	return req
}

func responseCreateHandler(resp *http.Response, i *entities.Issue) *entities.Issue {
	r := entities.Response{}
	defer resp.Body.Close()
	err := json.NewDecoder(resp.Body).Decode(&r)
	if err != nil {
		ErrorConsole.Println(err)
		return nil
	}

	return i.CopyForWrite(r.Iid)
}

func responseMilestoneHandler(resp *http.Response) (*[]entities.Milestone, error) {
	var m []entities.Milestone
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}
