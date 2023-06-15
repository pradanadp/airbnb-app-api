package controller

import (
	"be-api/app/middlewares"
	"be-api/features"
	paymentInterface "be-api/features/payment"
	"be-api/utils"
	"net/http"

	"github.com/labstack/echo/v4"
)

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func (handler *paymentController) AddPayment(c echo.Context) error {
	id := middlewares.ExtracTokenUserId(c)

	// dataBooking,errBooking := handler.paymentService.GetPayment(uint(id))
	// if errBooking != nil {
	// 	if errBooking == echo.ErrBadRequest {
	// 		return c.JSON(http.StatusBadRequest, utils.FailResponse("error get payload "+errBooking.Error(), nil))
	// 	}
	// }
	// bookingID := dataBooking.Booking.ID 

	var payment features.ResponMidtrans
	err := c.Bind(&payment)
	if err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}

	idOrder, errCreate := handler.paymentService.CreatePayment(payment,uint(id))
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("review add successfully", idOrder))
}



func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}