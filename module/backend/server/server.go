package server

import (
	"net/http"

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
)

func NewBackendServer(cfg *config.Config) (http.Handler, error) {
	db, err := gorm.Open(mysql.Open(cfg.MysqlConfig.BuildDsn()), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	registerUserUsecase := user.NewUserRegistrationUsecase(db)
	loginUserUsecase := user.NewUserLoginUsecase(cfg.AuthConfig, db)
	verifyUserTokenUsecase := user.NewUserTokenVerificationUsecase(cfg.AuthConfig, db)
	updateUserUsecase := user.NewUserUpdateUsecase(db)
	listUseUsecase := user.NewUserListUsecase(db)
	superadminUC := user.NewSuperAdminUsecase(db)

	createUnitUsecase := unit.NewUnitCreationUsecase(db)
	updateUnitUsecase := unit.NewUnitUpdateUsecase(db)
	verifyOwnerUsecase := unit.NewUnitOwnerVerificationUsecase(db)
	listUnitUsecase := unit.NewUnitListUsecase(db)

	createTenantUsecase := tenant.NewTenantCreationUsecase(db)
	listTenantUsecase := tenant.NewTenantListUsecase(db)

	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/v1/user/register", middleware.WithDefaultNoAuthMiddlewares(userHandler.NewUserRegistrationHandler(registerUserUsecase)).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/v1/user/login", middleware.WithDefaultNoAuthMiddlewares(userHandler.NewUserLoginHandler(loginUserUsecase)).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/v1/user/me", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, userHandler.NewUserSelfInfoHandler()).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/v1/user/me", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, userHandler.NewUserUpdateHandler(updateUserUsecase)).ServeHTTP).Methods(http.MethodPatch)

	router.HandleFunc("/v1/unit", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, unitHandler.NewUnitUpdateHandler(verifyOwnerUsecase, updateUnitUsecase)).ServeHTTP).Methods(http.MethodPatch)
	router.HandleFunc("/v1/unit", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, unitHandler.NewUnitListHandler(listUnitUsecase)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/v1/unit/tenant", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, tenantHandler.NewTenantCreationHandler(verifyOwnerUsecase, createTenantUsecase)).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/v1/unit/tenant", middleware.WithDefaultMiddlewares(verifyUserTokenUsecase, tenantHandler.NewTenantListHandler(verifyOwnerUsecase, listTenantUsecase)).ServeHTTP).Methods(http.MethodGet)

	router.HandleFunc("/v1/admin/users", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUserListHandler(listUseUsecase)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/v1/admin/users", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUserUpdateHandler(updateUserUsecase)).ServeHTTP).Methods(http.MethodPatch)
	router.HandleFunc("/v1/admin/units", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUnitCreationHandler(createUnitUsecase)).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/v1/admin/units", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewUnitListHandler(listUnitUsecase)).ServeHTTP).Methods(http.MethodGet)
	router.HandleFunc("/v1/admin/tenants", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewTenantCreationHandler(createTenantUsecase)).ServeHTTP).Methods(http.MethodPost)
	router.HandleFunc("/v1/admin/tenants", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewTenantListHandler(listTenantUsecase)).ServeHTTP).Methods(http.MethodGet)

	// API to set user as SuperAdmin
	router.HandleFunc("/v1/admin/set", middleware.WithDefaultAdminMiddlewares(verifyUserTokenUsecase, adminHandler.NewSuperAdminHandler(superadminUC)).ServeHTTP).Methods(http.MethodPut)

	return router, nil
}
