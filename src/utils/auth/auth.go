package auth

import (
	"errors"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
)

const (
	userIDKey = "user_id"
	// iat と exp は登録済みクレーム名。それぞれの意味は https://tools.ietf.org/html/rfc7519#section-4.1 を参照
	iatKey = "iat"
	expKey = "exp"
	// lifetime は jwt の発行から失効までの期間を表す。
	lifetime = 72 * time.Hour
)

func NewJWT(userID int64) (string, error) {
	fmt.Println("create jwt userID: ", userID)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		userIDKey: userID,
		iatKey:    time.Now().Unix(),
		expKey:    time.Now().Add(lifetime).Unix(),
	})

	return token.SignedString([]byte(os.Getenv("SIGNINGKEY")))
}

func DecodeJWT(token string) (int64, error) {
	decodedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return "", fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SIGNINGKEY")), nil
	})

	if err != nil {
		fmt.Println("token parse error: ", err)
		return 0, err
	}

	claims, ok := decodedToken.Claims.(jwt.MapClaims)
	if !ok {
		fmt.Println("jwt.MapClaims type cast error")
		return 0, err
	}
	// reflectで確認したらfloat64になっていて直でint64にキャストできないため一旦float64
	userID, ok := claims["user_id"].(float64)
	if !ok {
		fmt.Println("type cast error. userId to float64")
		return 0, errors.New("type cast error. userId to float64")
	}

	return int64(userID), nil
}

func GetTokenFromHeader(bearer string) string {
	// bearer := r.Header.Get("Authorization")
	return strings.Replace(bearer, "Bearer ", "", 1)
}
