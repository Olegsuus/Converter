package adapters

import (
	"converter/internal/services/domain"
	"fmt"
	"os/exec"
	"strings"
)

type AppleScriptAdapter struct{}

func NewAppleScriptAdapter() *AppleScriptAdapter {
	return &AppleScriptAdapter{}
}

func (a *AppleScriptAdapter) ConvertFile(
	inputPath, outputPath string, inputFmt, outputFmt domain.FileFormat,
) error {

	appName, asExportType, err := a.mapFormats(inputFmt, outputFmt)
	if err != nil {
		return err
	}

	appleScriptCmd := fmt.Sprintf(`
    tell application "%s"
        set doc to open POSIX file "%s"
        export doc to POSIX file "%s" as %s
        close doc
    end tell
	`, appName, escapePath(inputPath), escapePath(outputPath), asExportType)

	cmd := exec.Command("osascript", "-e", appleScriptCmd)
	out, err2 := cmd.CombinedOutput()
	if err2 != nil {
		return fmt.Errorf("AppleScript error: %v, output: %s", err2, string(out))
	}

	return nil
}

func (a *AppleScriptAdapter) mapFormats(
	inputFmt, outputFmt domain.FileFormat,
) (string, string, error) {

	var appName string
	var asExportType string

	switch inputFmt {
	case domain.FileFormatPages:
		appName = "Pages"
	case domain.FileFormatNumbers:
		appName = "Numbers"
	case domain.FileFormatKey:
		appName = "Keynote"
	default:
		return "", "", fmt.Errorf("unsupported input format for AppleScript: %s", inputFmt)
	}

	switch appName {
	case "Pages":
		switch outputFmt {
		case domain.FileFormatPDF:
			asExportType = "PDF"
		case domain.FileFormatDocx, domain.FileFormatDoc:
			asExportType = "Microsoft Word"
		default:
			return "", "", fmt.Errorf("unsupported output format for Pages: %s", outputFmt)
		}

	case "Numbers":
		switch outputFmt {
		case domain.FileFormatPDF:
			asExportType = "PDF"
		case domain.FileFormatXlsx:
			asExportType = "Excel"
		case domain.FileFormatDocx, domain.FileFormatDoc:
			return "", "", fmt.Errorf("Numbers can export to PDF or Excel, but not DOCX")
		default:
			return "", "", fmt.Errorf("unsupported output format for Numbers: %s", outputFmt)
		}

	case "Keynote":
		switch outputFmt {
		case domain.FileFormatPDF:
			asExportType = "PDF"
		case domain.FileFormatPptx:
			asExportType = "PowerPoint"
		case domain.FileFormatDocx, domain.FileFormatDoc:
			return "", "", fmt.Errorf("Keynote can export to PDF or PowerPoint, but not DOCX")
		default:
			return "", "", fmt.Errorf("unsupported output format for Keynote: %s", outputFmt)
		}

	}

	return appName, asExportType, nil
}

func escapePath(path string) string {
	return strings.ReplaceAll(path, `"`, `\"`)
}
