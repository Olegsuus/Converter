package format

import (
	"converter/internal/services/domain"
	"path/filepath"
	"strings"
)

func GetFileFormat(path string) domain.FileFormat {
	ext := strings.ToLower(filepath.Ext(path))
	switch ext {
	case ".pages":
		return domain.FileFormatPages
	case ".numbers":
		return domain.FileFormatNumbers
	case ".key", ".keynote":
		return domain.FileFormatKey
	case ".doc":
		return domain.FileFormatDoc
	case ".docx":
		return domain.FileFormatDocx
	case ".pdf":
		return domain.FileFormatPDF
	default:
		return domain.FileFormatUnknown
	}
}

func ToFileFormat(fmtStr string) domain.FileFormat {
	fmtStr = strings.ToLower(fmtStr)
	switch fmtStr {
	case "pages":
		return domain.FileFormatPages
	case "numbers":
		return domain.FileFormatNumbers
	case "key", "keynote":
		return domain.FileFormatKey
	case "pdf":
		return domain.FileFormatPDF
	case "docx":
		return domain.FileFormatDocx
	case "doc":
		return domain.FileFormatDoc
	default:
		return domain.FileFormatUnknown
	}
}
