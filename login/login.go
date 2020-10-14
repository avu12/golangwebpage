package login

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/types"
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
	Data := types.Logindata{}
	if pwd == pwdhashindb {
		Data.Username = name
		Data.Isloggedin = true
		AddUserToRedis(c, name)
		c.HTML(http.StatusOK, "index.html", Data)
	} else {
		Data.Username = ""
		Data.Isloggedin = false
		c.HTML(http.StatusOK, "index.html", Data)
	}

}

func AddUserToRedis(c *gin.Context, username string) {
	session := sessions.Default(c)
	//set true for logged in
	session.Set("user_username", username)
	session.Save()
}
