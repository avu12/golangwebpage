package webpagego

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/avu12/golangwebpage/webpagego/internal/dailymail"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

var (
	router *gin.Engine
	opt    sessions.Options
)

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("static/*")
	/*store, err := sessions.NewRedisStore(10, "tcp", os.Getenv("REDIS_URL"), "", []byte("secret"))
	if err != nil {

		log.Println("Problem with redis store in init", err)
	}
	opt.MaxAge = 86400
	opt.Path = "/"
	opt.Secure = true
	opt.HttpOnly = true

	store.Options(opt)

	//Redis need correct config to use!

	router.Use(sessions.Sessions("mysession", store))
	log.Println("No problem with redis store in init")*/

}

//StartApp is the main entrypoint to the app
func StartApp() {
	//Mapping urls with handlers
	mapUrls()

	//sending email every day
	go dailymail.SendDailyMail()

	//for infinite running:
	go NoSleep()

	port := os.Getenv("PORT")

	if port == "" {
		log.Fatal("$PORT must be set")
	}
	router.Use(static.Serve("/", static.LocalFile("./static", true)))
	err := router.Run(":" + port)
	if err != nil {
		panic(err)
	}

	/* FOR LOCAL: err := router.Run(":3000")
	if err != nil {
		panic(err)
	}*/

}

//NoSleep is to prevent Heroku to stop the app automatically, only needed if Free dyno used.
func NoSleep() {
	for range time.NewTicker(time.Minute * 15).C {
		log.Println(time.Now(), "I am a still not sleeping dyno!")
		req, err := http.NewRequest("GET", "https://golangwebpagev2.herokuapp.com/", nil)
		if err != nil {
			log.Println("ERROR in self request!")
		}
		client := http.Client{}
		resp, err := client.Do(req)
		//log.Println("RESP:", resp)
		if err != nil {
			log.Println("ERROR in self resp!")
		}
		resp.Body.Close()
	}
}
