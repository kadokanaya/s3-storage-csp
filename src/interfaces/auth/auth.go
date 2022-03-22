package auth

import (
	"fmt"
	"sp/src/domains/entities"
	"sp/src/usecases/port"
	"time"

	"github.com/dgrijalva/jwt-go"
	// jwtmiddleware "github.com/auth0/go-jwt-middleware"
	// jwt "github.com/form3tech-oss/jwt-go"
)

type Auth struct {
	UserID string
	Iat    int64
}

const (
	// secret は openssl rand -base64 40 コマンドで作成した。
	secret = "2FMd5FNSqS/nW2wWJy5S3ppjSHhUnLt8HuwBkTD6HqfPfBBDlykwLA=="

	// userIDKey はユーザーの ID を表す。
	userIDKey = "user_id"

	// iat と exp は登録済みクレーム名。それぞれの意味は https://tools.ietf.org/html/rfc7519#section-4.1 を参照。{
	iatKey = "iat"
	expKey = "exp"
	// }

	// lifetime は jwt の発行から失効までの期間を表す。
	lifetime = 30 * time.Minute
)

type RequestAuth struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type authHandler struct {
}

func NewAuthHandler() port.Auth {
	return &authHandler{}
}

func (ua *authHandler) Login(user *entities.LoginUser) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"exp":     time.Now().Add(lifetime).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(secret))
}

func Parse(sign string) (*Auth, error) {
	token, err := jwt.Parse(sign, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(secret), nil
	})

	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, fmt.Errorf("%s is expired: %s", sign, err)
			} else {
				return nil, fmt.Errorf("%s is invalid: %s", sign, err)
			}
		} else {
			return nil, fmt.Errorf("%s is invalid : %s", sign, err)
		}
	}

	if token == nil {
		return nil, fmt.Errorf("not found token in %s", sign)
		//        return nil, err.Errorf("not found token in %s:", sign, err)
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return nil, fmt.Errorf("not found claims in %s", sign)
	}
	userID, ok := claims[userIDKey].(string)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", userIDKey, sign)
	}
	iat, ok := claims[iatKey].(float64)
	if !ok {
		return nil, fmt.Errorf("not found %s in %s", iatKey, sign)
	}

	return &Auth{
		UserID: userID,
		Iat:    int64(iat),
	}, nil
}

