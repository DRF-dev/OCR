package app

import (
	"OCR/infra/factories"
	"OCR/infra/models/code"
	"OCR/infra/models/file_content"
)

type FromFileContentToCodeModel struct {
	Factories factories.Interface
}

type FromFileContentToCodeInterface interface {
	Execute(model file_content.Model) []code.Model
}

func (f FromFileContentToCodeModel) Execute(model file_content.Model) []code.Model {
	strContent := model.FromBytesToString()
	contentEntity := f.Factories.GetContentEntity(strContent)
	codes := contentEntity.GetCaracteresFromEachLinesColumns()

	codeEntities := make([]code.Model, len(codes))
	for i, c := range codes {
		codeEntities[i] = f.Factories.GetCodeEntity(c)
	}

	return codeEntities
}