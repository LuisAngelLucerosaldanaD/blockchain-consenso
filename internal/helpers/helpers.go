package helpers

import (
	"bjungle-consenso/internal/env"
	"bjungle-consenso/internal/logger"
	"bjungle-consenso/internal/models"
	"crypto/rsa"
	"encoding/json"
	"fmt"
	"github.com/form3tech-oss/jwt-go"
	"github.com/gofiber/fiber/v2"
	"io/ioutil"
	"strings"
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

func SliceToString(s []string) string {
	var elements string

	for i, e := range s {
		if e == "" {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%s'", strings.ToLower(e))
		} else {
			elements = fmt.Sprintf("%s,'%s'", elements, strings.ToLower(e))
		}

	}
	return elements
}

func SliceInt64ToString(ids []int64) string {
	var elements string
	for i, e := range ids {
		if e == 0 {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%d'", e)
		} else {
			elements = fmt.Sprintf("%s,'%d'", elements, e)
		}

	}
	return elements
}

func SliceInt64ToStringInteger(ids []int64) string {
	var elements string
	for i, e := range ids {
		if e == 0 {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("%d", e)
		} else {
			elements = fmt.Sprintf("%s,%d", elements, e)
		}

	}
	return elements
}

func SliceIntToString(ids []int) string {
	var elements string
	for i, e := range ids {
		if i == 0 {
			elements = fmt.Sprintf("'%d'", e)
		} else {
			elements = fmt.Sprintf("%s,'%d'", elements, e)
		}

	}
	return elements
}

func SlicePointerToString(s []*string) string {
	var elements string
	for i, e := range s {
		if e == nil {
			continue
		}
		if i == 0 {
			elements = fmt.Sprintf("'%s'", strings.ToLower(*e))
		} else {
			elements = fmt.Sprintf("%s,'%s'", elements, strings.ToLower(*e))
		}

	}
	return elements
}

func InterfaceToMapInterface(s interface{}) (map[string]interface{}, error) {
	rs := make(map[string]interface{}, 0)
	data, err := json.Marshal(s)
	if err != nil {
		logger.Error.Println("couldn't convert interface to map[string]interface{}: ", err)
		return rs, err
	}
	err = json.Unmarshal(data, &rs)
	if err != nil {
		logger.Error.Println("couldn't convert interface to map[string]interface{}: ", err)
		return rs, err
	}
	return rs, nil

}
