package controller

import (
	paymentInterface "be-api/features/payment"
)

type paymentController struct {
	paymentService paymentInterface.PaymentService
}

func New(service paymentInterface.PaymentService) *paymentController {
	return &paymentController{
		paymentService: service,
	}
}
