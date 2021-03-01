package templates

import (
	"github.com/gin-gonic/gin"
)

func InitTemplates(router *gin.Engine) {
	router.SetFuncMap(GetMainTemplatesFuncMap())

	// load templates
	router.LoadHTMLGlob("templates/*")

	// add static files route
	router.Static("/static", "static/")
}
