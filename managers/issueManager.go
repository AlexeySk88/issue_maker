package managers

import (
	"fmt"
	"issue_maker/entities"
	"issue_maker/helpers"
	"strings"
)

type IssueManager struct {
	rm *RestManager
	fm *FileManager
	r  *entities.Request
}

func NewIssueManager(rm *RestManager, fm *FileManager, r *entities.Request) *IssueManager {
	return &IssueManager{rm: rm, fm: fm, r: r}
}

func (im *IssueManager) Send() (*entities.Request, error) {
	newReq := entities.Request{
		AccessToken: im.r.AccessToken,
		ProjectId:   im.r.ProjectId,
		Milestone:   im.r.Milestone,
	}

	for index, issue := range im.r.Issues {
		desc, err := im.descriptionUploadImage(issue.Description)
		if err != nil {
			return nil, err
		}
		issue.Description = desc
		im.r.Issues[index] = issue

		var iid int
		if issue.IsCreate() {
			fmt.Printf("Записываю в gitlab задачу с заголовком: %s\n", issue.Title)
			iid, err = im.rm.Create(im.r.RequestParam(index))
		} else if issue.IsUpdate() {
			fmt.Printf("Записываю в gitlab задачу с заголовком: %s\n", issue.Title)
			iid, err = im.rm.Update(im.r.Issues[index].Id, im.r.RequestParam(index))
		} else {
			ErrorConsole.Printf("Задача с заголовком %s не будет отправлена в gitlab\n", issue.Title)
			continue
		}
		if err != nil {
			return nil, err
		}

		newReq.Issues = append(newReq.Issues, *issue.CopyForWrite(iid))
	}

	fmt.Println("Задачи успешно записаны!")
	return &newReq, nil
}

func (im *IssueManager) descriptionUploadImage(desc string) (string, error) {
	links := helpers.FindAllImageLinks(desc)
	if !im.fm.CheckExistFilesInBasePath(links) {
		return "", fmt.Errorf("найдены не все изображения")
	}

	newDesc := desc
	for _, imageName := range links {
		imagePath := im.fm.basePath + imageName
		markdown, err := im.rm.UploadFile(imagePath)
		if err != nil {
			return "", err
		}
		newDesc = strings.ReplaceAll(newDesc, imagePath, markdown)
	}

	return newDesc, nil
}
