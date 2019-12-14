package services

import (
	"CustomIM/models"
)

func Config() (*models.Emails, error) {
	email := new(models.Emails)
	if err := db.First(email).Error; err != nil {
		return nil, err
	}
	return email, nil
}


