package authstore

import (
	"context"
	"fmt"
	"time"
)

func (s *store) SaveRefreshToken(ctx context.Context, token, userId string) error {
	key := fmt.Sprintf("user:%v", userId)
	cmd := s.rdb.Set(ctx, key, token, 7*24*60*60*time.Second)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}
