package crypto

import (
	"github.com/dgrijalva/jwt-go"
	"github.com/z13z/Kiosks/master-server/db/users"
	"time"
)

//todo zaza
var signatureKey = []byte("changeIt")

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
			IssuedAt: time.Now().Unix(),
		},
	}
	return jwt.NewWithClaims(jwt.SigningMethodHS256, userClaims).SignedString(signatureKey)
}
