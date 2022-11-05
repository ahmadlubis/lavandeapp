package main

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase/user"
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

	registrar := user.NewUserRegistrationUsecase(db)
	usr, err := registrar.Register(context.Background(), request.RegisterUserRequest{
		Name:     "Test user01",
		Email:    "test.email@gmial.com",
		Password: "password",
	})
	if err != nil {
		panic(err)
	}

	db.Where("email = ?", usr.Email).First(&usr)
	err = bcrypt.CompareHashAndPassword(usr.Password, []byte("password"))
	if err != nil {
		panic(err)
	}

	err = bcrypt.CompareHashAndPassword(usr.Password, []byte("password2"))
	if err != nil {
		panic(err)
	}
}
