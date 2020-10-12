package webpagego

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/avu12/golangwebpage/webpagego/internal/dailymail"
	"github.com/gin-gonic/contrib/static"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	"github.com/pkg/sftp"
	"golang.org/x/crypto/ssh"
	"golang.org/x/net/proxy"
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
func TestSFTP() {

	remote := os.Getenv("OWN_IPV6")
	port := ":22"
	pass := os.Getenv("SFTPTESTPWD")
	user := os.Getenv("SFTPTESTUSER")

	config := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	conn, err := ssh.Dial("tcp", remote+port, &config)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client, err := sftp.NewClient(conn)
	if err != nil {
		log.Fatal(err)
	}
	defer client.Close()
}
func TestFixieSFTP() {
	remote := os.Getenv("OWN_IPV6")
	port := ":22"
	pass := os.Getenv("SFTPTESTPWD")
	user := os.Getenv("SFTPTESTUSER")

	config := ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{
			ssh.Password(pass),
		},
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	}
	_ = config
	fixie_data := strings.Split(os.Getenv("FIXIE_SOCKS_HOST"), "@")
	fixie_addr := fixie_data[1]
	auth_data := strings.Split(fixie_data[0], ":")
	auth := proxy.Auth{
		User:     auth_data[0],
		Password: auth_data[1],
	}
	dialer, err := proxy.SOCKS5("tcp", fixie_addr, &auth, proxy.Direct)
	if err != nil {
		log.Println(os.Stderr, "can't connect to the proxy:", err)

	}
	conn, err := dialer.Dial("tcp", remote+port)
	if err != nil {
		log.Println(err)
	}
	defer conn.Close()

}
