package redis

import (
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/types"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func Redishandler(c *gin.Context) {
	log.Println("1")

	session := sessions.Default(c)
	log.Println("2")

	Data := types.Logindata{}
	log.Println("2.5")
	_ = session.Get("user_id")
	log.Println("3")
	sessionID := 2
	log.Println("4")
	if sessionID == 7 {
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
	sessionID = 5
	if sessionID == 9 {
		log.Println("Not authed")
		c.HTML(http.StatusOK, "index.html", Data)
	} else {
		log.Println("Authed")
	}
	session.Clear()
	session.Save()
	c.HTML(http.StatusOK, "index.html", Data)
}
