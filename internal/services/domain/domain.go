package domain

type FileFormat string

const (
	FileFormatUnknown FileFormat = "unknown"
	FileFormatPages   FileFormat = "pages"
	FileFormatNumbers FileFormat = "numbers"
	FileFormatKey     FileFormat = "key"
	FileFormatXlsx    FileFormat = "xlsx"
	FileFormatPptx    FileFormat = "pptx"

	FileFormatDoc  FileFormat = "doc"
	FileFormatDocx FileFormat = "docx"
	FileFormatPDF  FileFormat = "pdf"
)

type ConversionRequest struct {
	InputPath  string
	OutputPath string
	InputFmt   FileFormat
	OutputFmt  FileFormat
}

type ConversionResult struct {
	Success bool
	Error   error
}
