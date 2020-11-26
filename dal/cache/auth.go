package cache

import (
	"math/rand"
	"strings"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/fgunawan1995/julo/model"
	"github.com/pkg/errors"

	cache "github.com/patrickmn/go-cache"
)

//GetToken checks whether token exists and return user_id
func (dal *DAL) GetToken(token string) (string, error) {
	user, found := dal.Cache.Get(token)
	if !found {
		return "0", errors.New("token not found")
	}
	userInfo := strings.Split(user.(string), " ")
	return userInfo[1], nil
}

func (dal *DAL) SetToken(userID, token string) {
	dal.Cache.Set(token, "user_id "+userID, cache.DefaultExpiration)
}

func (dal *DAL) SetTokenForUser(userID string) (string, error) {
	token, err := dal.generateToken()
	if err != nil {
		return token, err
	}
	dal.Cache.Set(token, "user_id "+userID, cache.DefaultExpiration)
	return token, nil
}

func (dal *DAL) generateToken() (string, error) {
	createdAt := time.Now().UTC()
	randomVar := rand.Int()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"random":     randomVar,
		"created_at": createdAt,
	})
	return token.SignedString([]byte(model.SecretKey))
}
