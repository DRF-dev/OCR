package factories

import (
	"OCR/infra/models/code"
	"OCR/infra/models/content"
	"OCR/infra/models/file_content"
)

func (Model) GetFileContentEntity(content []byte) file_content.Model {
	return file_content.Model{Content: content}
}

func (Model) GetContentEntity(contentParam string) content.Model {
	return content.Model{Content: contentParam}
}

func (Model) GetCodeEntity(value string) code.Model  {
	return code.Model{Value: value}
}
