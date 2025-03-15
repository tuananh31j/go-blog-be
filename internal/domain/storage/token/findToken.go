package tokenStorage

import (
	"context"
	"fmt"
)

func (s *store) FindToken(ctx context.Context, token string) (string, error) {
	key := fmt.Sprintf("user:%v", token)
	tk, err := s.rdb.Get(ctx, key).Result()
	if err != nil {
		return "", err
	}
	return tk, nil
}
