package http

import (
	"app/internal/delivery/http/templates"
	"app/internal/delivery/http/v1"
	"app/internal/service"
	"app/pkg/logs"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	exampleService service.Example
	logger logs.Logger
}

func NewHandler(exampleService service.Example, logger logs.Logger) *Handler {
	return &Handler{
		exampleService: exampleService,
		logger: logger,
	}
}

func (h Handler) Init(mode string) *gin.Engine {
	// Set server mode
	gin.SetMode(mode)

	// Init routes
	router := gin.New()

	// Init templates
	templates.InitTemplates(router)

	h.initAPI(router)

	return router
}

func (h Handler) initAPI(router *gin.Engine) {
	handlerV1 := v1.NewHandler(h.exampleService, h.logger)

	api := router.Group("/api")
	{
		handlerV1.Init(api)
	}
}