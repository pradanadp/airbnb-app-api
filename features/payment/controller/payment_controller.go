package controller

import (
	"be-api/features"
	paymentInterface "be-api/features/payment"
	"be-api/utils"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func (handler *paymentController) AddPayment(c echo.Context) error {
	idParam := c.Param("booking_id")
	BookingID, errParam := strconv.Atoi(idParam)
	if errParam != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid Booking ID", nil))
	}

	var payment features.ResponMidtrans
	err := c.Bind(&payment)
	if err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}
	id, errCreate := handler.paymentService.CreatePayment(payment,uint(BookingID))
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(errCreate.Error(), nil))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("review add successfully", id))

}

func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}
