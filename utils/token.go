package utils

import (
	"encoding/base64"
	"fmt"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// refs https://codevoweb.com/golang-mongodb-jwt-authentication-authorization/

func CreateToken(ttl time.Duration, id interface{}, ref int, privateKey string) (string, error) {
	decodedPrivateKey, err := base64.StdEncoding.DecodeString(privateKey)
	if err != nil {
		return "", fmt.Errorf("could not decode key: %w", err)
	}
	key, err := jwt.ParseRSAPrivateKeyFromPEM(decodedPrivateKey)

	if err != nil {
		return "", fmt.Errorf("create: parse key: %w", err)
	}

	now := time.Now().UTC()

	claims := make(jwt.MapClaims)
	claims["sub"] = id
	claims["ref"] = ref
	claims["exp"] = now.Add(ttl).Unix()
	claims["iat"] = now.Unix()
	claims["nbf"] = now.Unix()

	token, err := jwt.NewWithClaims(jwt.SigningMethodRS256, claims).SignedString(key)

	if err != nil {
		return "", fmt.Errorf("create: sign token: %w", err)
	}

	return token, nil
}

type TokenModel struct {
	Id   string  `json:"id" bson:"id"`
	Role float64 `json:"role" bson:"role"`
}

func ValidateToken(token string, publicKey string) (*TokenModel, error) {
	decodedPublicKey, err := base64.StdEncoding.DecodeString(publicKey)

	if err != nil {
		return nil, fmt.Errorf("could not decode: %w", err)
	}

	key, err := jwt.ParseRSAPublicKeyFromPEM(decodedPublicKey)

	if err != nil {
		return nil, fmt.Errorf("validate: parse key: %w", err)
	}

	parsedToken, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodRSA); !ok {
			return nil, fmt.Errorf("unexpected method: %s", t.Header["alg"])
		}
		return key, nil
	})

	if err != nil {
		return nil, fmt.Errorf("validate: %w", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		return nil, fmt.Errorf("validate: invalid token")
	}

	r := TokenModel{
		Id:   claims["sub"].(string),
		Role: claims["ref"].(float64),
	}
	return &r, nil
}

// @ พี่กล้าแนะนำมา 13-07-2022 23:25
func Permision(listPermission []int) fiber.Handler {
	return func(c *fiber.Ctx) error {
		// TODO อยากรับ Argument เพิ่มอ่ะครับ
		for _, a := range listPermission {
			// if a == c.Locals("role") {
			if a == 6 {
				return c.Next()
			}
		}
		return c.Status(fiber.StatusForbidden).JSON(fiber.Map{
			"code":    fiber.StatusForbidden,
			"status":  false,
			"message": "permission denied",
			"data":    "",
		})
	}
}
