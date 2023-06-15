package homestay

import (
	models "be-api/features"
)

type HomestayRepository interface {
	Insert(homestay models.HomestayEntity) (uint, error)
	Select(homestayID uint) (models.HomestayEntity, error)
	SelectAll() ([]models.HomestayEntity, error)
	SelectAllByHostID(hostID uint) ([]models.HomestayEntity, error)
	Update(homestayID uint, updatedHomestay models.HomestayEntity) error
	Delete(homestayID uint) error
}

type HomestayService interface {
	CreateHomestay(homestay models.HomestayEntity) (uint, error)
	GetHomestay(homestayID uint) (models.HomestayEntity, error)
	GetAllHomestay() ([]models.HomestayEntity, error)
	GetAllHomestayByHostID(hostID uint) ([]models.HomestayEntity, error)
	UpdatedHomestay(homestayID uint, updatedHomestay models.HomestayEntity) error
	DeleteHomestay(homestayID uint) error
}
