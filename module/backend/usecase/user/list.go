package user

import (
	"context"
	"github.com/ahmadlubis/lavandeapp/module/backend/entity"
	"github.com/ahmadlubis/lavandeapp/module/backend/model"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/request"
	"github.com/ahmadlubis/lavandeapp/module/backend/model/response"
	"github.com/ahmadlubis/lavandeapp/module/backend/usecase"
	"gorm.io/gorm"
	"strings"
)

const (
	minLimit = 10
	maxLimit = 50
)

type listUserRequest struct {
	Name     string
	NIK      string
	Email    string
	PhoneNo  string
	Status   entity.UserStatus
	Religion entity.UserReligion
	Limit    int
	Offset   int
}

type userListUsecase struct {
	db *gorm.DB
}

func NewUserListUsecase(db *gorm.DB) usecase.UserListUsecase {
	return &userListUsecase{db: db}
}

func (u *userListUsecase) List(ctx context.Context, request request.ListUserRequest) (response.ListUserResponse, error) {
	var users []entity.User
	var total int64

	req, err := u.normalizeListRequest(request)
	if err != nil {
		return response.ListUserResponse{}, err
	}

	// Build WHERE query and its parameters
	var conditions []string
	var params []interface{}
	if req.Name != "" {
		conditions = append(conditions, "name LIKE ?")
		params = append(params, req.Name+"%")
	}
	if req.NIK != "" {
		conditions = append(conditions, "nik = ?")
		params = append(params, req.NIK)
	}
	if req.Email != "" {
		conditions = append(conditions, "email = ?")
		params = append(params, req.Email)
	}
	if req.PhoneNo != "" {
		conditions = append(conditions, "phone_no = ?")
		params = append(params, req.PhoneNo)
	}
	if req.Status != 0 {
		conditions = append(conditions, "status = ?")
		params = append(params, req.Status)
	}
	if req.Religion != 0 {
		conditions = append(conditions, "religion = ?")
		params = append(params, req.Religion)
	}
	query := strings.Join(conditions, " AND ")

	// Fetch the list of users considering pagination
	if result := u.db.Where(query, params...).Offset(req.Offset).Limit(req.Limit).Order("id").Find(&users); result.Error != nil {
		return response.ListUserResponse{}, model.NewUnknownError("", result.Error)
	}

	// Fetch the total of all users with NO regards to the pagination
	if result := u.db.Model(&users).Where(query, params...).Count(&total); result.Error != nil {
		return response.ListUserResponse{}, model.NewUnknownError("", result.Error)
	}

	return response.ListUserResponse{
		Data: users,
		Meta: response.PaginationMeta{
			Limit:  uint64(req.Limit),
			Offset: uint64(req.Offset),
			Count:  uint64(len(users)),
			Total:  uint64(total),
		},
	}, nil
}

func (u *userListUsecase) normalizeListRequest(req request.ListUserRequest) (listUserRequest, error) {
	var err error
	var status entity.UserStatus
	var religion entity.UserReligion

	if req.Limit < minLimit {
		req.Limit = minLimit
	}
	if req.Limit > maxLimit {
		req.Limit = maxLimit
	}
	if req.Offset < 0 {
		req.Offset = 0
	}

	if req.Status != "" {
		status, err = entity.ParseUserStatus(req.Status)
		if err != nil {
			return listUserRequest{}, err
		}
	}

	if req.Religion != "" {
		religion, err = entity.ParseUserReligion(req.Religion)
		if err != nil {
			return listUserRequest{}, err
		}
	}

	return listUserRequest{
		Name:     req.Name,
		NIK:      req.NIK,
		Email:    req.Email,
		PhoneNo:  req.PhoneNo,
		Status:   status,
		Religion: religion,
		Limit:    int(req.Limit),
		Offset:   int(req.Offset),
	}, nil
}
