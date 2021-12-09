package controllers

import (
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetFeedback(ctx *gin.Context) {
	feedback, err := ctrl.Model.GetAllFeedbackToView()
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")
	ctx.HTML(http.StatusOK, "feedback.tmpl", gin.H{
		"name":         session.Get("name"),
		"base_url":     baseUrl,
		"title":        "Weber Insight - Feedback",
		"userfeedback": true,
		"data":         feedback,
	})

}
