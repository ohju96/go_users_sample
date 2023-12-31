package main

import (
	"ginSample/config"
	"ginSample/config/db"
	"ginSample/router"
	"ginSample/validator"
	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()

	Init(app)

	app.Run()
}

func Init(app *gin.Engine) {
	// toml
	toml := config.InitToml("./config/config.toml")

	// db
	db.InitMySQL(&toml)
	db.InitRedis(&toml)

	// router
	router.MainRouter(app)

	// == VALIDATE == //
	validator.InitValidate()
}
