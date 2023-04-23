package use_case

import (
	"capstone/constant"
	"capstone/data_type"
	"capstone/delivery/dto_request"
	"capstone/delivery/dto_response"
	jwtInternal "capstone/internal/jwt"
	"capstone/model"
	"capstone/repository"
	"capstone/util"
	"context"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type AuthUseCase interface {
	LoginEmail(ctx context.Context, request dto_request.AuthEmailLoginRequest) model.Token
	RegisterEmail(ctx context.Context, request dto_request.AuthEmailRegisterRequest) model.Token

	Parse(ctx context.Context, token string) (*model.User, error)
}

type authUseCase struct {
	userAccessTokenRepository repository.UserAccessTokenRepository
	userRepository            repository.UserRepository

	jwt jwtInternal.Jwt
}

func NewAuthUseCase(
	userAccessTokenRepository repository.UserAccessTokenRepository,
	userRepository repository.UserRepository,
	jwt jwtInternal.Jwt,
) AuthUseCase {
	return &authUseCase{
		userAccessTokenRepository: userAccessTokenRepository,
		userRepository:            userRepository,

		jwt: jwt,
	}
}

func (u *authUseCase) generateJwt(ctx context.Context, userId string) (*jwtInternal.Token, error) {
	var (
		now                   = util.CurrentDateTime()
		expiredAt             = now.Add(time.Hour * 24)
		maxGenerationAttempts = 10
	)

	userAccessToken := &model.UserAccessToken{
		Id:           util.NewUuid(),
		UserId:       userId,
		Revoked:      false,
		ExpiredAt:    expiredAt,
		IpAddress:    "",
		Longitude:    0,
		Latitude:     0,
		LocationName: "",
		Timestamp:    model.Timestamp{},
	}

	for maxGenerationAttempts > 0 {
		maxGenerationAttempts--

		userAccessToken.Id = util.NewUuid()

		if err := u.userAccessTokenRepository.Insert(ctx, userAccessToken); err != nil && maxGenerationAttempts == 0 {
			log.Println(err)
			return nil, fmt.Errorf("max access token generation attemps exceeded")
		} else if err == nil {
			break
		}
	}

	accessToken, err := u.jwt.Generate(jwtInternal.Payload{
		Id:        userAccessToken.Id,
		UserId:    userAccessToken.UserId,
		CreatedAt: userAccessToken.CreatedAt.DateTime().Time(),
		ExpiredAt: userAccessToken.ExpiredAt.Time(),
	})
	if err != nil {
		return nil, err
	}

	return accessToken, nil
}

func (u *authUseCase) mustGetHashedPassword(originalPassword string) string {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(originalPassword), bcrypt.DefaultCost)
	panicIfErr(err)
	return string(hashedPassword)
}

func (u *authUseCase) mustValidateComparePassword(hashedPassword string, originalPassword string) {
	panicIfErr(
		bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(originalPassword)),
	)
}

func (u *authUseCase) LoginEmail(ctx context.Context, request dto_request.AuthEmailLoginRequest) model.Token {
	user, err := u.userRepository.GetByEmail(ctx, request.Email)
	if err != nil {
		if err == constant.ErrNoData {
			panic(dto_response.NewBadRequestResponse("Email not registered"))
		}
		panic(err)
	}

	u.mustValidateComparePassword(user.Password, request.Password)

	accessToken, err := u.generateJwt(ctx, user.Id)
	panicIfErr(err)

	return model.Token{
		AccessToken:          accessToken.AccessToken,
		AccessTokenExpiredAt: data_type.NewDateTime(accessToken.ExpiredAt),
		TokenType:            accessToken.Type,
	}
}

func (u *authUseCase) RegisterEmail(ctx context.Context, request dto_request.AuthEmailRegisterRequest) model.Token {
	checkUser, err := u.userRepository.GetByEmail(ctx, request.Email)
	if err != nil && err != constant.ErrNoData {
		panic(err)
	} else if checkUser != nil {
		panic(dto_response.NewBadRequestResponse("Email already registered"))
	}

	user := model.User{
		Id:       util.NewUuid(),
		Name:     request.Name,
		Email:    util.StandarizeEmail(request.Email),
		Password: u.mustGetHashedPassword(request.Password),
	}

	panicIfErr(
		u.userRepository.Insert(ctx, &user),
	)

	accessToken, err := u.generateJwt(ctx, user.Id)
	panicIfErr(err)

	return model.Token{
		AccessToken:          accessToken.AccessToken,
		AccessTokenExpiredAt: data_type.NewDateTime(accessToken.ExpiredAt),
		TokenType:            accessToken.Type,
	}
}

func (u *authUseCase) Parse(ctx context.Context, token string) (*model.User, error) {
	payload, err := u.jwt.Parse(token)
	if err != nil {
		return nil, constant.ErrNotAuthenticated
	}

	var (
		tokenId = payload.Id
		userId  = payload.UserId
	)

	isExist, err := u.userAccessTokenRepository.IsExist(ctx, tokenId)
	if err != nil {
		return nil, err
	}

	if isExist {
		user, err := u.userRepository.Get(ctx, userId)
		if err != nil && err != constant.ErrNoData {
			return nil, err
		}

		if user != nil {
			return user, nil
		}
	}

	return nil, constant.ErrNotAuthenticated
}
