package auth_jwt

import (
	"errors"
	"github.com/golang-jwt/jwt/v5"
	"net/http"

	"github.com/gin-gonic/gin"
)

var (
	Secret = []byte("wzhnb")
	// TokenExpireDuration = time.Hour * 2 过期时间
)

type JWTClaims struct {
	UserId      int    `json:"user_id"`
	Username    string `json:"user_name"`
	PhoneNumber string `json:"phone_number"`
	Avatar      string `json:"avatar"`
	Identity    string `json:"identity"`
	IpPosition  string `json:"ip_position"`
	jwt.RegisteredClaims
}

// 生成token
func GenToken(userid int, userName string, phoneNumber string, avatar string, identity string, ipPosition string) (string, error) {
	claims := JWTClaims{
		UserId:      userid,
		Username:    userName,
		PhoneNumber: phoneNumber,
		Avatar:      avatar,
		Identity:    identity,
		IpPosition:  ipPosition,
		RegisteredClaims: jwt.RegisteredClaims{
			Issuer: "WZH",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(Secret)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ParseToken 解析token
func ParseToken(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return Secret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(*JWTClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// VerifyToken 验证token
func VerifyToken(tokenString string) (int, error) {
	if tokenString == "" {
		return 0, nil
	}
	claims, err := ParseToken(tokenString)
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}

// AuthMiddleware 生成权限中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		tokenString := c.PostForm("token")
		if tokenString == "" {
			tokenString = c.Query("token")
		}

		userId, err := VerifyToken(tokenString)
		if err != nil || userId == 0 {
			c.Abort()
			c.JSON(http.StatusInternalServerError, "token验证失败")
			return
		}

		c.Set("UserId", userId)
		c.Next()
	}
}
