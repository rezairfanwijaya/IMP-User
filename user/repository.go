package user

import "gorm.io/gorm"

type IRepository interface {
	Save(user User) (User, error)
	FindByID(ID int) (User, error)
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

func (r *repository) FindByUsername(username string) (User, error) {
	var user User
	if err := r.db.Where("username = ?", username).Find(&user).Error; err != nil {
		return user, err
	}

	return user, nil
}
