package main

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
	"time"
)

const fileReadName = "issues"
const fileWriteName = "done"
const fileExtension = ".yaml"

func fileRead() (*Request, error) {
	fileName := fileReadName + fileExtension
	if _, err := os.Stat(fileName); err != nil {
		return nil, fmt.Errorf("файла %s в директории с проектом не найдено", fileName)
	}

	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}

	r := Request{}
	err = yaml.Unmarshal(data, &r)
	if err != nil {
		return nil, err
	}

	err = r.validation()
	if err != nil {
		return nil, err
	}

	return &r, nil
}

func fileWrite(request *Request) error {
	fileName := fileWriteName + "_" + time.Now().Format("02-01-2006") + fileExtension
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

	return nil
}
