package review

import "be-api/features"

type ReviewDataInterface interface {
	Insert(input features.ReviewEntity, costumer_id uint) (uint,error)
	Delete(review_id uint) error
	SelectId(review_id uint) (features.ReviewEntity,error)
	SelectAll(homestay_id uint) ([]features.ReviewEntity, error)
}

type ReviewServiceInterface interface {
	AddRiview(input features.ReviewEntity, costumer_id uint) (uint,error)
	DeleteRiview(review_id uint) error
	GetId(review_id uint) (features.ReviewEntity,error)
	GetAll(homestay_id uint) ([]features.ReviewEntity, error)
}