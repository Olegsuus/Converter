package router

import (
	handlers "converter/internal/controllers/rest/handlers/converter"
	"github.com/gin-gonic/gin"
)

func RegisterRoutes(r *gin.Engine, convHandlers *handlers.ConverterHandlers) {

	r.POST("/convert", convHandlers.ConvertFile)

}
