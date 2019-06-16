package main

import (
	"log"
	"os"
	"strconv"
	"time"

	_ "compgen-api-docs/routers"

	"github.com/astaxie/beego"
	"github.com/astaxie/beego/orm"
	"github.com/astaxie/beego/plugins/cors"
	_ "github.com/go-sql-driver/mysql"
)

func init() {
	orm.RegisterDriver("mysql", orm.DRMySQL)
	dbURL := beego.AppConfig.String("mysqluser") + ":" + beego.AppConfig.String("mysqlpass") + "@tcp(" + beego.AppConfig.String("mysqlurls") + ":3306)/" + beego.AppConfig.String("mysqldb")
	orm.RegisterDataBase("default", "mysql", dbURL)
	orm.DefaultTimeLoc, _ = time.LoadLocation("America/Sao_Paulo")
	orm.Debug = true
}

func main() {
	if os.Getenv("PORT") != "" {
		log.Println("Env $PORT :", os.Getenv("PORT"))
		port, err := strconv.Atoi(os.Getenv("PORT"))
		if err != nil {
			log.Fatal(err)
			log.Fatal("$PORT must be set")
		}
		log.Println("port : ", port)
		beego.BConfig.Listen.HTTPSPort = port
	}
	if os.Getenv("BEEGO_ENV") != "" {
		log.Println("Env $BEEGO_ENV :", os.Getenv("BEEGO_ENV"))
		beego.BConfig.RunMode = os.Getenv("BEEGO_ENV")
	}

	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		AllowOrigins:     []string{"localhost:8080", "localhost:3000", "*"},
		AllowMethods:     []string{"PUT", "PATCH", "GET", "POST", "OPTIONS", "DELETE"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "X-Requested-With", "Content-Type", "Accept", "Connection", "Upgrade", "Token", "Authorization", "Websocket", "Set-Cookie", "withCredentials"},
		ExposeHeaders:    []string{"Content-Length", "Access-Control-Allow-Origin", "Origin", "Connection", "Upgrade", "Token", "Authorization", "Websocket", "Set-Cookie", "withCredentials"},
		AllowCredentials: true,
	}))

	beego.Run()
}
