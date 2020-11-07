package managers

import (
	"fmt"
	"issue_maker/helpers"
	"strings"
)

func descriptionUploadImage(desc string, projectId int, accessToken string) (string, error) {
	links := helpers.FindAllImageLinks(desc)
	if !CheckExistFiles(links) {
		return "", fmt.Errorf("найдены не все изображения")
	}

	newDesc := desc
	for _, imagePath := range links {
		markdown, err := uploadFile(imagePath, projectId, accessToken)
		if err != nil {
			return "", err
		}
		newDesc = strings.ReplaceAll(newDesc, imagePath, markdown)
	}

	return newDesc, nil
}
