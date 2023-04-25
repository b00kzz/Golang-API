package repo

import (
	"ticket/goapi/internal/core/port"

	"gorm.io/gorm"
)

type reviewRepo struct {
	db *gorm.DB
}

func NewReviewRepo(db *gorm.DB) port.ReviewRepo {
	return reviewRepo{
		db: db,
	}
}

func (c reviewRepo) GetAll() ([]port.Review, error) {
	reviews := []port.Review{}
	err := c.db.Find(&reviews).Error
	if err != nil {
		return nil, err
	}
	return reviews, nil
}

func (c reviewRepo) GetById(id int) (*port.Review, error) {
	review := port.Review{}
	err := c.db.First(&review, id).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (c reviewRepo) Create(review port.Review) (*port.Review, error) {
	err := c.db.Create(&review).Error
	if err != nil {
		return nil, err
	}
	return &review, nil
}

func (c reviewRepo) Update(id int, review port.Review) error {
	err := c.db.Model(&port.Review{}).Where("rev_id = ?", id).Updates(review).Error
	if err != nil {
		return err
	}
	return nil
}

func (c reviewRepo) Delete(id int) error {
	err := c.db.Delete(&port.Review{}, id).Error
	if err != nil {
		return err
	}
	return nil
}
