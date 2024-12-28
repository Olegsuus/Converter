package handlers

import (
	"converter/internal/services/domain"
	"log/slog"
)

type ConverterHandlers struct {
	csP ConvertServicesProvider
	l   *slog.Logger
}

type ConvertServicesProvider interface {
	Convert(req domain.ConversionRequest) (domain.ConversionResult, error)
}

func RegisterConverterHandlers(csP ConvertServicesProvider, l *slog.Logger) *ConverterHandlers {
	return &ConverterHandlers{
		csP: csP,
		l:   l,
	}
}
