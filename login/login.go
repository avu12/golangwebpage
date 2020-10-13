package login

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/types"
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
	data := types.Logindata{}
	if pwd == pwdhashindb {
		data.Username = name
		data.Isloggedin = true
		c.HTML(http.StatusOK, "index.html", data)
	} else {
		data.Username = ""
		data.Isloggedin = false
		c.HTML(http.StatusOK, "index.html", data)
	}

}