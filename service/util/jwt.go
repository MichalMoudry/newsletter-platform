package util

import (
	"context"
	"fmt"
	"time"

	"github.com/go-chi/jwtauth/v5"
	"github.com/google/uuid"
)

// Function for obtaining a map of claims from a JWT token.
func GetClaimsFromContext(ctx context.Context) (map[string]interface{}, error) {
	_, claims, err := jwtauth.FromContext(ctx)
	if err != nil {
		return map[string]interface{}{}, err
	}
	return claims, nil
}

// Function for obtaining content of the issued JWTs.
func GetJwtTokenContent(subject string, id uuid.UUID) map[string]interface{} {
	return map[string]interface{}{
		"sub":     subject,
		"iss":     "https://localhost:443",
		"aud":     []string{"https://localhost:443", "https://localhost:443"},
		"exp":     time.Now().UTC().AddDate(0, 0, 3),
		"user_id": id,
	}
}

// Function for obtaining a user id from context.
func GetUserIdFromContext(ctx context.Context) (uuid.UUID, error) {
	claims, err := GetClaimsFromContext(ctx)
	if err != nil {
		return uuid.Nil, err
	}
	return uuid.Parse(fmt.Sprint(claims["user_id"]))
}
