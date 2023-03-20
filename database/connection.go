package database

import (
	"fmt"
	"imp/helper"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func NewConnection(pathENV string) (*gorm.DB, error) {
	env, err := helper.GetENV(pathENV)
	if err != nil {
		return nil, err
	}

	dbUsername := env["DATABASE_USERNAME"]
	dbPassword := env["DATABASE_PASSWORD"]
	dbHost := env["DATABASE_HOST"]
	dbPort := env["DATABASE_PORT"]
	dbName := env["DATABASE_NAME"]

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v?charset=utf8mb4&parseTime=True&loc=Local",
		dbUsername,
		dbPassword,
		dbHost,
		dbPort,
		dbName,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf(
			"gagal koneksi ke database : %v",
			err.Error(),
		)
	}

	return db, nil
}