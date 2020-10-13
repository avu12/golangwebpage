package login

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/database"
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
		log.Println("You logged in!")
	} else {
		log.Println("You did not log in!")
	}
	c.HTML(http.StatusOK, "index.html", nil)

}
