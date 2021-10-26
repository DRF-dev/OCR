package main

import (
	"OCR/app"
	"OCR/infra/factories"
	"OCR/infra/services"
)

func main()  {
	var fileService services.FileServiceInterface = services.FileService{}
	var factories factories.Interface = factories.Model{}

	var getFileContent app.GetFileContentInterface = app.GetFileContentModel{
		FileService: fileService,
		Factories: factories,
	}
	var fromFileContentToCode app.FromFileContentToCodeInterface = app.FromFileContentToCodeModel{
		Factories: factories,
	}
	var writeResultToFile app.WriteFileContentInterface = app.WriteFileContentModel{
		FileService: fileService, Factories: factories,
	}

	content := getFileContent.Execute()
	codes := fromFileContentToCode.Execute(content)
	writeResultToFile.Execute(codes)
}
