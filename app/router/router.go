package router

import (
	homestayControllerInit "be-api/features/homestay/controller"
	homestayRepoInit "be-api/features/homestay/data"
	homestayServiceInit "be-api/features/homestay/service"

	userController "be-api/features/user/controller"
	userData "be-api/features/user/data"
	userService "be-api/features/user/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// Homestay Router
	homestayRepo := homestayRepoInit.New(db)
	homestayService := homestayServiceInit.New(homestayRepo)
	homestayControllerAPI := homestayControllerInit.New(homestayService)

	homestaysGroup := e.Group("/homestays")
	{
		homestaysGroup.POST("", homestayControllerAPI.CreateHomestay)
		homestaysGroup.GET("/:homestay_id", homestayControllerAPI.ReadHomestay)
	}

	//User Router
	UserData := userData.New(db)
	UserService := userService.New(UserData)
	UserController := userController.New(UserService)

	e.POST("/login",UserController.LoginUser)
	e.POST("/users",UserController.AddUser)
}
