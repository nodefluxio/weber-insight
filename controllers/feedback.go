package controllers

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/gocarina/gocsv"
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

func (ctrl *Controller) ExportFeedback(ctx *gin.Context) {
	feedback, err := ctrl.Model.GetAllFeedbackToView()
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	csvContent, err := gocsv.MarshalString(feedback)
	if err != nil {
		fmt.Println(err)
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	ctx.Header("Content-Type", "text/csv")
	ctx.Header("Content-Disposition", "attachment;filename=user-feedback.csv")
	ctx.Data(http.StatusOK, "text/csv", []byte(fmt.Sprintf("%v", csvContent)))

}
