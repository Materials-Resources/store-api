package session

import (
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type Parser struct {
	keyFunc jwt.Keyfunc
}

func NewParser(jwksUrl string) *Parser {
	jwks, err := keyfunc.NewDefault([]string{jwksUrl})
	if err != nil {
		log.Fatal(err)
	}
	return &Parser{
		keyFunc: jwks.Keyfunc,
	}
}

type JWTClaims struct {
	Metadata struct {
		ContactID string `json:"contact_id"`
		BranchID  string `json:"branch_id"`
	} `json:"urn:zitadel:iam:user:metadata"`
	jwt.RegisteredClaims
}

func (p *Parser) ParseJwt(tokenString string) (*JWTClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, p.keyFunc)
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*JWTClaims); !ok {
		log.Fatal("unknown claims type, cannot proceed")
	}

	return token.Claims.(*JWTClaims), nil
}
