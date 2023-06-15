package controller

import (
	models "be-api/features"
	bookingInterface "be-api/features/booking"
	"be-api/utils"
	"fmt"
	"net/http"
	"strings"

	"github.com/labstack/echo/v4"
)

type bookingController struct {
	bookingService bookingInterface.BookingService
}

func New(service bookingInterface.BookingService) *bookingController {
	return &bookingController{
		bookingService: service,
	}
}

func (bc *bookingController) CreateBooking(c echo.Context) error {
	var booking models.BookingEntity
	err := c.Bind(&booking)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind booking data", nil))
	}

	bookingID, orderID, err := bc.bookingService.CreateBooking(booking)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse(err.Error(), nil))
		}
	}

	fmt.Println(bookingID)
	fmt.Println(orderID)

	booking.ID = bookingID
	booking.OrderID = orderID
	bookingResponse := BookingEntityToResponse(booking)

	return c.JSON(http.StatusOK, utils.SuccessResponse("booking created successfully, complete the payment immediately", bookingResponse))
}

func (bc *bookingController) DeleteBooking(c echo.Context) error {
	return nil
}
