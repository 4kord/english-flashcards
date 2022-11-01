package auth

import (
	"context"
)

func (s *service) Refresh(ctx context.Context, refreshToken string) (string, string, error) {
	return "", "", nil
}
