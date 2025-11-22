package redis

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

func NewRedis(addr, pass string) *redis.Client {
	return redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: pass,
		DB:       0,
	})
}

func Subscribe(client *redis.Client, channel string) *redis.PubSub {
	return client.Subscribe(context.Background(), channel)
}

func Publish(client *redis.Client, channel string, msg string) error {
	return client.Publish(context.Background(), channel, msg).Err()
}

// ---- Presence / rooms helpers ----

const (
	keyOnlineUsers = "presence:online-users"
	keyRoomsSet    = "presence:rooms"
)

func roomKey(room string) string {
	return "process:room:" + room
}

func userConnKey(userID uint) string {
	return fmt.Sprintf("process:user:%d:connections", userID)
}

// TrackUserJoin updates:
// - global online users set
// - per-room set
// - rooms set
// - per-user connection count
func TrackUserJoin(ctx context.Context, c *redis.Client, room string, userID uint) error {
	pipe := c.TxPipeline()
	pipe.SAdd(ctx, keyOnlineUsers, userID)
	pipe.SAdd(ctx, keyRoomsSet, room)
	pipe.SAdd(ctx, roomKey(room), userID)
	pipe.Incr(ctx, userConnKey(userID))
	_, err := pipe.Exec(ctx)
	return err
}

func TrackUserLeave(ctx context.Context, c *redis.Client, room string, userID uint) error {
	pipe := c.TxPipeline()
	pipe.SRem(ctx, roomKey(room), userID)
	pipe.Decr(ctx, userConnKey(userID))
	if _, err := pipe.Exec(ctx); err != nil {
		return err
	}

	// If no more connections for this user, remove from global online set.
	n, err := c.Get(ctx, userConnKey(userID)).Int()
	if err == redis.Nil || n <= 0 {
		pipe2 := c.TxPipeline()
		pipe2.SRem(ctx, keyOnlineUsers, userID)
		pipe2.Del(ctx, userConnKey(userID))
		_, err = pipe2.Exec(ctx)
		if err != nil {
			return err
		}
	}

	// Optional: if room is empty, remove it from rooms set
	cnt, err := c.SCard(ctx, roomKey(room)).Result()
	if err == nil && cnt == 0 {
		_ = c.SRem(ctx, keyRoomsSet, room).Err()
	}
	return nil
}

func GetOnlineUsers(ctx context.Context, c *redis.Client) ([]string, error) {
	vals, err := c.SMembers(ctx, keyOnlineUsers).Result()
	if err == redis.Nil {
		return []string{}, nil
	}
	return vals, err
}

func GetRooms(ctx context.Context, c *redis.Client) ([]string, error) {
	vals, err := c.SMembers(ctx, keyRoomsSet).Result()
	if err == redis.Nil {
		return []string{}, nil
	}
	return vals, err
}
func GetRoomUserCount(ctx context.Context, c *redis.Client, room string) (int64, error) {
	cnt, err := c.SCard(ctx, roomKey(room)).Result()
	if err == redis.Nil {
		return 0, nil
	}
	return cnt, err
}
