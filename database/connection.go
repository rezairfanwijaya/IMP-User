package database

import (
	"fmt"
	"imp/helper"
	"imp/user"

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

	// connection to database
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf(
			"gagal koneksi ke database : %v",
			err.Error(),
		)
	}

	// migration schema
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return db, fmt.Errorf(
			"gagal migration schema : %v",
			err.Error(),
		)
	}

	// migaration seed users
	if err := MigarationUserSeed(db); err != nil {
		return db, err
	}
	return db, nil
}
