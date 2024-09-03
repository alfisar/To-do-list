package service

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"strings"

	repoRedis "todolist/application/redis/repository"
	"todolist/application/user/repository"
	"todolist/config"
	"todolist/domain"
	"todolist/helper"
	"todolist/internal/consts"
	"todolist/internal/errorhandler"
	validations "todolist/internal/validation"

	"github.com/go-redis/redis/v8"
	"github.com/valyala/fasthttp"
	"gorm.io/gorm"
)

type authService struct {
	repo      repository.UserContract
	repoRedis repoRedis.RedisRepositoryContract
}

func NewAuthService(repo repository.UserContract, repoRedis repoRedis.RedisRepositoryContract) *authService {
	return &authService{
		repo:      repo,
		repoRedis: repoRedis,
	}
}

func (s *authService) Registration(ctx context.Context, r *fasthttp.Request, data domain.User) (result domain.User, err domain.ErrorData) {
	var (
		errChan      chan domain.ErrorData
		successImage bool
		dirImage     string
		sqlTrx       *gorm.DB
	)

	defer func() {
		if err.Code != 0 {
			if successImage {
				errData := helper.DeleteImage(dirImage)
				log.Println(errData)
			}

		}
	}()
	errData := helper.ValidationDataUser(data)
	if errData != nil {
		log.Printf("Error validation data user on func registration %s : %s", data.Email, errData.Error())
		err = errorhandler.ErrValidation(errData)
		return
	}

	where := map[string]any{
		"email": data.Email,
	}

	resultUser, errData := s.repo.Get(config.AppConfig.DBSql, where)
	fmt.Println(errData.Error())
	fmt.Println(gorm.ErrRecordNotFound.Error())
	if errData != nil && !strings.Contains(errData.Error(), gorm.ErrRecordNotFound.Error()) {
		log.Printf("Error get data user on func registration %s : %s", data.Email, errData.Error())
		err = errorhandler.ErrGetData(errData)
		return
	}

	if resultUser.ID != 0 {
		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidLogicBisnis, errorhandler.ErrMsgDataExist)
		return
	}

	data.Password, err = helper.GeneratePass(data.Password)
	if err.Code != 0 {
		log.Printf("Error hashing pass on func registration %s : %s", data.Email, err.Errors)
		return
	}

	sqlTrx = config.AppConfig.DBSql.Begin()

	errChan = make(chan domain.ErrorData, 2)
	multipart, _ := r.MultipartForm()

	photos, errData := validations.SaveImage(config.AppConfig, config.AppConfig.DIR.DirUser, multipart.File["image"][0])
	if errData != nil {
		log.Printf("Error save image on func registration %s : %s", data.Email, errData.Error())
		errChan <- errorhandler.ErrValidation(errData)
		return
	}
	data.Photo = photos
	successImage = true
	dirImage = config.AppConfig.DIR.DirUser + "/" + photos

	errData = s.repo.Create(sqlTrx, data)
	if errData != nil {
		sqlTrx.Rollback()
		log.Printf("Error create user on func registration %s : %s", data.Email, errData.Error())
		errChan <- errorhandler.ErrInsertData(errData)
		return
	}

	sqlTrx.Commit()
	result = data
	return
}

func (s *authService) Login(ctx context.Context, data domain.Login) (token string, err domain.ErrorData) {
	key := "LOGIN_" + data.Username

	result, errData := s.repoRedis.Get(ctx, config.AppConfig.DBRedis[consts.Token], key)
	if errData != nil && !strings.Contains(errData.Error(), string(redis.Nil)) {
		log.Printf("Error get data user redis on func login : %s", errData.Error())

		err = errorhandler.ErrGetData(errData)
		return
	} else if result != "" {

		errData = json.Unmarshal([]byte(result), &token)
		if errData != nil {
			log.Printf("Error parsing data on func login : %s", errData.Error())
			err = errorhandler.ErrHashing(errData)
			return
		}
		return
	}

	where := map[string]any{}
	if data.Email != "" {
		where["email"] = data.Email
	} else if data.Username != "" {
		where["username"] = data.Username
	} else {
		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeRequired, errorhandler.ErrMsgLoginRequired)
		return
	}

	resultUser, errData := s.repo.Get(config.AppConfig.DBSql, where)
	if errData != nil && !strings.Contains(errData.Error(), gorm.ErrRecordNotFound.Error()) {
		log.Printf("Error get data user on func login : %s", errData.Error())

		err = errorhandler.ErrGetData(errData)
		return
	} else if errData != nil && strings.Contains(errData.Error(), gorm.ErrRecordNotFound.Error()) {
		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidLogicBisnis, errorhandler.ErrMsgLoginFailed)
		return
	}

	errData = helper.Verify(resultUser.Password, data.Password)
	if errData != nil {
		log.Printf("Error verify password on func login : %s", errData.Error())

		err = errorhandler.ErrInvalidLogic(errorhandler.ErrCodeInvalidLogicBisnis, errorhandler.ErrMsgLoginFailed)
		return
	}

	token, errData = config.AppConfig.Jwt.GetToken(config.AppConfig.ExpData.ExpLogin, resultUser.ID)
	if errData != nil {
		log.Printf("Error generate token on func login : %s", errData.Error())

		err = errorhandler.ErrInternal(errorhandler.ErrCodeGenerateToken, errData)
		return
	}

	token, errData = helper.EncryptAES256CBC([]byte(token))
	if errData != nil {
		log.Printf("Error hashing aes 256 token on func login : %s", errData.Error())

		err = errorhandler.ErrHashing(errData)
		return
	}

	errData = s.repoRedis.Insert(ctx, config.AppConfig.DBRedis[consts.Token], key, "", config.AppConfig.ExpData.ExpLogin)

	if errData != nil {
		log.Printf("Error insert redis data token on func login : %s", errData.Error())

		err = errorhandler.ErrInsertData(errData)
		return
	}
	return
}
