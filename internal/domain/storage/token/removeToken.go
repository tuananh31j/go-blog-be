package tokenStorage

import (
	"context"
	"fmt"
)

func (s *store) RemoveToken(ctx context.Context, token string) error {
	key := fmt.Sprintf("user:%v", token)
	cmd := s.rdb.Del(ctx, key)
	if err := cmd.Err(); err != nil {
		return err
	}
	return nil
}
