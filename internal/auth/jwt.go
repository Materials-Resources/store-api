package auth

import (
	"fmt"
	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type CustomerClaims struct {
	Metadata struct {
		ContactId string `json:"contact_id"`
	} `json:"urn:zitadel:iam:user:metadata"`
	jwt.RegisteredClaims
}

func ParseJwt(tokenString string) (*CustomerClaims, error) {
	jwksUrl := "https://auth.materials-resources.com/oauth/v2/keys"
	jwks, err := keyfunc.NewDefault([]string{jwksUrl})
	if err != nil {
		log.Fatal(err)
	}
	token, err := jwt.ParseWithClaims(tokenString, &CustomerClaims{}, jwks.Keyfunc)
	if err != nil {
		return nil, err
	} else if claims, ok := token.Claims.(*CustomerClaims); ok {
		fmt.Println(claims.Metadata.ContactId, claims.RegisteredClaims.Issuer)
	} else {
		log.Fatal("unknown claims type, cannot proceed")
	}

	log.Print(token.Valid)
	log.Println(token.Claims.GetExpirationTime())

	return token.Claims.(*CustomerClaims), nil
}
