package server

import (
	"github.com/ahmadlubis/lavandeapp/config"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler/middleware"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase/user"
	"github.com/gorilla/mux"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

func NewBackendServer(cfg *config.Config) (http.Handler, error) {
	db, err := gorm.Open(mysql.Open(cfg.MysqlConfig.BuildDsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	registerUserUsecase := user.NewUserRegistrationUsecase(db)
	loginUserUsecase := user.NewUserLoginUsecase(cfg.AuthConfig, db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/user/register", middleware.WithDefaultMiddlewares(handler.NewUserRegistrationHandler(registerUserUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/user/login", middleware.WithDefaultMiddlewares(handler.NewUserLoginHandler(loginUserUsecase)).ServeHTTP).Methods("POST")

	return router, nil
}
