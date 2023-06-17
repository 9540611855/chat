package auth

import (
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
)

func GetJWTKey() ([]byte, error) {

	return []byte("mHPqAeyMhQzR3RdVcN79"), nil
}

func GenerateJWT(userID int64) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 3).Unix()

	jwtKey, err := GetJWTKey()
	if err != nil {
		return "", err
	}

	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func ValidateJWT(tokenString string) (jwt.MapClaims, error) {
	jwtKey, err := GetJWTKey()
	if err != nil {
		return nil, err
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Unix() > int64(exp) {
				return nil, fmt.Errorf("jwt 过期 请重新登录")
			}
		}
		return claims, nil
	}

	return nil, fmt.Errorf("invalid token")
}

func CheckJwtTokenTime(tokenString string, userID int64) bool {
	jwtKey, err := GetJWTKey()
	if err != nil {
		return false
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return jwtKey, nil
	})

	if err != nil {
		return false
	}
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		user, _ := claims["user"].(float64)

		if int64(user) != userID {
			return false
		}
		if exp, ok := claims["exp"].(float64); ok {
			if time.Now().Add(-time.Hour*2).Unix() > int64(exp) {
				//jwt 超过2个小时就给验证失败
				return false
			}
		}
	}
	return true
}
