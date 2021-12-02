package controllers

import (
	"fmt"
	"net/http"
	"os"
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
	baseUrl := os.Getenv("BASE_URL")
	ctx.HTML(http.StatusOK, "services.tmpl", gin.H{
		"name":           session.Get("name"),
		"base_url":       baseUrl,
		"title":          "Weber Insight - Manage Services",
		"manageservices": true,
		"data":           services,
	})

}

func (ctrl *Controller) DeleteService(ctx *gin.Context) {
	var service models.Service
	id := ctx.Param("id")
	fmt.Println(id)
	ctrl.Model.DeleteService(&service, id)

	ctx.JSON(http.StatusOK, gin.H{
		"ok":      true,
		"message": "Service has been deleted successfully",
	})
	return
}
