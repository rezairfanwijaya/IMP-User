package database

import (
	"fmt"
	"imp/user"

	"gorm.io/gorm"
)

func MigarationUserSeed(db *gorm.DB) error {
	// if existing
	isExist := isExisting(db)

	userSeeds := user.GenerateSeedUser()
	if !isExist {
		for _, userSeed := range userSeeds {
			if err := db.Create(&userSeed).Error; err != nil {
				return fmt.Errorf(
					"gagal migration seed user : %v",
					err.Error(),
				)
			}
		}
	}

	return nil
}

func isExisting(db *gorm.DB) bool {
	var users []user.User
	db.Find(&users)

	return len(users) > 0
}
