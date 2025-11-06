package token

import (
	"errors"
	"fmt"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

// Config 包括 Token 包的配置选项
type Config struct {
	key         string        // 用于签发和解析 Token 的密钥
	identityKey string        // Token 中用户身份的键，项目中可以用 userID 来
	expiration  time.Duration // 签发 Token 的过期时间
}

var (
	config = Config{
		key:         "Rtg8BPKNEf2mB4mgvKONGPZZQSaJWNLijxR42qRgq0iBb5",
		identityKey: "identityKey",
		expiration:  2 * time.Hour,
	}
	once sync.Once // 确保配置只被初始化一次
)

// Init 初始化 JWT 配置
func Init(key string, identityKey string, expiration time.Duration) {
	once.Do(func() {
		if key != "" {
			config.key = key // 设置密钥
		}
		if identityKey != "" {
			config.identityKey = identityKey // 设置身份键
		}
		if expiration != 0 {
			config.expiration = expiration
		}
	})
}

// Parse 解析密钥
// tokenStr 需要解析的 token
// key token 的密钥
func Parse(tokenStr, key string) (string, error) {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		// 确保传入的 Token 的加密算法是预期的加密算法，若不是则无法解析
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, jwt.ErrSignatureInvalid
		}
		return []byte(key), nil // 返回密钥
	})
	if err != nil {
		return "", err
	}
	var identityKey string
	// 如果解析成功，从 Token 中提取 identity
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		if key, exists := claims[config.identityKey]; exists {
			if identity, valid := key.(string); valid {
				// 获取身份键
				identityKey = identity
			}
		}
	}
	if identityKey == "" {
		return "", jwt.ErrSignatureInvalid
	}
	return identityKey, nil
}

// ParseRequest 从请求头中获取令牌，并将其传递给 Parse 函数以进行解析
func ParseRequest(c *gin.Context) (string, error) {
	h := c.Request.Header.Get("Authorization")
	if len(h) == 0 || h == "" {
		return "", errors.New("authorization header is empty")
	}
	var token string
	fmt.Sscanf(h, "Bearer %s", &token)
	return Parse(token, config.key)
}

// Sign 使用 jwtSecret 签发 token，token 的 claims 中会存放传入的 subject
func Sign(identityKey string) (string, time.Time, error) {
	expireAt := time.Now().Add(config.expiration)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		config.identityKey: identityKey,       // 存放用户身份
		"nbf":              time.Now().Unix(), // 生效时间
		"iat":              time.Now().Unix(), // 签发时间
		"exp":              expireAt.Unix(),   // 过期时间
	})
	if config.key == "" {
		return "", time.Time{}, jwt.ErrInvalidKey
	}
	tokenString, err := token.SignedString([]byte(config.key))
	if err != nil {
		return "", time.Time{}, err
	}
	return tokenString, expireAt, nil
}
