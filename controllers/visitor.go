package controllers

import (
	"net/http"
	"os"
	"weber-insight/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetVisitors(ctx *gin.Context) {
	var visitors []models.Visitor
	err := ctrl.Model.GetVisitors(&visitors)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")
	ctx.HTML(http.StatusOK, "user-list.tmpl", gin.H{
		"name":     session.Get("name"),
		"base_url": baseUrl,
		"title":    "Weber Insight - User List",
		"userlist": true,
		"data":     visitors,
	})

}

func (ctrl *Controller) GetAMLPEPVisitors(ctx *gin.Context) {
	visitors, err := ctrl.Model.GetAMLPEPVisitors()
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")
	ctx.HTML(http.StatusOK, "aml-pep-user-list.tmpl", gin.H{
		"name":           session.Get("name"),
		"base_url":       baseUrl,
		"title":          "Weber Insight - AML / PEP User List",
		"amlpepuserlist": true,
		"data":           visitors,
	})

}
