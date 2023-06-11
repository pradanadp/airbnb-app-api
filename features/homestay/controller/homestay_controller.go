package controller

import (
	homestayInterface "be-api/features/homestay"
)

type homestayController struct {
	homestayService homestayInterface.HomestayService
}

func New(service homestayInterface.HomestayService) *homestayController {
	return &homestayController{
		homestayService: service,
	}
}
