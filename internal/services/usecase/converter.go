package usecase

import (
	"converter/internal/services/domain"
)

type ConverterUseCase interface {
	Convert(req domain.ConversionRequest) (domain.ConversionResult, error)
}

type ConvertAdapter interface {
	ConvertFile(inputPath, outputPath string, inputFmt, outputFmt domain.FileFormat) error
}

type CompositeConverter struct {
	apple ConvertAdapter
	libre ConvertAdapter
}

func NewCompositeConverterUseCase(apple, libre ConvertAdapter) ConverterUseCase {
	return &CompositeConverter{
		apple: apple,
		libre: libre,
	}
}

func (c *CompositeConverter) Convert(req domain.ConversionRequest) (domain.ConversionResult, error) {

	switch req.InputFmt {
	case domain.FileFormatPages, domain.FileFormatNumbers, domain.FileFormatKey:
		err := c.apple.ConvertFile(
			req.InputPath,
			req.OutputPath,
			req.InputFmt,
			req.OutputFmt,
		)
		if err != nil {
			return domain.ConversionResult{Success: false, Error: err}, err
		}
		return domain.ConversionResult{Success: true}, nil

	default:
		err := c.libre.ConvertFile(
			req.InputPath,
			req.OutputPath,
			req.InputFmt,
			req.OutputFmt,
		)
		if err != nil {
			return domain.ConversionResult{Success: false, Error: err}, err
		}
		return domain.ConversionResult{Success: true}, nil
	}
}
