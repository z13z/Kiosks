package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/z13z/Kiosks/master-server/db/users"
	"log"
	"time"
)

//todo zaza
var signatureKey = []byte("changeIt")
var durationForKey, _ = time.ParseDuration("1m")

type KioskUserClaims struct {
	Username    string `json:"username"`
	Permissions string `json:"permissions"`
	jwt.StandardClaims
}

func GetJwtForUser(entity users.UserEntity) (string, error) {
	val, err := entity.Permissions.Value()
	if err != nil {
		return "", err
	}
	permissions := val.(string)
	userClaims := KioskUserClaims{
		entity.Name,
		permissions,
		jwt.StandardClaims{
			IssuedAt:  time.Now().Unix(),
			ExpiresAt: time.Now().Add(durationForKey).Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString(signatureKey)
}

func CheckJwt(tokenToCheck string) (*KioskUserClaims, bool) {
	if tokenToCheck == "" {
		return nil, false
	}
	token, err := jwt.ParseWithClaims(tokenToCheck, &KioskUserClaims{}, func(token *jwt.Token) (interface{}, error) {
		return signatureKey, nil
	})
	if err != nil {
		log.Print("Can't parse jwt: "+tokenToCheck, err)
		return nil, false
	}
	parsedClaims, ok := token.Claims.(*KioskUserClaims)
	if token.Valid && ok {
		return parsedClaims, ok
	} else {
		return nil, false
	}
}
