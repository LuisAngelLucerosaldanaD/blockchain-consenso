package helpers

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"bytes"
	"crypto/rsa"
	"encoding/binary"
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
)

var (
	signKey   *rsa.PublicKey
	publicKey string
)

type UserClaims struct {
	jwt.StandardClaims
	User string `json:"user"`
	Role int    `json:"role"`
}

func init() {
	c := env.NewConfiguration()
	publicKey = c.App.RSAPublicKey
	keyBytes, err := ioutil.ReadFile(publicKey)
	if err != nil {
		logger.Error.Printf("leyendo el archivo privado de firma: %s", err)
	}

	signKey, err = jwt.ParseRSAPublicKeyFromPEM(keyBytes)
	if err != nil {
		logger.Error.Printf("realizando el parse en auth RSA private: %s", err)
	}
}

func GetUserContext(c *fiber.Ctx) *models.User {
	userUrl := c.Locals("user").(*jwt.Token)
	claims := userUrl.Claims.(jwt.MapClaims)
	for i, cl := range claims {
		if i == "user" {
			u := models.User{}
			ub, _ := json.Marshal(cl)
			_ = json.Unmarshal(ub, &u)
			return &u
		}
	}
	return nil
}

func GetUserContextV2(c *fiber.Ctx) (*models.User, error) {
	tokenStr := c.Get("Authorization")
	claims := jwt.MapClaims{}
	_, err := jwt.ParseWithClaims(tokenStr[7:], &claims, func(token *jwt.Token) (interface{}, error) {
		return signKey, nil
	})
	if err != nil {
		return nil, err
	}

	for i, cl := range claims {
		if i == "user" {
			u := models.User{}
			ub, _ := json.Marshal(cl)
			_ = json.Unmarshal(ub, &u)
			return &u, nil
		}
	}

	return nil, nil
}

func ToHex(num int64) []byte {
	buff := new(bytes.Buffer)
	err := binary.Write(buff, binary.BigEndian, num)
	if err != nil {
		log.Panic(err)

	}
	return buff.Bytes()
}
