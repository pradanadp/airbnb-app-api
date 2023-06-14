package service

import (
	models "be-api/features"
	homestayInterface "be-api/features/homestay"
	"errors"
	"fmt"
)

type homestayService struct {
	homestayRepository homestayInterface.HomestayRepository
}

// CreateHomestay implements homestay.HomestayService.
func (hs *homestayService) CreateHomestay(homestay models.HomestayEntity) (uint, error) {
	switch {
	case homestay.HostID == 0:
		return 0, errors.New("error, host id is required")
	case homestay.Title == "":
		return 0, errors.New("error, title is required")
	case homestay.Description == "":
		return 0, errors.New("error, description is required")
	case homestay.Location == "":
		return 0, errors.New("error, location is required")
	case homestay.Address == "":
		return 0, errors.New("error, address is required")
	case homestay.Price == 0.0:
		return 0, errors.New("error, price is required")
	case homestay.Facilities == "":
		return 0, errors.New("error, facilities is required")
	}

	homestayID, err := hs.homestayRepository.Insert(homestay)
	if err != nil {
		return 0, fmt.Errorf("%v", err)
	}

	return homestayID, nil
}

// DeleteHomestay implements homestay.HomestayService.
func (hs *homestayService) DeleteHomestay(homestayID uint) error {
	err := hs.homestayRepository.Delete(homestayID)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

// GetAllHomestay implements homestay.HomestayService.
func (hs *homestayService) GetAllHomestay() ([]models.HomestayEntity, error) {
	homestayEntities, err := hs.homestayRepository.SelectAll()
	if err != nil {
		return nil, fmt.Errorf("error: %v", err)
	}

	return homestayEntities, nil
}

// GetHomestay implements homestay.HomestayService.
func (hs *homestayService) GetHomestay(homestayID uint) (models.HomestayEntity, error) {
	homestayEntity, err := hs.homestayRepository.Select(homestayID)
	if err != nil {
		return models.HomestayEntity{}, fmt.Errorf("error: %v", err)
	}

	return homestayEntity, nil
}

// UpdatedHomestay implements homestay.HomestayService.
func (hs *homestayService) UpdatedHomestay(homestayID uint, updatedHomestay models.HomestayEntity) error {
	err := hs.homestayRepository.Update(homestayID, updatedHomestay)
	if err != nil {
		return fmt.Errorf("error: %v", err)
	}

	return nil
}

func New(repo homestayInterface.HomestayRepository) homestayInterface.HomestayService {
	return &homestayService{
		homestayRepository: repo,
	}
}
