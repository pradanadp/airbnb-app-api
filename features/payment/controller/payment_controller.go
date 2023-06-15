package controller

import (
	paymentInterface "be-api/features/payment"

	"github.com/labstack/echo/v4"
)

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func (handler *paymentController) AddUser(c echo.Context) error {
	return nil
}

func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}
