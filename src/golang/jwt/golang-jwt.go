package jwt

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type JwtEngine struct {
	secret []byte
	logger *zap.SugaredLogger
	//refreshTokens map[string]string
}

type Tokens struct {
	AccessToken        string `json:"accessToken"`
	AccessTokenExpiry  int64  `json:"accessExpiry"`
	RefreshToken       string `json:"refreshToken"`
	RefreshTokenExpiry int64  `json:"refreshTokenExpiry"`
}

type CustomClaims struct {
	UserId string `json:"userId"`
	jwt.RegisteredClaims
}

func New(secret []byte, logger *zap.SugaredLogger) *JwtEngine {
	return &JwtEngine{
		secret: secret,
		logger: logger,
	}
}

func (j *JwtEngine) GenTokens(user string) (*Tokens, error) {
	now := time.Now()
	accessDuration := time.Duration(5) * time.Minute
	refreshDuration := time.Duration(5) * time.Hour

	claims := CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(accessDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "RigelLedger",
			NotBefore: jwt.NewNumericDate(now),
			Subject:   "RigelLedger AccessToken",
		},
	}
	accessToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
	if err != nil {
		return nil, fmt.Errorf("JWT Access Token Creation error: %v", err)
	}

	claims = CustomClaims{
		user,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(now.Add(refreshDuration)),
			IssuedAt:  jwt.NewNumericDate(now),
			Issuer:    "RigelLedger",
			NotBefore: jwt.NewNumericDate(now),
			Subject:   "RigelLedger RefreshToken",
		},
	}
	refreshToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString(j.secret)
	if err != nil {
		return nil, fmt.Errorf("JWT Refresh Token Creation error: %v", err)
	}

	tokens := &Tokens{
		AccessToken:        accessToken,
		AccessTokenExpiry:  now.Add(accessDuration).Unix(),
		RefreshToken:       refreshToken,
		RefreshTokenExpiry: now.Add(refreshDuration).Unix(),
	}

	return tokens, nil
}

func (j *JwtEngine) RefreshToken(refreshToken string) {}

func (j *JwtEngine) ValidateToken(accessToken string) {}
