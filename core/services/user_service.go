package services

import (
	"encoding/json"
	"fmt"

	"hexagonal/common/cache"
	"hexagonal/common/logs"
	"hexagonal/core/models"
	"hexagonal/core/repositories"

	"hexagonal/utils"
	"strings"
	"time"

	"github.com/spf13/viper"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

type userService struct {
	cache    cache.AppCache
	log      logs.AppLog
	userRepo repositories.UserRepository
}

func NewUserService(log logs.AppLog, cache cache.AppCache, userRepo repositories.UserRepository) UserService {
	return userService{userRepo: userRepo, cache: cache, log: log}
}

func (s userService) GetUsers(p models.UserPaginationModel) (*models.UserResGetAllModel, error) {
	// TODO Query Redis
	data, err := s.cache.Get(fmt.Sprintf("user:*:page:%v:row:%v", p.Page, p.Row))
	if err != nil {
		if err.Error() != "cache: no documents in result" {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}
	}

	var toHandler models.UserResGetAllModel

	toHandler.Pagination.Page = p.Page
	toHandler.Pagination.Row = p.Row
	toHandler.Pagination.Total = 0

	usersResponses := []models.UserResModel{}
	var users []models.UserModel

	if data == nil {
		users, err = s.userRepo.GetAll(models.UserPaginationModel{
			Page: toHandler.Pagination.Page,
			Row:  toHandler.Pagination.Row,
		})
		if err != nil {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

		// # DTO Data Tranfer Object
		for _, user := range users {
			userRes := models.UserResModel{
				UserID: user.UserID,
				Email:  user.Email,

				Name: user.Name,
				Role: user.Role,
			}
			usersResponses = append(usersResponses, userRes)

		}
		// TODO Set Cache
		json, err := json.Marshal(usersResponses)
		if err != nil {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

		err = s.cache.Set(fmt.Sprintf("user:*:page:%v:row:%v", p.Page, p.Row), string(json))
		if err != nil {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

	} else {
		json.Unmarshal([]byte(*data), &usersResponses)
	}

	toHandler.Items = usersResponses

	countUser, err := s.userRepo.GetCountAll()
	if err != nil {
		s.log.Error(err)
		return nil, utils.HandlerError{
			Code:    500,
			Message: "unexpected error",
		}
	}

	toHandler.Pagination.Total = int(countUser)

	return &toHandler, nil
}

func (s userService) GetUser(userid string) (*models.UserResModel, error) {
	// TODO Query Redis
	data, err := s.cache.Get("user:" + userid)
	if err != nil {
		if err.Error() != "cache: no documents in result" {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}
	}
	var user *models.UserModel
	var usersResponses models.UserResModel

	if data == nil {
		user, err = s.userRepo.GetById(userid)
		if err != nil {
			if err == mongo.ErrNoDocuments {
				return nil, utils.HandlerError{
					Code:    200,
					Message: "user not found",
				}
			}
			// # Tech Error
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

		usersResponses = models.UserResModel{
			UserID: user.UserID,
			Email:  user.Email,
			// Password: user.Password,
			Name: user.Name,
			Role: user.Role,
			// Status:   user.Status,
		}

		// TODO Set Cache
		json, err := json.Marshal(usersResponses)
		if err != nil {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

		err = s.cache.Set("user:"+userid, string(json))
		if err != nil {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}
	} else {
		json.Unmarshal([]byte(*data), &usersResponses)
	}

	return &usersResponses, nil
}

func (s userService) SignIn(payload *models.SignInReqModel) (*models.SignInResModel, error) {
	user, err := s.userRepo.GetByEmail(strings.ToLower(payload.Username))
	if err != nil {
		if err == mongo.ErrNoDocuments {
			return nil, utils.HandlerError{
				Code:    401,
				Message: "username or password is incorrect",
			}
		}
		// # Tech Error
		s.log.Error(err)
		return nil, utils.HandlerError{
			Code:    500,
			Message: "unexpected error",
		}
	}

	// TODO Get Array is local Provider

	for _, v := range user.Oauth {
		if v.Provider == "local" {
			err = bcrypt.CompareHashAndPassword([]byte(v.Password), []byte(payload.Password))
			if err != nil {
				// # Tech Error
				s.log.Error(err)
				return nil, utils.HandlerError{
					Code:    500,
					Message: "unexpected error",
				}

			}
			private := viper.GetString("app.access_token_private_key")
			token, err := utils.CreateToken(30*time.Minute, user.UserID, user.Role, private)
			if err != nil {
				return nil, utils.HandlerError{
					Code:    500,
					Message: "unexpected error",
				}
			}

			//# DTO
			t := models.SignInResModel{
				Accesstoken: token,
				Status:      user.Status,
				Name:        user.Name,
				Email:       strings.ToLower(user.Email),
				Role:        user.Role,
			}

			// TODO 5.Return To handler
			return &t, nil
		}
	}
	return nil, utils.HandlerError{
		Code:    401,
		Message: "username or password is incorrect",
	}
}

func (s userService) SignUp(r *models.SignUpReqModel) (*models.SignUpResModel, error) {
	// TODO 2.Generate new password use bcryp
	newPass, err := bcrypt.GenerateFromPassword([]byte(r.Password), 10)
	if err != nil {
		return nil, utils.HandlerError{
			Code:    500,
			Message: "unexpected error",
		}
	}

	// TODO 3.make payload to repositiry
	data := models.UserCreateModel{
		Name:     r.Name,
		Email:    strings.ToLower(r.Email),
		Password: string(newPass),
		Status:   false,
		Role:     1,
		Provider: "local",
	}

	// TODO 4.insert to db
	newUser, err := s.userRepo.Create(data)
	// TODO 5.response
	if err != nil {
		if err.Error() == "email already exist" {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    400,
				Message: "email already exist",
			}
		} else {
			s.log.Error(err)
			return nil, utils.HandlerError{
				Code:    500,
				Message: "unexpected error",
			}
		}

	}

	// # Data Tranfer object
	u := models.SignUpResModel{
		Email:  newUser.Email,
		Name:   newUser.Name,
		Role:   newUser.Role,
		Status: newUser.Status,
	}

	return &u, nil
}

func (s userService) UpdateUser(models.UserUpdateReqModel) error {
	return nil
}
