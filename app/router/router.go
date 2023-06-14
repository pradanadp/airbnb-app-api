package router

import (
	"be-api/app/middlewares"
	bookingControllerInit "be-api/features/booking/controller"
	bookingRepoInit "be-api/features/booking/data"
	bookingServiceInit "be-api/features/booking/service"

	homestayControllerInit "be-api/features/homestay/controller"
	homestayRepoInit "be-api/features/homestay/data"
	homestayServiceInit "be-api/features/homestay/service"

	userController "be-api/features/user/controller"
	userData "be-api/features/user/data"
	userService "be-api/features/user/service"

	reviewController "be-api/features/review/controller"
	reviewData "be-api/features/review/data"
	reviewService "be-api/features/review/service"

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
		homestaysGroup.GET("", homestayControllerAPI.ReadAllHomestay)
		homestaysGroup.GET("/:homestay_id", homestayControllerAPI.ReadHomestay)
		homestaysGroup.PUT("/:homestay_id", homestayControllerAPI.UpdateHomestay)
		homestaysGroup.DELETE("/:homestay_id", homestayControllerAPI.DeleteHomestay)
	}

	// Booking router
	bookingRepo := bookingRepoInit.New(db)
	bookingService := bookingServiceInit.New(bookingRepo)
	bookingControllerAPI := bookingControllerInit.New(bookingService)

	bookingsGroup := e.Group("/bookings")
	{
		bookingsGroup.POST("", bookingControllerAPI.CreateBooking)
	}

	//User Router
	UserData := userData.New(db)
	UserService := userService.New(UserData)
	UserController := userController.New(UserService)

	e.POST("/login", UserController.LoginUser)
	e.POST("/users", UserController.AddUser)
	e.GET("/users", UserController.GetUser, middlewares.JWTMiddleware())
	e.DELETE("/users", UserController.DeleteUser, middlewares.JWTMiddleware())
	e.PUT("/users", UserController.UpdateUser, middlewares.JWTMiddleware())
	e.PUT("/upgrades", UserController.UpgradeUser, middlewares.JWTMiddleware())


	//review Router
	ReviewData := reviewData.New(db)
	reviewService := reviewService.New(ReviewData)
	ReviewController := reviewController.New(reviewService)

	e.POST("/reviews", ReviewController.AddReview,middlewares.JWTMiddleware())
	e.DELETE("/reviews/:review_id", ReviewController.DeleteReview,middlewares.JWTMiddleware())
	e.GET("/homestays/:homestay_id/reviews", ReviewController.GetAllReview)
}
