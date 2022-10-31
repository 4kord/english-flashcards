package maker

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var (
	ErrInvalidToken = errors.New("token is invalid")
	ErrExpiredToken = errors.New("token has expired")
)

type AccessPayload struct {
	Sub   int32     `json:"sub"`
	Admin bool      `json:"admin"`
	Exp   time.Time `json:"exp"`
}

type RefreshPayload struct {
	Sub int32     `json:"sub"`
	Exp time.Time `json:"exp"`
}

func (payload *AccessPayload) Valid() error {
	if time.Now().After(payload.Exp) {
		return ErrExpiredToken
	}

	return nil
}

func (payload *RefreshPayload) Valid() error {
	if time.Now().After(payload.Exp) {
		return ErrExpiredToken
	}

	return nil
}

type Maker interface {
	CreateAccessToken(sub int32, admin bool, exp time.Duration) (string, error)
	CreateRefreshToken(sub int32, exp time.Duration) (string, error)
	VerifyAccessToken(token string) (*AccessPayload, error)
	VerifyRefreshToken(token string) (*RefreshPayload, error)
}

type JWTMaker struct {
	secretKey string
}

func NewAccessPayload(sub int32, admin bool, exp time.Duration) *AccessPayload {
	return &AccessPayload{
		Sub:   sub,
		Admin: admin,
		Exp:   time.Now().Add(exp),
	}
}

func NewRefreshPayload(sub int32, exp time.Duration) *RefreshPayload {
	return &RefreshPayload{
		Sub: sub,
		Exp: time.Now().Add(exp),
	}
}

func NewJWTMaker(secretKey string) Maker {
	return &JWTMaker{secretKey}
}

func (maker *JWTMaker) CreateAccessToken(sub int32, admin bool, exp time.Duration) (string, error) {
	payload := NewAccessPayload(sub, admin, exp)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) CreateRefreshToken(sub int32, exp time.Duration) (string, error) {
	payload := NewRefreshPayload(sub, exp)

	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)

	return jwtToken.SignedString([]byte(maker.secretKey))
}

func (maker *JWTMaker) VerifyAccessToken(token string) (*AccessPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &AccessPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*AccessPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}

func (maker *JWTMaker) VerifyRefreshToken(token string) (*RefreshPayload, error) {
	keyFunc := func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, ErrInvalidToken
		}

		return []byte(maker.secretKey), nil
	}

	jwtToken, err := jwt.ParseWithClaims(token, &RefreshPayload{}, keyFunc)
	if err != nil {
		verr, ok := err.(*jwt.ValidationError)
		if ok && errors.Is(verr.Inner, ErrExpiredToken) {
			return nil, ErrExpiredToken
		}

		return nil, ErrInvalidToken
	}

	payload, ok := jwtToken.Claims.(*RefreshPayload)
	if !ok {
		return nil, ErrInvalidToken
	}

	return payload, nil
}
