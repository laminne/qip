package controller

import (
	"errors"

	"github.com/approvers/qip/pkg/utils/password/argon2"

	"github.com/approvers/qip/pkg/application/user"
	"github.com/approvers/qip/pkg/repository"
	"github.com/approvers/qip/pkg/utils/password"
	"github.com/approvers/qip/pkg/utils/token"
)

type AuthController struct {
	repo            repository.UserRepository
	findUserService user.FindUserService
	passwordEncoder password.Encoder
	tokenGenerator  token.JWTTokenGenerator
	tokenParser     token.JWTTokenParser
}

func NewAuthController(r repository.UserRepository, key string) *AuthController {
	return &AuthController{
		repo:            r,
		findUserService: *user.NewFindUserService(r),
		passwordEncoder: argon2.NewArgon2PasswordEncoder(),
		tokenGenerator:  *token.NewJWTTokenGenerator(key),
		tokenParser:     *token.NewJWTTokenParser(key),
	}
}

func (a *AuthController) CheckToken(t string) bool {
	_, err := a.tokenParser.Parse(t)
	if err != nil {
		return false
	}
	return true
}

// Login ログイン
func (a *AuthController) Login(name string, pass string) (string, error) {
	// ユーザー情報を読み出す
	u, err := a.findUserService.FindByName(name)
	if err != nil {
		return "", err
	}
	var usr *user.UserData
	for _, v := range u {
		if v.IsLocalUser() {
			usr = &v
		}
	}
	if usr == nil {
		// 存在しないときはエラー
		return "", errors.New("user is not local-user")
	}

	// パスワードを検証する
	pw := *usr.Password()
	if !a.passwordEncoder.IsMatchPassword(pass, password.EncodedPassword(pw)) {
		// 検証失敗
		return "", errors.New("password is not matched")
	}

	// トークンを生成して返却する
	tk, err := a.tokenGenerator.NewToken(usr.Id())
	if err != nil {
		return "", err
	}

	return tk, nil
}
