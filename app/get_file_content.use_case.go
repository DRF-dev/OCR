package app

import (
	"OCR/infra/factories"
	"OCR/infra/models/file_content"
	"OCR/infra/services"
)

type GetFileContentModel struct {
	FileService services.FileServiceInterface
	Factories factories.Interface
}

type GetFileContentInterface interface {
	Execute() file_content.Model
}

func (c GetFileContentModel) Execute() file_content.Model  {
	bytes := c.FileService.GetFileContent()
	fileContentEntity := c.Factories.GetFileContentEntity(bytes)
	return fileContentEntity
}
