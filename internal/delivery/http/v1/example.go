package v1

import (
	"app/internal/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) initExampleRoutes(api *gin.RouterGroup) {
	reports := api.Group("/example")
	{
		reports.POST("/create", h.exampleCreate)
		reports.GET("/getAll", h.exampleGetAll)
	}
}

type ExampleCreateInput struct {
	Name string
}

type ExampleCreateOutput struct {
	Example model.Example `json:"example"`
}

// exampleCreate godoc
// @Tags example
// @Produce json
// @Param params body ExampleCreateInput true " "
// @Success 200 {object} ExampleCreateOutput
// @Failure 500 {object} Errors
// @Router /api/v1/example/create [post]
func (h *Handler) exampleCreate(ctx *gin.Context) {
	var input ExampleCreateInput

	err := ctx.ShouldBindJSON(&input)
	if err != nil {
		h.logger.Error(err)
		ValidationError(ctx, err)
		return
	}

	createdExample, err := h.exampleService.Create(ctx.Request.Context(), input.Name)
	if err != nil {
		h.logger.Error(err)
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ExampleCreateOutput{
		Example: createdExample,
	})
}

type ExampleGetAllOutput struct {
	Examples []model.Example `json:"examples"`
}

// exampleGetAll godoc
// @Tags example
// @Produce json
// @Success 200 {object} ExampleGetAllOutput
// @Failure 500 {object} Errors
// @Router /api/v1/example/getAll [get]
func (h *Handler) exampleGetAll(ctx *gin.Context) {
	examples, err := h.exampleService.GetAll(ctx.Request.Context())
	if err != nil {
		h.logger.Error(err)
		InternalServerError(ctx, err)
		return
	}

	ctx.JSON(http.StatusOK, ExampleGetAllOutput{
		Examples: examples,
	})
}