package fakers

import (
	"github.com/bxcodec/faker/v4"
	"github.com/google/uuid"
	"github.com/ilhamadikusuma31/ditoko/app/models"
	"gorm.io/gorm"
	"time"
)

func UserFaker(db *gorm.DB) *models.User {
	return &models.User{
		ID:            uuid.New().String(),
		Addresses:     nil,
		FirstName:     faker.FirstName(),
		LastName:      faker.LastName(),
		Email:         faker.Email(),
		Password:      faker.PASSWORD,
		RememberToken: "",
		CreatedAt:     time.Time{},
		UpdatedAt:     time.Time{},
		DeleteAt:      gorm.DeletedAt{},
	}
}
