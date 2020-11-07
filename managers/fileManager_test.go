package managers

import (
	"github.com/spf13/afero"
	"testing"
)

func TestFileManager_FileRead(t *testing.T) {
	fm := NewFileManager(afero.NewMemMapFs())
	err := afero.WriteFile(fm.manager, "issues.yaml", []byte(getDataYaml()), 0666)
	if err != nil {
		t.Errorf("TestFileManager_FileRead failed, got error: %v", err)
	}

	req, err := fm.FileRead()
	if err != nil {
		t.Errorf("TestFileManager_FileRead failed, got error: %v", err)
	}
	i1 := req.Issues[0]
	i2 := req.Issues[1]
	if req.AccessToken != "abc123" || req.ProjectId != 321123 || req.Milestone != "my-milestone" ||
		i1.Title != "title1" || i1.Description != "desc1" || i1.Labels[0] != "new label" || i1.Labels[1] != "lab 7" ||
		i1.Milestone != "my-milestone" || i1.Weight != 1 || i2.Title != "title2" || i2.Description != "desc2" ||
		i2.Labels[0] != "lab5" || i2.Milestone != "my-milestone" || i2.Weight != 1 {
		t.Errorf("TestFileManager_FileRead failed, got %v", req)
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
