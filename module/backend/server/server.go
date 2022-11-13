package server

import (
	"github.com/ahmadlubis/lavandeapp/config"
	adminHandler "github.com/ahmadlubis/lavandeapp/module/backend/handler/admin"
	"github.com/ahmadlubis/lavandeapp/module/backend/handler/middleware"
	tenantHandler "github.com/ahmadlubis/lavandeapp/module/backend/handler/tenant"
	unitHandler "github.com/ahmadlubis/lavandeapp/module/backend/handler/unit"
	userHandler "github.com/ahmadlubis/lavandeapp/module/backend/handler/user"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase/tenant"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase/unit"
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
	selfUpdateUserUsecase := user.NewUserUpdateUsecase(db)
	listUseUsecase := user.NewUserListUsecase(db)

	createUnitUsecase := unit.NewUnitCreationUsecase(db)
	updateUnitUsecase := unit.NewUnitUpdateUsecase(db)
	verifyOwnerUsecase := unit.NewUnitOwnerVerificationUsecase(db)

	createTenantUsecase := tenant.NewTenantCreationUsecase(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/user/register", middleware.WithDefaultNoAuthMiddlewares(userHandler.NewUserRegistrationHandler(registerUserUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/user/login", middleware.WithDefaultNoAuthMiddlewares(userHandler.NewUserLoginHandler(loginUserUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/user/me", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, userHandler.NewUserSelfInfoHandler()).ServeHTTP).Methods("GET")
	router.HandleFunc("/v1/user/me", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, userHandler.NewUserUpdateHandler(selfUpdateUserUsecase)).ServeHTTP).Methods("PATCH")

	router.HandleFunc("/v1/unit", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, unitHandler.NewUnitUpdateHandler(verifyOwnerUsecase, updateUnitUsecase)).ServeHTTP).Methods("PATCH")
	router.HandleFunc("/v1/unit/tenant", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, tenantHandler.NewTenantCreationHandler(verifyOwnerUsecase, createTenantUsecase)).ServeHTTP).Methods("POST")

	router.HandleFunc("/v1/admin/users", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUserListHandler(listUseUsecase)).ServeHTTP).Methods("GET")
	router.HandleFunc("/v1/admin/units", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUnitCreationHandler(createUnitUsecase)).ServeHTTP).Methods("POST")
	router.HandleFunc("/v1/admin/tenants", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewTenantCreationHandler(createTenantUsecase)).ServeHTTP).Methods("POST")

	return router, nil
}
