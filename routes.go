package main

import (
	"net/http"
	"weber-insight/controllers"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func loginMiddleware(ctrl *controllers.Controller) gin.HandlerFunc {
	return func(c *gin.Context) {
		if c.Request.URL.String() != "/login" {
			if !ctrl.CheckLoggedIn(c) {
				c.Redirect(http.StatusTemporaryRedirect, "/login")
			}
		}
	}
}

func setupRouter(ctrl *controllers.Controller) *gin.Engine {
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	r.Static("/assets", "./views/assets")
	r.Static("/node_modules", "./views/node_modules")
	r.LoadHTMLGlob("views/pages/*")

	r.Use(loginMiddleware(ctrl))

	base := r.Group("")
	services := base.Group("/services")
	{
		services.GET("", ctrl.GetServices)
	}

	r.GET("/", ctrl.Index)

	r.GET("/error", func(c *gin.Context) {
		c.HTML(http.StatusOK, "error.tmpl", gin.H{
			"title": "Weber Insight - Error",
		})
	})

	r.GET("/login", func(c *gin.Context) {
		c.HTML(http.StatusOK, "login.html", gin.H{
			"title": "Weber Insight - Login",
		})
	})

	r.GET("/dashboard", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "dashboard.tmpl", gin.H{
			"name":      session.Get("name"),
			"title":     "Weber Insight - Dashboard",
			"dashboard": true,
		})
	})

	r.GET("/notification", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "notification.tmpl", gin.H{
			"name":              session.Get("name"),
			"title":             "Weber Insight - Notification",
			"emailnotification": true,
		})
	})

	r.GET("/export-data", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "export.tmpl", gin.H{
			"name":       session.Get("name"),
			"title":      "Weber Insight - Export Data",
			"exportdata": true,
		})
	})

	r.GET("/user-activities", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "user-activities.tmpl", gin.H{
			"name":       session.Get("name"),
			"title":      "Weber Insight - User Activities",
			"userlookup": true,
		})
	})

	r.GET("/user-feedback", func(c *gin.Context) {
		session := sessions.Default(c)
		c.HTML(http.StatusOK, "user-feedback.tmpl", gin.H{
			"name":         session.Get("name"),
			"title":        "Weber Insight - User Feedback",
			"userfeedback": true,
		})
	})

	r.POST("/login", ctrl.Login)

	r.GET("/logout", ctrl.Logout)

	return r
}
