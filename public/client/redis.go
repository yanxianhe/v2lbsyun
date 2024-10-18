package client

import (
	"context"
	"fmt"
	"time"
	"v2lbsyun/config"
	"v2lbsyun/public/utils"

	"github.com/go-redis/redis/v8"
)

// InitRedis 初始化 Redis 客户端
//
// 根据配置文件中的 Redis 配置信息，创建一个 Redis 客户端并返回。
// 如果没有设置密码，则 Password 字段为空字符串。
// 如果未指定 Redis 数据库名，则默认使用默认数据库。
//
// 返回值：
//
//	*redis.Client：Redis 客户端实例
func InitRedis() *redis.Client {
	// 根据实际情况配置连接参数
	rdb := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.RY_MAP_REDIS_HOST,
		Password: utils.GetDesDecrypt(config.Cfg.Redis.RY_MAP_REDIS_PWD), // no password set
		DB:       config.Cfg.Redis.RY_MAP_REDIS_NAME,                     // use default DB
	})
	fmt.Printf("redis client init success %v\n", rdb.String())
	return rdb
}

// KeyCount 根据给定的模式查询 Redis 中的键数量
//
// 参数：
//   - rdb: Redis 客户端指针
//   - pattern: Redis 键的模式字符串，支持通配符 * 和 ?
//
// 返回值：
//   - int64: 查询到的键的数量，如果查询失败则返回 -1
func KeyCount(rdb *redis.Client, pattern string) int64 {
	keys, err := rdb.Keys(context.Background(), pattern).Result()
	if err != nil {
		fmt.Printf("redis.Keys failed: %v\n", err)
		return -1
	}
	return int64(len(keys))
}

// KeyCount 从Redis数据库中根据模式pattern查询匹配的键的数量
//
// 参数：
// - rdb: Redis客户端指针
// - pattern: 匹配模式
//
// 返回值：
// - int64: 匹配键的数量，如果查询失败则返回-1
func GetValue(rdb *redis.Client, key string) (string, error) {
	ctx := context.Background()
	val, err := rdb.Get(ctx, key).Result()
	if err == redis.Nil {
		fmt.Printf("redis.Get failed: %v\n", err)
		return "", nil // Key not found
	} else if err != nil {
		fmt.Printf("redis.Get failed: %v\n", err)
		return "", err
	}
	return val, nil
}

// SetValue 将给定的键值对存储到Redis中，并设置过期时间
// 参数:
//
//	rdb: *redis.Client Redis客户端实例
//	key: string 键名
//	value: interface{} 键值
//	duration: time.Duration 键值的有效期
//
// 返回值:
//
//	error: 如果存储失败，则返回错误信息；否则返回nil
func SetValue(rdb *redis.Client, key string, value interface{}, duration int64) error {
	ctx := context.Background()
	err := rdb.Set(ctx, key, value, time.Duration(duration)*time.Second).Err()
	if err != nil {
		fmt.Printf("redis.Set failed: %v\n", err)
		return err
	}
	return nil
}

func DelValue(rdb *redis.Client, key string) error {
	return rdb.Del(context.Background(), key).Err()
}
