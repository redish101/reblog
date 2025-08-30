package jwt

import (
	"crypto/ed25519"
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/redish101/reblog/internal/env"
)

// Claims JWT claims 结构
type Claims struct {
	Username string `json:"username"`
	Nickname string `json:"nickname"`
	Email    string `json:"email"`
	jwt.RegisteredClaims
}

// GenerateToken 生成 JWT token
func GenerateToken(username, nickname, email string) (string, error) {
	// 设置过期时间为24小时
	expirationTime := time.Now().Add(24 * time.Hour)

	// 创建 claims
	claims := &Claims{
		Username: username,
		Email:    email,
		Nickname: nickname,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    "reblog",
			Subject:   username,
		},
	}

	// 使用 ed25519 算法创建 token
	token := jwt.NewWithClaims(jwt.SigningMethodEdDSA, claims)

	// 使用私钥签名
	tokenString, err := token.SignedString(env.JWTPrivateKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

// ParseToken 解析 JWT token
func ParseToken(tokenString string) (*Claims, error) {
	// 解析 token
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// 验证签名算法
		if _, ok := token.Method.(*jwt.SigningMethodEd25519); !ok {
			return nil, errors.New("unexpected signing method")
		}
		// 返回公钥用于验证
		return ed25519.PublicKey(env.JWTPublicKey), nil
	})
	if err != nil {
		return nil, err
	}

	// 验证 token 有效性并提取 claims
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims, nil
	}

	return nil, errors.New("invalid token")
}

// ValidateToken 验证 token 是否有效
func ValidateToken(tokenString string) bool {
	_, err := ParseToken(tokenString)
	return err == nil
}
