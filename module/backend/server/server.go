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
	verifyUserTokenUsecase := user.NewUserTokenVerificationUsecase(cfg.AuthConfig, db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/user/register", middleware.WithDefaultNoAuthMiddlewares(handler.NewUserRegistrationHandler(registerUserUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/user/login", middleware.WithDefaultNoAuthMiddlewares(handler.NewUserLoginHandler(loginUserUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/user/me", middleware.WithDefaultMiddlewares(handler.NewUserSelfInfoHandler(), verifyUserTokenUsecase).ServeHTTP).Methods("GET")

	return router, nil
}
