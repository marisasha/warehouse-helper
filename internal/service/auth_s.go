package service

import (
	"crypto/rand"
	"crypto/subtle"
	"encoding/base64"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	dto "github.com/marisasha/warehouse-helper/internal/dto/request"
	"github.com/marisasha/warehouse-helper/internal/repository"
	"golang.org/x/crypto/argon2"
)

// публичные структуры
type AuthService struct {
	repos repository.Authorization
	cfg   Argon2id
}

type Argon2id struct {
	SigningKey       string `mapstructure:"signing_key"`
	TokenTTL         int    `mapstructure:"token_ttl"`
	ArgonMemory      uint32 `mapstructure:"argon_memory"`
	ArgonIterations  uint32 `mapstructure:"argon_iterations"`
	ArgonParallelism uint8  `mapstructure:"argon_parallelism"`
	ArgonSaltLength  uint32 `mapstructure:"argon_salt_length"`
	ArgonKeyLength   uint32 `mapstructure:"argon_key_length"`
}

// приватные структуры
type tokenClaims struct {
	jwt.StandardClaims
	UserId int `json:"user_id"`
}

func NewAuthService(repos repository.Authorization, cfg Argon2id) *AuthService {
	return &AuthService{
		repos: repos,
		cfg:   cfg,
	}
}

func (s *AuthService) CreateUser(user *dto.User) error {
	hashedPassword, err := s.generateArgon2Hash(user.Password)
	if err != nil {
		return err
	}
	user.Password = hashedPassword
	return s.repos.CreateUser(user)
}

func (s *AuthService) GenerateToken(username, password string) (string, error) {
	userId, userPasswordHash, err := s.repos.GetUser(username)
	if err != nil {
		return "", err
	}
	match, err := s.comparePasswordAndHash(password, userPasswordHash)
	if err != nil || !match {
		return "", errors.New("invalid username or password")
	}
	tokenTTL := time.Duration(s.cfg.TokenTTL) * time.Second
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(tokenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		},
		userId,
	})

	return token.SignedString([]byte(s.cfg.SigningKey))
}

func (s *AuthService) ParseToken(accesToken *string) (int, error) {
	token, err := jwt.ParseWithClaims(*accesToken, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("invalid signing method")
		}
		return []byte(s.cfg.SigningKey), nil
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

func (s *AuthService) generateArgon2Hash(password string) (string, error) {
	salt := make([]byte, s.cfg.ArgonSaltLength)
	if _, err := rand.Read(salt); err != nil {
		return "", err
	}

	hash := argon2.IDKey([]byte(password), salt, s.cfg.ArgonIterations, s.cfg.ArgonMemory, s.cfg.ArgonParallelism, s.cfg.ArgonKeyLength)

	b64Salt := base64.RawStdEncoding.EncodeToString(salt)
	b64Hash := base64.RawStdEncoding.EncodeToString(hash)

	fullHash := fmt.Sprintf("$argon2id$v=%d$m=%d,t=%d,p=%d$%s$%s",
		argon2.Version, s.cfg.ArgonMemory, s.cfg.ArgonIterations, s.cfg.ArgonParallelism, b64Salt, b64Hash)

	return fullHash, nil
}

func (s *AuthService) comparePasswordAndHash(password, encodedHash string) (bool, error) {
	parts := strings.Split(encodedHash, "$")
	if len(parts) != 6 {
		return false, errors.New("invalid hash format")
	}

	if parts[1] != "argon2id" {
		return false, errors.New("incompatible argon2 variant")
	}

	var version int
	_, err := fmt.Sscanf(parts[2], "v=%d", &version)
	if err != nil || version != argon2.Version {
		return false, errors.New("incompatible argon2 version")
	}

	var memory, iterations uint32
	var parallelism uint8
	_, err = fmt.Sscanf(parts[3], "m=%d,t=%d,p=%d", &memory, &iterations, &parallelism)
	if err != nil {
		return false, errors.New("invalid argon2 parameters")
	}

	salt, err := base64.RawStdEncoding.DecodeString(parts[4])
	if err != nil {
		return false, err
	}

	expectedHash, err := base64.RawStdEncoding.DecodeString(parts[5])
	if err != nil {
		return false, err
	}

	actualHash := argon2.IDKey([]byte(password), salt, iterations, memory, parallelism, uint32(len(expectedHash)))

	// Защита от атак по времени (timing attacks)
	if subtle.ConstantTimeCompare(expectedHash, actualHash) == 1 {
		return true, nil
	}

	return false, nil
}
