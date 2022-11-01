package main

import (
	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	cfg := config.NewConfig()
	db, err := gorm.Open(mysql.Open(cfg.MysqlConfig.BuildDsn()), &gorm.Config{})
	if err != nil {
		panic(err)
	}

	pass, err := bcrypt.GenerateFromPassword([]byte("password"), bcrypt.DefaultCost)
	if err != nil {
		panic(err)
	}

	user := entity.User{
		Name:            "Test User",
		Email:           "test@email.com",
		Role:            entity.UserRoleResident,
		ResidenceStatus: entity.UserResidenceStatusRenter,
		Password:        pass,
	}

	result := db.Create(&user)
	if result.Error != nil {
		panic(err)
	}

	db.Where("email = ?", "test@email.com").First(&user)
	err = bcrypt.CompareHashAndPassword(user.Password, []byte("password"))
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(user.Password, []byte("password2"))
	if err != nil {
		panic(err)
	}
}
