package interfaces_test

import (
	"fmt"
	"testing"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func TestAuth(t *testing.T) {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": 12345678,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))
	t.Fatal("tokenString:", tokenString) // tokenString: eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJleHAiOjE2MzQwNTEzNjMsInVzZXJfaWQiOjEyMzQ1Njc4fQ.OooYrharapD5X2LV5UUWBOkEqH57wDfMd5ibkIpJHYM
}

func TestAuthVerify(t *testing.T) {
	// Claimsオブジェクトの作成
	claims := jwt.MapClaims{
		"user_id": 12345678,
		"exp":     time.Now().Add(time.Hour * 24).Unix(),
	}

	// ヘッダーとペイロードの生成
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// トークンに署名を付与
	tokenString, _ := token.SignedString([]byte("SECRET_KEY"))

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Don't forget to validate the alg is what you expect:
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte("SECRET_KEY"), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		fmt.Printf("user_id: %v\n", int64(claims["user_id"].(float64)))
		fmt.Printf("exp: %v\n", int64(claims["exp"].(float64)))
		t.Fatal(claims)
	} else {
		t.Fatal(err)
	}
}
