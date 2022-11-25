package unit

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"github.com/ahmadlubis/lavandeapp/module/backend/utility"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
	"net/http"
)

type unitCreationUsecase struct {
	db *gorm.DB
}

func NewUnitCreationUsecase(db *gorm.DB) usecase.UnitCreationUsecase {
	return &unitCreationUsecase{db: db}
}

func (u *unitCreationUsecase) Create(_ context.Context, req request.CreateUnitRequest) (entity.Unit, error) {
	unit := entity.Unit{
		GovID:  req.GovID,
		Tower:  req.Tower,
		Floor:  req.Floor,
		UnitNo: req.UnitNo,
	}
	if err := unit.Validate(); err != nil {
		return entity.Unit{}, err
	}

	if result := u.db.Create(&unit); result.Error != nil {
		if mysqlErr, ok := result.Error.(*mysql.MySQLError); ok {
			if mysqlErr.Number == utility.MysqlErrorConflictNumber {
				return entity.Unit{}, model.NewExpectedError("another account with the same GovID / (Tower, Floor, UnitNo) already exists", "UNIT_CONFLICT", http.StatusConflict, req.GovID)
			}
		}
		return entity.Unit{}, model.NewUnknownError(req.GovID, result.Error)
	}

	return unit, nil
}
