package types

import (
	"log"
	"net/http"

	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

type Book struct {
	Author string
	Title  string
}

type Logindata struct {
	Username   string
	Isloggedin bool
}
type UserController struct{}

func (ctr UserController) Redishandler(c *gin.Context) {
	log.Println("1")

	session := sessions.Default(c)
	log.Println("2")
	Data := Logindata{}
	sessionID := session.Get("user_id")
	log.Println("3")

	if sessionID == nil {
		log.Println("Not authed")
		c.HTML(http.StatusOK, "index.html", Data)
	} else {
		log.Println("Authed")
	}
	session.Set("user_id", 1)
	session.Set("user_email", "demo@demo.com")
	session.Set("user_username", "demo")
	session.Save()

	Data.Isloggedin = true
	Data.Username = "redistest"
	sessionID = session.Get("user_id")
	if sessionID == nil {
		log.Println("Not authed")
		c.HTML(http.StatusOK, "index.html", Data)
	} else {
		log.Println("Authed")
	}
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index.html", Data)
}
