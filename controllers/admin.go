package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-contrib/sessions"

	"github.com/gin-gonic/gin"
)

type GoogleToken struct {
	token string `json:"google-token" binding:"required"`
}

func (ctrl *Controller) Login(c *gin.Context) {
	googleToken := c.PostForm("google-token")
	resp, err := http.Get("https://oauth2.googleapis.com/tokeninfo?id_token=" + googleToken)
	if err != nil {
		fmt.Println("Get: " + err.Error() + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}
	defer resp.Body.Close()

	response, err2 := ioutil.ReadAll(resp.Body)
	if err2 != nil {
		fmt.Println("ReadAll: " + err.Error() + "\n")
		c.Redirect(http.StatusTemporaryRedirect, "/")
		return
	}

	var dataUser map[string]string
	json.Unmarshal(response, &dataUser)
	var email = dataUser["email"]
	var name = dataUser["name"]
	var picture = dataUser["picture"]
	var requiredEmail = "@nodeflux.io"

	if strings.Contains(email, requiredEmail) {
		session := sessions.Default(c)
		session.Set("logged_in", true)
		session.Set("name", name)
		session.Set("picture", picture)
		session.Save()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}
}

func (ctrl *Controller) Index(c *gin.Context) {
	if ctrl.CheckLoggedIn(c) {
		c.Redirect(http.StatusTemporaryRedirect, "/dashboard")
	} else {
		c.Redirect(http.StatusTemporaryRedirect, "/login")
	}
}

func (ctrl *Controller) CheckLoggedIn(c *gin.Context) bool {
	session := sessions.Default(c)
	if session.Get("logged_in") == true {
		return true
	}
	return false
}

func (ctrl *Controller) Logout(c *gin.Context) {
	session := sessions.Default(c)
	session.Clear()
	session.Options(sessions.Options{MaxAge: -1})
	session.Save()
	c.Redirect(http.StatusTemporaryRedirect, "/login")
}
