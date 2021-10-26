package services

import (
	"OCR/infra/models/code"
	"fmt"
	"io/ioutil"
	"os"
)

type FileService struct {}

type FileServiceInterface interface {
	GetFileContent() []byte
	CreateFile(path string) (*os.File, error)
	AddResultsOnFile(file *os.File, model []code.Model) error
}

func (FileService) GetFileContent() []byte {
	bytes, err := ioutil.ReadFile("/Users/drf/Desktop/Essai_perso/Golang/OCR/file.txt")
	if err != nil {
		panic(fmt.Sprintf("cannot get content of the file %s\n", err))
	}
	return bytes
}

func (FileService) CreateFile(path string) (*os.File,error) {
	file, err := os.Create(path)
	if err != nil {
		return nil, err
	}
	return file, nil
}

func (FileService) AddResultsOnFile(file *os.File, model []code.Model) error  {
	for _, m := range model {
		_, err := file.WriteString(m.Value + "\n")
		if err != nil {
			return err
		}
	}
	return nil
}