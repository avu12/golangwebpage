package mail

import (
	"crypto/sha256"
	"encoding/hex"
	"log"
	"net/http"
	"net/smtp"
	"os"
	"text/template"

	"github.com/avu12/golangwebpage/database"
	"github.com/domodwyer/mailyak"
	"github.com/gin-gonic/gin"
)

type Url struct {
	Url string
}

func MailHandler(c *gin.Context) {
	email := c.PostForm("mail")
	hash := sha256.Sum256([]byte(email))
	hashnosize := hash[:]
	hashencoded := hex.EncodeToString(hashnosize)
	database.InsertToMailTableWithoutConfirm(email, hashencoded)
	SendConfirmation(email, hashencoded)
	pwd := c.PostForm("pwd")
	hashpwd := sha256.Sum256([]byte(pwd))
	pwd = hex.EncodeToString(hashpwd[:])
	name := c.PostForm("name")
	Registeruserandpasword(name, pwd)
	c.HTML(http.StatusOK, "emailnotify.html", nil)
}

func SendConfirmation(email string, hashencoded string) {
	sender := os.Getenv("MAIL_SENDER")
	pwd := os.Getenv("GMAIL_APPPWD")
	to := []string{}
	to = append(to, email)
	data := Url{
		Url: "https://golangwebpagev2.herokuapp.com/emailregistered/" + hashencoded,
	}
	mail := mailyak.New("smtp.gmail.com:587", smtp.PlainAuth("", sender, pwd, "smtp.gmail.com"))
	mail.To(email)
	mail.From(sender)
	mail.FromName("Tamas")
	mail.Subject("Email verification")
	t, err := template.ParseFiles("./templates/emailverificationtemplate.html")
	if err != nil {
		log.Println(err)
	}
	err = t.Execute(mail.HTML(), data)
	if err != nil {
		log.Println(err)
	}

	err = mail.Send()
	if err != nil {
		log.Println(err)
	}
}

func ConfirmRegistration(c *gin.Context) {
	emailhash := c.Param("emailhash")
	log.Println(emailhash)
	//Data exis
	if len(database.SelectMail(emailhash)) != 0 {
		err := database.UpdateConfirmData(emailhash)
		if err != nil {
			//Something error happened during db update
			log.Println("Problem in DB update")
		}
	} else {
		//no data, show error page:
	}
	c.HTML(http.StatusOK, "emailregistered.html", nil)
}

func Registeruserandpasword(username string, hashpwd string) {
	log.Println(username, hashpwd)
}
