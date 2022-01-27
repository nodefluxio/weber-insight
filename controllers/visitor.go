package controllers

import (
	"fmt"
	"net/http"
	"os"
	"weber-insight/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
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

func (ctrl *Controller) ExportVisitors(ctx *gin.Context) {
	var visitors []models.Visitor
	err := ctrl.Model.GetVisitors(&visitors)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	csvContent, err := gocsv.MarshalString(visitors)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment;filename=user-list.csv")
	ctx.Data(http.StatusOK, "text/csv", []byte(fmt.Sprintf("%v", csvContent)))

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
