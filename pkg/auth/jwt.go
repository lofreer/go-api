package auth

import (
	"fmt"
	"go-api/api/app/models/admin"
	"go-api/api/pkg/database"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

// 定义授权保存信息
type CustomClaims struct {
	Id      uint64
	ExpTime int64
	jwt.StandardClaims
}

// 私钥
const (
	SECRETARY = "lofreer-secret-key"
)

// 获取用户token值
func GetToken(data *admin.Admin) (map[string]interface{}, error) {
	// 7200秒过期
	maxAge := 7200
	expTime := time.Now().Add(time.Duration(maxAge) * time.Second).Unix()
	customClaims := &CustomClaims{
		Id:      data.Id,
		ExpTime: expTime,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expTime,   // 过期时间，必须设置
			Issuer:    "lofreer", // 非必须，也可以填充用户名，
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, customClaims)
	tokenString, err := token.SignedString([]byte(SECRETARY))
	if err != nil {
		return nil, err
	}
	rlt := make(map[string]interface{})
	rlt["expTime"] = expTime
	rlt["token"] = tokenString
	return rlt, nil
}

// 使用token换取user信息
func GetUser(tokenString string) (interface{}, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(SECRETARY), nil
	})
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		id := int(claims["Id"].(float64))
		admin := admin.Admin{}
		database.DB.First(&admin, id)
		return admin, nil
	} else {
		return nil, err
	}
}
