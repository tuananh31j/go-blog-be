package tokenStorage

import (
	"context"
	"fmt"
)

func (s *store) RemoveToken(ctx context.Context, userId string) error {
	key := fmt.Sprintf("user:%v", userId)
	cmd := s.rdb.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}
