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

	registrar := user.NewUserRegistrationUsecase(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/user/register", middleware.WithDefaultMiddlewares(handler.NewUserRegistrationHandler(registrar)).ServeHTTP).Methods("POST")

	return router, nil
}
