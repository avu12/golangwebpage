package dailymail

import (
	"fmt"
	"log"
	"net/smtp"
	"os"
	"time"

	"github.com/avu12/golangwebpage/database"
	"github.com/avu12/golangwebpage/webpagego/internal/controller/nameday"
)

func SendDailyMail() {
	sender := os.Getenv("MAIL_SENDER")
	pwd := os.Getenv("GMAIL_APPPWD")
	for range time.NewTicker(time.Hour * 24).C {
		quote, err := database.SelectQuote()
		if err != nil {
			quote = "Sorry, no quotes today!"
		}

		data := nameday.GetNamedayNow()
		day, ok := data["day"].(int)
		if ok != true {
			log.Println("Error in type assertion of day")
		}
		month, ok := data["month"].(int)
		if ok != true {
			log.Println("Error in type assertion of month")
		}
		to := database.SelectAllConfirmedMail()
		if to == nil {
			log.Println("Error happened during filling the recipients field, so do not send emails.")
			continue
		}
		auth := smtp.PlainAuth("", sender, pwd, "smtp.gmail.com")
		//TODO: Remake it with mailyak!
		msg := []byte(fmt.Sprintf("Subject: Daily Mail Service from Tamás!\r\n"+
			"\r\n"+
			"Today date is %d %d. Today these people has namedays: %s \r\n  Quote of the day: %s", time.Month(month), day, data["namedays"], quote))
		err = smtp.SendMail("smtp.gmail.com:587", auth, sender, to, msg)
		if err != nil {
			log.Println(err)
		}
	}

}
