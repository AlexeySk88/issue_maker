package managers

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"issue_maker/entities"
	"os"
	"time"
)

const fileReadName = "issues"
const fileWriteName = "done"
const fileExtension = ".yaml"
const fileLog = "issue_maker.log"

func FileRead() (*entities.Request, error) {
	fileName := fileReadName + fileExtension
	if _, err := os.Stat(fileName); err != nil {
		return nil, fmt.Errorf("файла %s в директории с проектом не найдено", fileName)
	}

	data, err := ioutil.ReadFile(fileName)
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

func FileWrite(request *entities.Request) error {
	fileName := fileWriteName + "_" + time.Now().Format("02-01-2006_15-04-05") + fileExtension
	file, err := os.OpenFile(fileName, os.O_CREATE, 0777)
	if err != nil {
		return err
	}
	data, err := yaml.Marshal(request)
	if err != nil {
		return err
	}

	if _, err = file.Write(data); err != nil {
		return err
	}

	fmt.Printf("Файл с именем %s создан\n", fileName)
	return nil
}

func GetFileLog() (*os.File, error) {
	return os.OpenFile(fileLog, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
}
