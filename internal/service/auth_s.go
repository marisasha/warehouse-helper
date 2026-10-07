package service

import (
	"crypto/rand"
	"crypto/sha1"
	"encoding/hex"
	"errors"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/marisasha/warehouse-helper/internal/dto"
	"github.com/marisasha/warehouse-helper/internal/repository"
)

// публичные структуры
type AuthService struct {
	repos repository.Authorization
}

// приватные структуры
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

// константы
const (
	salt       = "vfzgz25f2sdf4gsf.fsg246ydhd.gh3ilof10"
	signingKey = "fnhj52..254nfslmnl8hfsvbnjs.2fjisg"
	tokenTTL   = 12 * time.Hour
	queueName  = "email_verification"
)

func NewAuthService(repos repository.Authorization) *AuthService {
	return &AuthService{
		repos: repos,
	}
}

func (s *AuthService) CreateUser(user *dto.User) error {
	user.Password = *generatePasswordHash(user.Password)
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password *string) (string, error) {
	user, err := s.repos.GetUser(username, generatePasswordHash(*password))
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		user.ID,
	})

	return token.SignedString([]byte(signingKey))
}

func (s *AuthService) ParseToken(accesToken *string) (int, error) {
	token, err := jwt.ParseWithClaims(*accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(signingKey), nil
	})

	if err != nil {
		return 0, err
	}

	claims, ok := token.Claims.(*tokenClaims)
	if !ok {
		return 0, errors.New("token claims are not of type *tokenClaims")
	}
	return claims.UserId, nil
}

func generatePasswordHash(password string) *string {
	hash := sha1.New()
	hash.Write([]byte(password))
	passwordHash := fmt.Sprintf("%x", hash.Sum([]byte(salt)))
	return &passwordHash
}

func generateToken() string {
	b := make([]byte, 16)
	rand.Read(b)
	return hex.EncodeToString(b)
}
