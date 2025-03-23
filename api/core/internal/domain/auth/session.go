package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

// session: セッション情報の構造体
type session struct {
	// 各認証機関で発行されたユニークなID
	Sub string `json:"sub"`

	// アプリケーションで作成されたユニークなID
	UID string `json:"uid"`

	// 認証機関を識別するためのタグ（firebase, auth0, github...）
	Provider string `json:"provider"`
}

// NewGetSessionData: セッション情報を取得するためのコンストラクタ
func NewGetSession(sub string) *session {
	return &session{
		Sub:      sub,
		UID:      "",
		Provider: "",
	}
}

// func NewSaveSession: セッション情報を保存するためのコンストラクタ
func NewSaveSession(
	sub string,
	uid string,
	provider string,
) *session {
	return &session{
		Sub:      sub,
		UID:      uid,
		Provider: provider,
	}
}

// ToJSON JSON 文字列に変換
func (s *session) ToJSON() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session data: %w", err)
	}
	return string(data), nil
}

// Save: firebase の sub を KEYとしてセッション情報を保存する
func (s *session) Save(ctx context.Context, client *redis.Client) error {
	if s.Sub == "" {
		return errors.New("provider's sub is empty")
	}

	data, err := s.ToJSON()
	if err != nil {
		return err
	}

	key := fmt.Sprintf("session:%s", s.Sub)

	// Redisにセッション情報を格納する
	// NOTE: firebaseの認証の有効期限が1時間のためそちらに合わせる。※1時間経つと自動的にRedisから削除される
	// NOTE: ただし本番運用の場合は1時間は短すぎるので24時間くらいが妥当だと思う。
	if err := client.Set(ctx, key, data, time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to save session data to redis: %w", err)
	}

	return nil
}

// Get: 指定したキーからセッション情報を取得する
func (s *session) Get(ctx context.Context, client *redis.Client) error {
	if s.Sub == "" {
		return errors.New("provider's sub is empty")
	}

	key := fmt.Sprintf("session:%s", s.Sub)

	// Redisにセッション情報を取得する
	data, err := client.Get(ctx, key).Result()
	if err != nil {
		// セッションが存在しない、または有効期限切れで削除された場合はnilを返す
		if err == redis.Nil {
			return nil
		}
		return fmt.Errorf("failed to get session data from redis: %w", err)
	}

	if err = json.Unmarshal([]byte(data), s); err != nil {
		return fmt.Errorf("failed to unmarshal session data: %w", err)
	}

	return nil
}
