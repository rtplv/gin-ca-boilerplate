package v1

import (
	"app/internal/service"
	"app/pkg/logs"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	exampleService service.Example
	logger *logs.Logger
}

func NewHandler(exampleService service.Example, logger *logs.Logger) *Handler {
	return &Handler{
		exampleService: exampleService,
		logger: logger,
	}
}

func (h Handler) Init(api *gin.RouterGroup) {
	v1 := api.Group("/v1")
	{
		h.initExampleRoutes(v1)
		h.initSwaggerRoutes(v1)
	}
}