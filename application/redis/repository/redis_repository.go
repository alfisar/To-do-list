package repository

import (
	"context"
	"fmt"
	"time"
	"todolist/internal/errorhandler"

	"github.com/go-redis/redis/v8"
)

type redisRepository struct{}

func NewRedisRepository() *redisRepository {
	return &redisRepository{}
}

func (r *redisRepository) Insert(ctx context.Context, conn *redis.Client, key string, data string, exp time.Duration) (err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
	}

	errData := conn.Set(ctx, key, data, exp).Err()
	if errData != nil {
		err = fmt.Errorf("insert redis error : %w", errData)
		return
	}

	return

}

func (r *redisRepository) Get(ctx context.Context, conn *redis.Client, key string) (result string, err error) {
	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf(fmt.Sprintf("%s", r))
		}

	}()

	if conn == nil {
		err = fmt.Errorf(errorhandler.ErrMsgConnEmpty)
	}

	errData := conn.Get(ctx, key).Err()
	if errData != nil {
		err = fmt.Errorf("get redis error : %w", errData)
		return
	}

	return
}
