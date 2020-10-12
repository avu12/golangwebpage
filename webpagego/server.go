package webpagego

import (
	"log"
	"net"
	"net/http"
	"os"
	"time"

	"github.com/avu12/golangwebpage/webpagego/internal/dailymail"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/tatsushid/go-fastping"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.LoadHTMLGlob("static/*")

}

//StartApp is the main entrypoint to the app
func StartApp() {
	//Mapping urls with handlers
	mapUrls()
	Testping()
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
func Testping() {

	p := fastping.NewPinger()
	ra, err := net.ResolveIPAddr("ip6:icmp", os.Getenv("OWN_IPV6"))
	if err != nil {
		log.Println(err)
		os.Exit(1)
	}
	p.AddIPAddr(ra)
	p.OnRecv = func(addr *net.IPAddr, rtt time.Duration) {
		log.Printf("IP Addr: %s receive, RTT: %v\n", addr.String(), rtt)
	}
	p.OnIdle = func() {
		log.Println("finish")
	}
	err = p.Run()
	if err != nil {
		log.Println(err)
	}
}
