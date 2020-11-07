package managers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"issue_maker/entities"
	"mime/multipart"
	"net/http"
)

const baseUrl string = "https://gitlab.com/api/v4/projects"

type RestManager struct {
	projectId   int
	accessToken string
	client      *http.Client
	fm          *FileManager
}

func NewRestManager(projectId int, accessToken string, fm *FileManager) *RestManager {
	return &RestManager{projectId: projectId, accessToken: accessToken, client: &http.Client{}, fm: fm}
}

func (rm *RestManager) GetMilestones() (*[]entities.Milestone, error) {
	url := fmt.Sprintf(
		"%s/%d/milestones",
		baseUrl,
		rm.projectId,
	)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("PRIVATE-TOKEN", rm.accessToken)
	resp, err := rm.client.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		return nil, fmt.Errorf("ошибка при выполнении запроса %s, код %d", url, resp.StatusCode)
	}

	return rm.responseMilestoneHandler(resp)
}

func (rm *RestManager) UploadFile(path string, projectId int, accessToken string) (string, error) {
	file, err := rm.fm.GetFile(path)
	if err != nil {
		return "", err
	}

	body := &bytes.Buffer{}
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile("file", path)
	if err != nil {
		return "", err
	}

	if _, err = io.Copy(part, file); err != nil {
		return "", nil
	}
	if err = writer.Close(); err != nil {
		return "", nil
	}

	url := fmt.Sprintf("%s/%d/uploads",
		baseUrl,
		projectId)
	req, err := http.NewRequest("POST", url, body)
	if err != nil {
		return "", nil
	}
	req.Header.Add("PRIVATE-TOKEN", accessToken)
	req.Header.Set("Content-Type", writer.FormDataContentType())

	resp, err := rm.client.Do(req)
	if err != nil {
		return "", nil
	}

	uploads, err := rm.responseUploadsHandler(resp)
	if err != nil {
		return "", nil
	}

	return uploads.Markdown, nil
}

func (rm *RestManager) Create(request *entities.Request, index int) (int, error) {
	url := fmt.Sprintf(
		"%s/%d/issues?%s",
		baseUrl,
		rm.projectId,
		request.RequestParam(index),
	)
	req, err := http.NewRequest("POST", url, nil)
	if err != nil {
		ErrorConsole.Println(err)
	}
	return rm.send(req)
}

func (rm *RestManager) Update(request *entities.Request, index int) (int, error) {
	url := fmt.Sprintf(
		"%s/%d/issues/%d?%s",
		baseUrl,
		rm.projectId,
		request.Issues[index].Id,
		request.RequestParam(index),
	)
	req, err := http.NewRequest("PUT", url, nil)
	if err != nil {
		ErrorConsole.Println(err)
	}
	return rm.send(req)
}

func (rm *RestManager) send(req *http.Request) (int, error) {
	req.Header.Add("PRIVATE-TOKEN", rm.accessToken)
	resp, err := rm.client.Do(req)
	if err != nil {
		return 0, err
	}
	if resp.StatusCode > 300 {
		return 0, fmt.Errorf("ошибка при выполнении запроса %s, код %d", req.URL, resp.StatusCode)
	}

	return rm.responseCreateHandler(resp)
}

func (rm *RestManager) responseCreateHandler(resp *http.Response) (int, error) {
	re := entities.Response{}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&re); err != nil {
		return 0, err
	}

	return re.Iid, nil
}

func (rm *RestManager) responseMilestoneHandler(resp *http.Response) (*[]entities.Milestone, error) {
	var m []entities.Milestone
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&m); err != nil {
		return nil, err
	}
	return &m, nil
}

func (rm *RestManager) responseUploadsHandler(resp *http.Response) (*entities.Uploads, error) {
	u := entities.Uploads{}
	defer resp.Body.Close()
	if err := json.NewDecoder(resp.Body).Decode(&u); err != nil {
		return nil, err
	}

	return &u, nil
}
