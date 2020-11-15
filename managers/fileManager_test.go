package managers

import (
	"github.com/spf13/afero"
	"github.com/stretchr/testify/require"
	"strings"
	"testing"
)

var fm = NewFileManager(afero.NewMemMapFs())

func TestFileManager_ReadIssuesFile(t *testing.T) {
	_, err := fm.ReadIssuesFile()
	require.Error(t, err)

	err = afero.WriteFile(fm.manager, "issues.yaml", []byte(getDataYaml()), 0666)
	require.NoError(t, err)

	req, err := fm.ReadIssuesFile()
	require.NoError(t, err)

	i1 := req.Issues[0]
	i2 := req.Issues[1]
	require.True(t, req.AccessToken == "abc123")
	require.True(t, req.ProjectId == 321123)
	require.True(t, req.Milestone == "my-milestone")
	require.True(t, i1.Title == "title1")
	require.True(t, i1.Description == "desc1")
	require.True(t, i1.Labels[0] == "new label")
	require.True(t, i1.Labels[1] == "lab 7")
	require.True(t, i1.Milestone == "my-milestone")
	require.True(t, i1.Weight == 1)
	require.True(t, i2.Title == "title2")
	require.True(t, i2.Description == "desc2")
	require.True(t, i2.Labels[0] == "lab5")
	require.True(t, i2.Milestone == "my-milestone")
	require.True(t, i2.Weight == 1)

	_ = fm.manager.Remove("issues.yaml")
}

func TestFileManager_ReadIssuesFileFromPath(t *testing.T) {
	_, err := fm.ReadIssuesFileFromPath(".")
	require.Error(t, err)

	_ = afero.WriteFile(fm.manager, "issues.yaml", []byte("test message"), 0666)
	_, err = fm.ReadIssuesFileFromPath(".")
	require.Error(t, err)

	noValidIssues := strings.Replace(getDataYaml(), "project_id: 321123", "", 1)
	_ = afero.WriteFile(fm.manager, "issues.yaml", []byte(noValidIssues), 0666)
	_, err = fm.ReadIssuesFileFromPath(".")
	require.Error(t, err)

	_ = fm.manager.Remove("issues.yaml")
}

func TestFileManager_CheckExistFiles(t *testing.T) {
	arr := []string{"file1.txt", "file2.txt"}
	for _, path := range arr {
		_, err := fm.manager.Create(path)
		require.NoError(t, err)
	}

	require.True(t, fm.CheckExistFilesInBasePath(arr))
	arr = append(arr, "file3.txt")
	require.False(t, fm.CheckExistFilesInBasePath(arr))
}

func TestFileManager_GetFileLog(t *testing.T) {
	logFile, err := fm.GetFileLog()
	require.NoError(t, err)
	require.True(t, fm.checkExistFile(logFile.Name()))
}

func TestFileManager_GetFile(t *testing.T) {
	fileName := "testFile.txt"
	_, err := fm.manager.Create(fileName)
	require.NoError(t, err)

	file, err := fm.GetFile(fileName)
	require.NoError(t, err)
	require.True(t, file.Name() == fileName)
}

func TestFileManager_WriteDoneFile(t *testing.T) {
	doneIsExist := checkDoneFile()
	require.False(t, doneIsExist)

	err := fm.WriteDoneFile(getRequest())
	require.NoError(t, err)

	doneIsExist = checkDoneFile()
	require.True(t, doneIsExist)
}

func checkDoneFile() bool {
	files, _ := afero.ReadDir(fm.manager, "./")
	for _, file := range files {
		if strings.Contains(file.Name(), "done") {
			return true
		}
	}
	return false
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
