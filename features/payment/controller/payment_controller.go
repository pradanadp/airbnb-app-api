package controller

import (
	"be-api/features"
	paymentInterface "be-api/features/payment"
	"be-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func (handler *paymentController) AddPayment(c echo.Context) error {

	IdBooking := c.Param("booking_id")
	idConv, errConv := strconv.Atoi(IdBooking)
	if errConv != nil {
		if strings.Contains(errConv.Error(), "bind failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(errConv.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to bind data. "+errConv.Error(), nil))
		}
	}

	var payment features.ResponMidtrans
	err := c.Bind(&payment)
	if err != nil {
		if err == echo.ErrBadRequest {
			return c.JSON(http.StatusBadRequest, utils.FailResponse("error bind payload "+err.Error(), nil))
		}
	}

	idOrder, errCreate := handler.paymentService.CreatePayment(payment,uint(idConv))
	if errCreate != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(errCreate.Error(), nil))
	}

	data,errGet:=handler.paymentService.GetPayment(idOrder)
	if errGet != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(errGet.Error(), nil))
	}

	dataResult := PaymentEntityToMidstrans(data)



	return c.JSON(http.StatusOK, utils.SuccessResponse("payment add successfully", dataResult))

}



func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}
