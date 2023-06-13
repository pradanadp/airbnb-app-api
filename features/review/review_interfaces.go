package review

import "be-api/features"

type ReviewDataInterface interface {
	Insert(input features.ReviewEntity, costumer_id uint) (uint,error)
	Delete(review_id uint) error
}

type ReviewServiceInterface interface {
	AddRiview(input features.ReviewEntity, costumer_id uint) (uint,error)
	DeleteRiview(review_id uint) error
}