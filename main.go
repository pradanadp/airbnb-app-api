package main

import (
	"be-api/app/config"
	"be-api/app/database"
	"be-api/app/router"
	"be-api/features/image/controller"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	cfg := config.InitConfig()
	mysql := database.InitMysql(cfg)
	database.InitialMigration(mysql)

	database.InitUsersData(mysql)
	database.InitHomestaysData(mysql)
	database.InitReviewsData(mysql)

	e := echo.New()

	e.Pre(middleware.RemoveTrailingSlash())
	e.Use(middleware.CORS())
	e.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: `[${time_rfc3339}] ${status} ${method} ${host}${path} ${latency_human}` + "\n",
	}))

	router.InitRouter(mysql, e)
	e.POST("/users/images", controller.UploadHostDoc)

	e.Logger.Fatal(e.Start(":8080"))
}
