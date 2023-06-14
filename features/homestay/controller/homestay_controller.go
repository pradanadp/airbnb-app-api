package controller

import (
	models "be-api/features"
	homestayInterface "be-api/features/homestay"
	"be-api/utils"
	"net/http"
	"strconv"
	"strings"

	"github.com/labstack/echo/v4"
)

type homestayController struct {
	homestayService homestayInterface.HomestayService
}

func New(service homestayInterface.HomestayService) *homestayController {
	return &homestayController{
		homestayService: service,
	}
}

func (hc *homestayController) CreateHomestay(c echo.Context) error {
	var homestay models.HomestayEntity

	err := c.Bind(&homestay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind homestay data", nil))
	}

	homestayID, err := hc.homestayService.CreateHomestay(homestay)
	if err != nil {
		if strings.Contains(err.Error(), "insert failed") {
			return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
		} else {
			return c.JSON(http.StatusInternalServerError, utils.FailResponse("failed to insert data. "+err.Error(), nil))
		}
	}

	homestay.ID = homestayID
	homestayResponse := HomestayEntityToResponse(homestay)

	return c.JSON(http.StatusOK, utils.SuccessResponse("homestay created successfully", homestayResponse))
}

func (hc *homestayController) ReadHomestay(c echo.Context) error {
	idParam := c.Param("homestay_id")
	homestayID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	homestay, err := hc.homestayService.GetHomestay(uint(homestayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("homestay not found", nil))
	}

	homestayResponse := HomestayEntityToResponse(homestay)

	return c.JSON(http.StatusOK, utils.SuccessResponse("homestay retrieved successfully", homestayResponse))
}

func (hc *homestayController) ReadAllHomestay(c echo.Context) error {
	homestays, err := hc.homestayService.GetAllHomestay()
	if err != nil {
		return c.JSON(http.StatusNotFound, utils.FailResponse("homestays not found", nil))
	}

	var homestayResponses []HomestayResponse
	for _, homestay := range homestays {
		homestayResponses = append(homestayResponses, HomestayEntityToResponse(homestay))
	}

	return c.JSON(http.StatusOK, utils.SuccessResponse("homestays retrieved successfully", homestayResponses))
}

func (hc *homestayController) UpdateHomestay(c echo.Context) error {
	var updatedHomestay models.HomestayEntity
	err := c.Bind(&updatedHomestay)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("failed to bind homestay data", nil))
	}

	idParam := c.Param("homestay_id")
	homestayID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	err = hc.homestayService.UpdatedHomestay(uint(homestayID), updatedHomestay)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, utils.FailResponse("status internal error", nil))
	}

	homestay, err := hc.homestayService.GetHomestay(uint(homestayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("homestay not found", nil))
	}

	homestayResponse := HomestayEntityToResponse(homestay)

	return c.JSON(http.StatusOK, utils.SuccessResponse("homestay updated successfully", homestayResponse))
}

func (hc *homestayController) DeleteHomestay(c echo.Context) error {
	idParam := c.Param("homestay_id")
	homestayID, err := strconv.Atoi(idParam)
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("invalid homestay ID", nil))
	}

	homestay, err := hc.homestayService.GetHomestay(uint(homestayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse("homestay not found", nil))
	}

	// Delete homestay data from database
	err = hc.homestayService.DeleteHomestay(uint(homestayID))
	if err != nil {
		return c.JSON(http.StatusBadRequest, utils.FailResponse(err.Error(), nil))
	}

	homestayResponse := HomestayEntityToResponse(homestay)

	return c.JSON(http.StatusOK, utils.SuccessResponse("homestay deleted successfully", homestayResponse))
}
