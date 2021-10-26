package factories

import (
	"OCR/infra/models/code"
	"OCR/infra/models/content"
	"OCR/infra/models/file_content"
)

type Interface interface {
	GetFileContentEntity(content []byte) file_content.Model
	GetContentEntity(content string) content.Model
	GetCodeEntity(value string) code.Model
}
