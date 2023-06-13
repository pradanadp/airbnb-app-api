package data

import (
	"be-api/features"
	"be-api/features/review"
	"errors"

	"gorm.io/gorm"
)

type ReviewData struct {
	db *gorm.DB
}

// Delete implements review.ReviewDataInterface
func (repo *ReviewData) Delete(review_id uint) error {
	var review features.Review

	FoundRivewsId := repo.db.First(&review, "id=?",review_id)
	if FoundRivewsId.Error != nil {
		return errors.New(FoundRivewsId.Error.Error() + ", failed to get review id")
	}
	tx := repo.db.Delete(&review, review_id)
	if tx.Error != nil {
		return tx.Error
	}

	return nil
}

// Insert implements review.ReviewDataInterface
func (repo *ReviewData) Insert(input features.ReviewEntity, costumer_id uint) (uint, error) {
	ReviewModel := features.ReviewEntityToModel(input)

	Costumers := features.User{}
	FoundCostumersId := repo.db.First(&Costumers, "id=?", costumer_id)
	if FoundCostumersId.Error != nil {
		return 0, errors.New(FoundCostumersId.Error.Error() + ", failed to get costumers id")
	}
	ReviewModel.CustomerID = costumer_id
	createOpr := repo.db.Create(&ReviewModel)
	if createOpr.Error != nil {
		return 0, createOpr.Error
	}

	if createOpr.RowsAffected == 0 {
		return 0, errors.New("failed to insert, row affected is 0")
	}

	return ReviewModel.ID, nil
}

func New(db *gorm.DB) review.ReviewDataInterface {
	return &ReviewData{
		db: db,
	}

}
