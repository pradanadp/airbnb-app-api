package router

import (
	"be-api/app/middlewares"
	bookingControllerInit "be-api/features/booking/controller"
	bookingRepoInit "be-api/features/booking/data"
	bookingServiceInit "be-api/features/booking/service"

	homestayControllerInit "be-api/features/homestay/controller"
	homestayRepoInit "be-api/features/homestay/data"
	homestayServiceInit "be-api/features/homestay/service"

	imageControllerInit "be-api/features/image/controller"
	imageRepoInit "be-api/features/image/data"
	imageServiceInit "be-api/features/image/service"

	userController "be-api/features/user/controller"
	userData "be-api/features/user/data"
	userService "be-api/features/user/service"

	reviewController "be-api/features/review/controller"
	reviewData "be-api/features/review/data"
	reviewService "be-api/features/review/service"

	paymentController "be-api/features/payment/controller"
	paymentData "be-api/features/payment/data"
	paymentService "be-api/features/payment/service"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

func InitRouter(db *gorm.DB, e *echo.Echo) {
	// Homestay Router
	homestayRepo := homestayRepoInit.New(db)
	homestayService := homestayServiceInit.New(homestayRepo)
	homestayControllerAPI := homestayControllerInit.New(homestayService)

	// Image router
	imageRepo := imageRepoInit.New(db)
	imageService := imageServiceInit.New(imageRepo)
	imageControllerAPI := imageControllerInit.New(imageService)

	homestaysGroup := e.Group("/homestays")
	{
		homestaysGroup.POST("", homestayControllerAPI.CreateHomestay, middlewares.JWTMiddleware())
		homestaysGroup.GET("", homestayControllerAPI.ReadAllHomestay)
		homestaysGroup.GET("/:homestay_id", homestayControllerAPI.ReadHomestay)
		homestaysGroup.PUT("/:homestay_id", homestayControllerAPI.UpdateHomestay)
		homestaysGroup.DELETE("/:homestay_id", homestayControllerAPI.DeleteHomestay)
		homestaysGroup.POST("/:homestay_id/images", imageControllerAPI.UploadHomestayPhotos)
		homestaysGroup.POST("/:homestay_id/images/local", imageControllerAPI.UploadHomestayPhotosLocal)
		homestaysGroup.DELETE("/:homestay_id/images/:image_id", imageControllerAPI.DeleteImage)
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
	e.POST("/users/profile-picture", UserController.UploadProfilePicture, middlewares.JWTMiddleware())
	e.POST("/users/host-doc", UserController.UploadHostDoc, middlewares.JWTMiddleware())

	//review Router
	ReviewData := reviewData.New(db)
	reviewService := reviewService.New(ReviewData)
	ReviewController := reviewController.New(reviewService)

	e.POST("/reviews", ReviewController.AddReview, middlewares.JWTMiddleware())
	e.DELETE("/reviews/:review_id", ReviewController.DeleteReview, middlewares.JWTMiddleware())
	e.GET("/homestays/:homestay_id/reviews", ReviewController.GetAllReview)

	PaymentData := paymentData.New(db)
	PaymentService := paymentService.New(PaymentData)
	PaymentController := paymentController.New(PaymentService)

	e.POST("/payments", PaymentController.AddPayment)
}
