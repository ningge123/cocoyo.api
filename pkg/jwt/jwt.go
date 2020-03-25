package jwt

import (
	"cocoyo/models"
	"cocoyo/pkg/response"
	"cocoyo/pkg/setting"
	"errors"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
	"time"
)

var jwtSecret = []byte(setting.LoadJwt().Key("JWT_SECRET").String())

const ContextKeyUserObj = "jwt_token"
const bearerLength = len("Bearer ")

type Claims struct {
	models.User
	jwt.StandardClaims
}

func GenerateToken(user *models.User) (string, error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(3 * time.Hour)

	claims := Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt:	expireTime.Unix(),
			IssuedAt: 	time.Now().Unix(),
			Id:       	fmt.Sprintf("%d", user.Id),
		},
		User: *user,
	}

	tokenClaims := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	token, err := tokenClaims.SignedString(jwtSecret)

	return token, err
}

func ParseToken(token string) (*models.User, error)  {
	if token == "" {
		return nil, errors.New("no token is found in Authorization Bearer")
	}

	claims := Claims{}
	_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(jwtSecret), nil
	})

	if err != nil {
		return nil, err
	}
	
	return &claims.User, err
}

func (c Claims) Valid() error {
	if c.VerifyExpiresAt(time.Now().Unix(), true) == false {
		return errors.New("token is expired")
	}

	if c.User.Id < 1 {
		return errors.New("invalid user in jwt")
	}

	return nil
}

func TokenAuthMiddleware(c *gin.Context)  {
	token := c.GetHeader("Authorization")
	var err error

	if len(token) < bearerLength {
		err = errors.New("authorization")
	}

	token = strings.TrimSpace(token[bearerLength:])

	_, err = ParseToken(token)

	if err != nil {
		c.Abort()
		c.JSON(http.StatusUnauthorized, response.Authorization())
	} else {
		//store the user Model in the context
		c.Set(ContextKeyUserObj, token)
		c.Next()
	}
}