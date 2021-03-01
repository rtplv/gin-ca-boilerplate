package v1

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func (h *Handler) initSwaggerRoutes(api *gin.RouterGroup) {
	url := ginSwagger.URL("/api/v1/docs/swagger.json")

	api.StaticFile("/docs/swagger.json", "./docs/swagger.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))
}
