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

	ratingsNew, errRatings:=AverageRatingsDelete(repo.db,review_id,review.HomestayID)
	if errRatings != nil {
		return errRatings
	}

	tx := repo.db.Delete(&review, review_id)
	if tx.Error != nil {
		return tx.Error
	}

	errUpdateRating:=UpdateRatings(repo.db,review.HomestayID,ratingsNew)
	if errUpdateRating != nil{
		return errUpdateRating
	}

	return nil
}

// Insert implements review.ReviewDataInterface
func (repo *ReviewData) Insert(input features.ReviewEntity, costumer_id uint) (uint, error) {
	ReviewModel := features.ReviewEntityToModel(input)

	ratingUpdate,err := AverageRatingsInsert(repo.db, ReviewModel.Ratings,ReviewModel.HomestayID)
	if err != nil{
		return 0, err
	}

	errUpdateRating:=UpdateRatings(repo.db,ReviewModel.HomestayID,ratingUpdate)
	if errUpdateRating != nil{
		return 0,errUpdateRating
	}

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

func AverageRatingsInsert(db *gorm.DB, inputRating float64, homestay_id uint) (float64 ,error){

	var total float64
    var count int

	review := []features.Review{}
	FoundhomestayId := db.Find(&review, "homestay_id=?",homestay_id)
	if FoundhomestayId.Error != nil {
		return 0,errors.New(FoundhomestayId.Error.Error() + ", failed to get review id")
	}
	
	var tampungRating []float64
	for _,lastRatings := range review{
		tampungRating = append(tampungRating, lastRatings.Ratings)
	}
	tampungRating = append(tampungRating, inputRating)

	for _, value := range tampungRating{
        total += value
        count++
	}

	var average float64

    if count > 0 {
        average = total / float64(count)
    }

	return average,nil
}

func AverageRatingsDelete(db *gorm.DB,review_id uint, homestay_id uint) (float64 ,error){

	var total float64
    var count int

	review := []features.Review{}

	FoundhomestayId := db.Find(&review, "homestay_id=?",homestay_id)
	if FoundhomestayId.Error != nil {
		return 0,errors.New(FoundhomestayId.Error.Error() + ", failed to get review id")
	}
	
	var tampungRating []float64
	for _,lastRatings := range review{
		tampungRating = append(tampungRating, lastRatings.Ratings)
	}
	reviewId := features.Review{}
	FoundreviewId := db.First(&reviewId, "homestay_id=? AND id=?",homestay_id,review_id)
	if FoundreviewId.Error != nil {
		return 0,errors.New(FoundreviewId.Error.Error() + ", failed to get review id")
	}

	for _, value := range tampungRating{
        total += value
        count++
	}

	total = total-reviewId.Ratings
	count = count - 1

	var average float64

    if count > 0 {
        average = total / float64(count)
    }

	return average,nil
}

func UpdateRatings(db *gorm.DB, homestayID uint, averageRating float64) error {
	result := db.Model(&features.Homestay{}).Where("id = ?", homestayID).Update("rating", averageRating)
	if result.Error != nil {
		return result.Error
	}
	return nil
}