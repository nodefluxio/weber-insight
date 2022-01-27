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

func (ctrl *Controller) GetVisitorActivities(ctx *gin.Context) {
	var visitorActivities []models.VisitorActivities
	err := ctrl.Model.GetVisitorActivities(&visitorActivities)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")

	ctx.HTML(http.StatusOK, "user-activities.tmpl", gin.H{
		"name":           session.Get("name"),
		"base_url":       baseUrl,
		"title":          "Weber Insight - User Activities",
		"useractivities": true,
		"data":           visitorActivities,
	})
}

func (ctrl *Controller) ExportUserActivities(ctx *gin.Context) {
	var visitorActivities []models.VisitorActivities
	err := ctrl.Model.GetVisitorActivities(&visitorActivities)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}
	csvContent, err := gocsv.MarshalString(visitorActivities)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment;filename=user-activities.csv")
	ctx.Data(http.StatusOK, "text/csv", []byte(fmt.Sprintf("%v", csvContent)))
}
