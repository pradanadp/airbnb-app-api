package service

import (
	"be-api/features"
	"be-api/features/review"

	"github.com/go-playground/validator"
)

type ReviewService struct {
	reviewData review.ReviewDataInterface
	validate   *validator.Validate
}

// GetId implements review.ReviewServiceInterface
func (service *ReviewService) GetId(review_id uint) (features.ReviewEntity,error) {
	reviewUser, err := service.reviewData.SelectId(review_id)
	if err != nil {
		return features.ReviewEntity{}, err
	}
	return reviewUser, err
}

// DeleteRiview implements review.ReviewServiceInterface
func (service *ReviewService) DeleteRiview(review_id uint) error {
	err := service.reviewData.Delete(review_id)
	if err != nil {
		return err
	}
	return nil
}

// AddRiview implements review.ReviewServiceInterface
func (service *ReviewService) AddRiview(input features.ReviewEntity, costumer_id uint) (uint, error) {
	if errValidate := service.validate.Struct(input); errValidate != nil {
		return 0, errValidate
	}

	id_homestay, err := service.reviewData.Insert(input, costumer_id)
	if err != nil {
		return 0, err
	}
	return id_homestay, nil
}

// GetAll implements review.ReviewServiceInterface
func (service *ReviewService) GetAll(homestay_id uint) ([]features.ReviewEntity, error) {
	reviewUser, err := service.reviewData.SelectAll(homestay_id)
	if err != nil {
		return []features.ReviewEntity{}, err
	}
	return reviewUser, err
}

func New(ReviewData review.ReviewDataInterface) review.ReviewServiceInterface {
	return &ReviewService{
		reviewData: ReviewData,
		validate:   validator.New(),
	}
}
