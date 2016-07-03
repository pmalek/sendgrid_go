package main

import (
	"log"
	"os"

	"github.com/buaazp/fasthttprouter"
	"github.com/pmalek/sendgrid_go/mailer"

	"github.com/valyala/fasthttp"
)

var (
	GO_WEB_PORT      = os.Getenv("GO_WEB_PORT")
	SENDGRID_API_KEY = os.Getenv("SENDGRID_API_KEY")
	PASS_FOR_MAIL    = os.Getenv("PASS_FOR_MAIL")
	EMAIL            = os.Getenv("EMAIL")
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lmicroseconds | log.Lshortfile)
	log.Printf("Starting server at %q", GO_WEB_PORT)

	router := fasthttprouter.New()
	router.GET("/:password", requestHandler)

	if err := fasthttp.ListenAndServe(GO_WEB_PORT, router.Handler); err != nil {
		log.Fatalf("Error in ListenAndServe: %s", err)
	}
}

func requestHandler(ctx *fasthttp.RequestCtx, p fasthttprouter.Params) {
	log.Printf("Request method is %q", ctx.Method())
	log.Printf("RequestURI is %q", ctx.RequestURI())
	log.Printf("Requested path is %q", ctx.Path())
	log.Printf("Host is %q", ctx.Host())
	log.Printf("Query string is %q", ctx.QueryArgs())
	log.Printf("User-Agent is %q", ctx.UserAgent())
	log.Printf("Connection has been established at %s", ctx.ConnTime())
	log.Printf("Request has been started at %s", ctx.Time())
	log.Printf("Serial request number for the current connection is %d", ctx.ConnRequestNum())
	log.Printf("Your ip is %q", ctx.RemoteIP())
	log.Printf("Raw request is:\n%s", &ctx.Request)

	ctx.SetContentType("text/plain; charset=utf8")

	if len(EMAIL) > 0 && len(PASS_FOR_MAIL) > 0 && p.ByName("password") == PASS_FOR_MAIL {
		go mailer.SendHelloEmail(SENDGRID_API_KEY, "You guessed the password :)", EMAIL, EMAIL)
		go func() { log.Println("Sending an email :)") }()
	} else {
		go func() { log.Println("Wrong password for sending an email -.-") }()
	}
}
