package webpagego

import (
	"html/template"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Masterminds/sprig"

	"github.com/avu12/golangwebpage/webpagego/internal/dailymail"
	"github.com/gin-gonic/contrib/sessions"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
	_ "github.com/lib/pq"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.SetFuncMap(template.FuncMap{
		"kindIs": sprig.GenericFuncMap,
	})

	router.LoadHTMLGlob("static/*")
	opt, err := redis.ParseURL(os.Getenv("REDIS_URL"))
	if err != nil {
		log.Println(err)
	}
	store, err := sessions.NewRedisStore(10, "tcp", opt.Addr, opt.Password)
	if err != nil {

		log.Println("Problem with redis store in init", err)
	}

	router.Use(sessions.Sessions("mysession", store))
	log.Println("No problem with redis store in init")

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
