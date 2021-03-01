package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

type Errors struct {
	Errors []ErrorMessage `json:"errors"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func ValidationError(ctx *gin.Context, err error) {
	errorMessages := make([]ErrorMessage, 0)

	for _, msg := range strings.Split(err.Error(), "\n") {
		errorMessages = append(errorMessages, ErrorMessage{
			Message: msg,
		})
	}

	ctx.JSON(http.StatusBadRequest, Errors{
		Errors: errorMessages,
	})
}

func InternalServerError(ctx *gin.Context, err error) {
	ctx.JSON(http.StatusInternalServerError, Errors{
		Errors: []ErrorMessage{
			{Message: "Внутренняя ошибка сервера."},
			{Message: err.Error()},
		},
	})
}
