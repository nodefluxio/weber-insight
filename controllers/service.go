package controllers

import (
	"net/http"
	"weber-insight/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func (ctrl *Controller) GetServices(ctx *gin.Context) {
	var services []models.Service
	err := ctrl.Model.GetServices(&services)
	if err != nil {
		ctx.Redirect(http.StatusTemporaryRedirect, "/error")
		return
	}

	session := sessions.Default(ctx)
	ctx.HTML(http.StatusOK, "services.tmpl", gin.H{
		"name":            session.Get("name"),
		"title":           "Weber Insight - Manage Services",
		"managerservices": true,
		"data":            services,
	})

}
