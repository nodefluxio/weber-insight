package controllers

import (
	"net/http"
	"os"
	"weber-insight/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetVisitorActivities(ctx *gin.Context) {
	var visitorActivities []models.VisitorActivity
	err := ctrl.Model.GetVisitorActivities(&visitorActivities)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")

	ctx.HTML(http.StatusOK, "user-activities.tmpl", gin.H{
		"name": session.Get("name"),
		"base_url": baseUrl,
		"title": "Weber Insight - User Activities",
		"useractivities": true,
		"data": visitorActivities,
	})
}