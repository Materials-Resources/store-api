package session

import (
	"context"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/golang-jwt/jwt/v5"
	"log"
)

type JWTClaims struct {
	Metadata struct {
		ContactID string `json:"contact_id"`
		BranchID  string `json:"branch_id"`
	} `json:"urn:zitadel:iam:user:metadata"`
	jwt.RegisteredClaims
}

func (m *Manager) ParseJwt(ctx context.Context, tokenString string) (*JWTClaims, error) {
	claims := &JWTClaims{}
	token, err := jwt.ParseWithClaims(tokenString, claims, m.keyFunc)
	if err != nil {
		return nil, err
	}
	if _, ok := token.Claims.(*JWTClaims); !ok {
		log.Fatal("unknown claims type, cannot proceed")
	}

	// Decode metadata fields
	if err := decodeBase64Metadata(claims); err != nil {
		return nil, fmt.Errorf("decoding metadata: %w", err)
	}

	// Verify user is active in Zitadel
	if claims.Subject != "" {
		isActive, err := m.zitadelClient.IsUserActive(ctx, claims.Subject)
		if err != nil {
			return nil, fmt.Errorf("checking user active status: %w", err)
		}
		if !isActive {
			return nil, errors.New("user is not active")
		}
	}

	return claims, nil
}

func decodeBase64Metadata(claims *JWTClaims) error {
	// Helper function to decode a Base64 string to a regular string.
	decodeField := func(field *string) error {
		decoded, err := base64.RawStdEncoding.DecodeString(*field)
		if err != nil {
			return fmt.Errorf("base64 decoding: %w", err)
		}
		*field = string(decoded)
		return nil
	}

	if err := decodeField(&claims.Metadata.ContactID); err != nil {
		return err
	}
	if err := decodeField(&claims.Metadata.BranchID); err != nil {
		return err
	}

	return nil
}
