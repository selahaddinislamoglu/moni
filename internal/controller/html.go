package controller

import (
	"path/filepath"

	"github.com/gin-gonic/gin"
)

type HTMLController struct {
}

func NewHTMLController() HTML {
	return &HTMLController{}
}

func (h *HTMLController) Login(ctx *gin.Context) {
	ctx.File(filepath.Join("internal", "public", "pages", "login.html"))
}

func (h *HTMLController) Dashboard(ctx *gin.Context) {
	ctx.File(filepath.Join("internal", "public", "pages", "dashboard.html"))
}
