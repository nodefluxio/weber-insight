package controllers

import (
	"fmt"
	"net/http"
	"os"
	"time"
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

func (ctrl *Controller) CreateServiceView(ctx *gin.Context) {
	session := sessions.Default(ctx)
	baseUrl := os.Getenv("BASE_URL")

	ctx.HTML(http.StatusOK, "create-service.tmpl", gin.H{
		"name":           session.Get("name"),
		"base_url":       baseUrl,
		"title":          "Weber Insight - Create Service",
		"manageservices": true,
	})
	return
}

func (ctrl *Controller) CreateService(ctx *gin.Context) {
	service := models.Service{
		Name:               ctx.PostForm("name"),
		Type:               ctx.PostForm("type"),
		Slug:               ctx.PostForm("slug"),
		Thumbnail:          ctx.PostForm("thumbnail"),
		AccessKey:          ctx.PostForm("access_key"),
		Token:              ctx.PostForm("token"),
		Timestamp:          ctx.PostForm("timestamp"),
		ShortDescription:   ctx.PostForm("short_description"),
		LongDescription:    ctx.PostForm("long_description"),
		SpecialInstruction: ctx.PostForm("special_instruction"),
		CreatedAt:          time.Now(),
		UpdatedAt:          time.Now(),
	}

	ctrl.Model.CreateService(&service)

	ctrl.GetServices(ctx)
	return
}
