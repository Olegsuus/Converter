package handlers

import (
	"converter/internal/services/domain"
	"converter/internal/services/format"
	"fmt"
	"github.com/gin-gonic/gin"
	"log/slog"
	"net/http"
	"os"
	"path/filepath"
	"time"
)

func (h *ConverterHandlers) ConvertFile(c *gin.Context) {
	file, err := c.FormFile("file")
	if err != nil {
		c.String(http.StatusBadRequest, "file not found in form-data: %v", err)
		return
	}

	outputFormat := c.Query("to")
	if outputFormat == "" {
		outputFormat = "pdf"
	}
	h.l.With(slog.String("output format", outputFormat))

	inputPath := filepath.Join(os.TempDir(), file.Filename)
	if err := c.SaveUploadedFile(file, inputPath); err != nil {
		c.String(http.StatusInternalServerError, "failed to save file: %v", err)
		return
	}
	h.l.With(slog.String("input path", inputPath))

	inputFmt := format.GetFileFormat(inputPath)
	h.l.With(slog.String("input fmt", fmt.Sprintf(inputPath)))

	ext := extensionFor(outputFormat)
	resultName := fmt.Sprintf("%d_%s%s", time.Now().UnixNano(), file.Filename, ext)
	h.l.With(slog.String("result name", resultName))

	outputPath := filepath.Join(os.TempDir(), resultName)

	req := domain.ConversionRequest{
		InputPath:  inputPath,
		OutputPath: outputPath,
		InputFmt:   inputFmt,
		OutputFmt:  format.ToFileFormat(outputFormat),
	}
	h.l.With(slog.String("request", fmt.Sprintf("%v", req)))

	result, err := h.csP.Convert(req)
	if err != nil {
		c.String(http.StatusInternalServerError, "conversion failed: %v", err)
		return
	}
	h.l.With(slog.String("result", fmt.Sprintf("%v", result)))

	if !result.Success {
		c.String(http.StatusInternalServerError, "conversion not successful")
		return
	}

	c.FileAttachment(outputPath, filepath.Base(outputPath))
}
