package managers

import (
	"github.com/spf13/afero"
	"testing"
)

var fm = NewFileManager(afero.NewMemMapFs())

func TestFileManager_ReadIssuesFile(t *testing.T) {
	err := afero.WriteFile(fm.manager, "issues.yaml", []byte(getDataYaml()), 0666)
	if err != nil {
		t.Errorf("TestFileManager_ReadIssuesFile failed, got error: %v", err)
	}

	req, err := fm.ReadIssuesFile()
	if err != nil {
		t.Errorf("TestFileManager_ReadIssuesFile failed, got error: %v", err)
	}
	i1 := req.Issues[0]
	i2 := req.Issues[1]
	if req.AccessToken != "abc123" || req.ProjectId != 321123 || req.Milestone != "my-milestone" ||
		i1.Title != "title1" || i1.Description != "desc1" || i1.Labels[0] != "new label" || i1.Labels[1] != "lab 7" ||
		i1.Milestone != "my-milestone" || i1.Weight != 1 || i2.Title != "title2" || i2.Description != "desc2" ||
		i2.Labels[0] != "lab5" || i2.Milestone != "my-milestone" || i2.Weight != 1 {
		t.Errorf("TestFileManager_ReadIssuesFile failed, got %v", req)
	}
}

func TestFileManager_CheckExistFiles(t *testing.T) {
	arr := []string{"file1.txt", "file2.txt"}
	for _, path := range arr {
		_, err := fm.manager.Create(path)
		if err != nil {
			t.Errorf("TestFileManager_CheckExistFile failed, got error: %v", err)
		}
	}

	if !fm.CheckExistFiles(arr) {
		t.Errorf("TestFileManager_CheckExistFile failed, expected true, got false")
	}

	arr = append(arr, "file3.txt")
	if fm.CheckExistFiles(arr) {
		t.Errorf("TestFileManager_CheckExistFile failed, expected false, got true")
	}
}

func getDataYaml() string {
	return `
token: abc123
project_id: 321123
milestone: my-milestone
issues:
  - title: title1
    description: desc1
    labels:
      - new label
      - lab 7
    milestone: my-milestone
    weight: 1
  - title: title2
    description: desc2
    labels:
      - lab5
    milestone: my-milestone
    weight: 1
`
}
