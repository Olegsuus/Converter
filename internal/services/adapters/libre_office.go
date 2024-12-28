package adapters

import (
	"converter/internal/services/domain"
	"errors"
	"fmt"
	"os/exec"
	"path/filepath"
)

type LibreOfficeAdapter struct{}

func NewLibreOfficeAdapter() *LibreOfficeAdapter {
	return &LibreOfficeAdapter{}
}

func (l *LibreOfficeAdapter) ConvertFile(
	inputPath string,
	outputPath string,
	inputFmt domain.FileFormat,
	outputFmt domain.FileFormat,
) error {
	outDir := filepath.Dir(outputPath)

	convertParam, err := mapLibreOfficeParam(outputFmt)
	if err != nil {
		return err
	}

	cmd := exec.Command("libreoffice",
		"--headless",
		"--convert-to", convertParam,
		"--outdir", outDir,
		inputPath,
	)

	output, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("libreoffice error: %v, output: %s", err2, string(output))
	}

	return nil
}

func mapLibreOfficeParam(fmt domain.FileFormat) (string, error) {
	switch fmt {
	case domain.FileFormatPDF:
		return "pdf", nil
	case domain.FileFormatDocx:
		return "docx", nil
	case domain.FileFormatDoc:
		return "doc", nil
	default:
		return "", errors.New("unsupported output format for LibreOffice")
	}
}
