package repository

import (
	"github.com/Greeshmanth-Pulicallu/currency/config"
	"github.com/Greeshmanth-Pulicallu/currency/models"
)

func AddUserToDB(userID, hash string) error {
	user := models.Users{
		UserID: userID,
		Hash:   string(hash),
	}

	if err := config.DB.Create(&user).Error; err != nil {
		return err
	}

	return nil
}

func VerifyUserExistsInDB(userID string) (models.Users, error) {
	var user models.Users
	if err := config.DB.Where("user_id = ?", userID).First(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
