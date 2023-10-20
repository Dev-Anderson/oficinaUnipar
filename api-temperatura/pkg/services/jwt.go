package services

import (
	"api-temperatura/internal/config"
	"fmt"
	"log"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
)

type jwtServices struct {
	secretKey string
	issure    string
}

func NewJWTService() *jwtServices {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalln("Erro ao buscar as variaveis de ambiente ", err.Error())
	}

	return &jwtServices{
		secretKey: env.SecretKey,
		issure:    env.Issure,
	}
}

type Claim struct {
	Sub uint `json:"sub"`
	jwt.StandardClaims
}

func (s *jwtServices) GenerateToken(id uint) (string, error) {
	env, err := config.LoadEnv()
	if err != nil {
		log.Fatalln("Erro ao buscar as variaveis de ambiente", err.Error())
	}

	timeExpires, err := strconv.ParseInt(env.TimeExpires, 10, 64)
	if err != nil {
		log.Fatalln("Erro ao converter o timeExpires ", err.Error())
	}

	claim := &Claim{
		id,
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * time.Duration(timeExpires)).Unix(),
			Issuer:    s.issure,
			IssuedAt:  time.Now().Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claim)

	t, err := token.SignedString([]byte(s.secretKey))
	if err != nil {
		return "", nil
	}

	return t, nil
}

func (s *jwtServices) ValidateToken(token string) bool {
	_, err := jwt.Parse(token, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Token invali invalido %s", token)
		}
		return []byte(s.secretKey), nil
	})

	return err == nil
}
