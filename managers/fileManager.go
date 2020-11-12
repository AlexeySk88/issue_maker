package managers

import (
	"fmt"
	"github.com/spf13/afero"
	"gopkg.in/yaml.v2"
	"issue_maker/entities"
	"os"
	"path/filepath"
	"time"
)

const fileReadName = "issues"
const fileWriteName = "done"
const fileExtension = ".yaml"
const fileLog = "issue_maker.log"

type FileManager struct {
	manager afero.Fs
}

func NewFileManager(fs afero.Fs) *FileManager {
	return &FileManager{manager: fs}
}

func (fm *FileManager) ReadIssuesFile() (*entities.Request, error) {
	filePath := fileReadName + fileExtension
	if !fm.checkExistFile(filePath) {
		ex, err := os.Executable()
		if err != nil {
			panic(err)
		}
		exPath := filepath.Dir(ex)
		return nil, fmt.Errorf("файла %s в директории %s не найдено", filePath, exPath)
	}
	return fm.ReadIssuesFileFromPath(filePath)
}

func (fm *FileManager) ReadIssuesFileFromPath(filePath string) (*entities.Request, error) {
	if !fm.checkExistFile(filePath) {
		return nil, fmt.Errorf("файла %s не найдено", filePath)
	}

	data, err := afero.ReadFile(fm.manager, filePath)
	if err != nil {
		return nil, err
	}

	r := entities.Request{}
	err = yaml.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	err = r.Validation()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func (fm *FileManager) WriteDoneFile(request *entities.Request) error {
	fileName := fileWriteName + "_" + time.Now().Format("02-01-2006_15-04-05") + fileExtension
	data, err := yaml.Marshal(request)
	if err != nil {
		return err
	}

	if err = afero.WriteFile(fm.manager, fileName, data, 0777); err != nil {
		return err
	}

	fmt.Printf("Файл с именем %s создан\n", fileName)
	return nil
}

func (fm *FileManager) GetFileLog() (afero.File, error) {
	return fm.manager.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}

func (fm *FileManager) GetFile(str string) (*os.File, error) {
	return os.Open(str)
}

func (fm *FileManager) CheckExistFiles(arr []string) bool {
	res := true
	for _, s := range arr {
		if !fm.checkExistFile(s) {
			ErrorConsole.Println("Файл %s не найден", s)
			res = false
		}
	}
	return res
}

func (fm *FileManager) checkExistFile(path string) bool {
	_, err := fm.manager.Stat(path)
	return err == nil
}
