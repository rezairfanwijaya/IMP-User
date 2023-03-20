package user

import (
	"math"

	"gorm.io/gorm"
)

type IRepository interface {
	Save(user User) (User, error)
	FindByID(ID int) (User, error)
	FindAll(params ParamsGetAllUsers, offset int) ([]User, int, int, error)
	FindByUsername(username string) (User, error)
}

type repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (r *repository) Save(user User) (User, error) {
	if err := r.db.Create(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindByID(ID int) (User, error) {
	var user User
	if err := r.db.Where("id = ?", ID).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}

func (r *repository) FindAll(params ParamsGetAllUsers, offset int) ([]User, int, int, error) {
	var users []User
	var totalData int64 = 0
	var totalPage int = 0

	if err := r.db.Order(params.Order).Limit(params.Limit).Offset(offset).Find(&users).Error; err != nil {
		return users, int(totalData), totalPage, err
	}

	// total data
	if err := r.db.Model(&User{}).Count(&totalData).Error; err != nil {
		return users, int(totalData), totalPage, err
	}

	// total page
	totalPage = int(math.Ceil(float64(totalData)/float64(params.Limit))) - 1

	return users, int(totalData), totalPage, nil
}

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
