package app

import (
	"OCR/infra/factories"
	"OCR/infra/models/code"
	"OCR/infra/services"
)

type WriteFileContentModel struct {
	FileService services.FileServiceInterface
	Factories factories.Interface
}

type WriteFileContentInterface interface {
	Execute(codes []code.Model)
}

func (w WriteFileContentModel) Execute(codes []code.Model) {
	f, err := w.FileService.CreateFile("/Users/drf/Desktop/Essai_perso/Golang/OCR/result.txt")
	if err != nil {
		panic("cannot create the file successfully")
	}
	defer f.Close()

	err = w.FileService.AddResultsOnFile(f, codes)
	if err != nil {
		panic("cannot save the value on the file")
	}
}
