package login

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/database"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginHandler(c *gin.Context) {
	pwd := c.PostForm("pwd")
	hashpwd := sha256.Sum256([]byte(pwd))
	pwd = hex.EncodeToString(hashpwd[:])
	name := c.PostForm("name")
	pwdhashindb, err := database.SelectUsernameAndPwdhash(name)
	if err != nil {
		log.Println("Error happened in DB")
	}
	if pwd == pwdhashindb {
		AddUserToRedis(c, name)
		AddCookieToUser(c, name)
		datas := map[string]interface{}{
			"alert": false,
			"uname": name,
		}
		c.HTML(http.StatusOK, "index.html", datas)
	} else {
		uname, err := GetUsername(c)
		datas := map[string]interface{}{
			"alert": true,
			"uname": uname,
		}
		if err != nil {
			datas["uname"] = nil
		}
		c.HTML(http.StatusOK, "index.html", datas)
	}

}

func AddUserToRedis(c *gin.Context, username string) {
	session := sessions.Default(c)
	//set true for logged in
	session.Set("user_username", username)
	session.Save()
}

/*
func GetUserFromRedis(c *gin.Context) {
	session := sessions.Default(c)
	username := session.Get("user_username")
	c.HTML(http.StatusOK, "index.html", username)
}*/

func AddCookieToUser(c *gin.Context, name string) {
	c.SetCookie("username", name, 60*10, "/", "golangwebpagev2.herokuapp.com", true, true)
}

func HomepageHandler(c *gin.Context) {
	uname, err := GetUsername(c)
	datas := map[string]interface{}{
		"alert": false,
		"uname": uname,
	}
	if err != nil {
		datas["uname"] = nil
	}
	c.HTML(http.StatusOK, "index.html", datas)
}
func LogoutHandler(c *gin.Context) {
	uname, err := GetUsername(c)
	datas := map[string]interface{}{
		"alert": false,
		"uname": uname,
	}
	if err != nil {
		datas["uname"] = nil
		c.HTML(http.StatusOK, "index.html", datas)
		return
	}
	datas["uname"] = nil
	c.SetCookie("username", uname, -1, "/", "golangwebpagev2.herokuapp.com", true, true)
	c.HTML(http.StatusOK, "index.html", datas)
}

func GetUsername(c *gin.Context) (string, error) {
	uname, err := c.Cookie("username")
	return uname, err
}
