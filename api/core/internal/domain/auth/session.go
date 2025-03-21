package auth

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"log/slog"
	"time"

	"github.com/redis/go-redis/v9"
)

// SessionData セッション情報の構造体
type sessionData struct {
	// 各認証機関で発行されたユニークなID
	Sub string `json:"sub"`

	// 認証が完了した日時（unix time）
	AuthTime int64 `json:"auth_time"`

	// 認証が失効する日時（unix time）
	Exp int64 `json:"exp"`

	// アプリケーションで作成されたユニークなID
	UID string `json:"uid"`

	// 認証機関を識別するためのタグ（firebase, auth0, github...）
	Provider string `json:"provider"`
}

// NewDefaultSessionData デフォルト値付きのコンストラクタ
func NewDefaultSessionData(
	sub string,
	authTime int64,
	expire int64,
	uid string,
	provider string,
) *sessionData {
	return &sessionData{
		Sub:      sub,
		AuthTime: authTime,
		Exp:      expire,
		UID:      uid,
		Provider: provider,
	}
}

// IsExpired: セッションが有効期限切れかチェック
func (s *sessionData) IsExpired() bool {
	return time.Now().Unix() > s.Exp
}

// ToJSON JSON 文字列に変換
func (s *sessionData) ToJSON() (string, error) {
	data, err := json.Marshal(s)
	if err != nil {
		return "", fmt.Errorf("failed to marshal session data: %w", err)
	}
	return string(data), nil
}

// SaveToRedis: firebase の sub を KEYとしてセッション情報を保存する
func (s *sessionData) SaveToRedis(ctx context.Context, client *redis.Client) error {
	if s.Sub == "" {
		return errors.New("provider's sub is empty")
	}

	data, err := s.ToJSON()
	if err != nil {
		return err
	}

	key := fmt.Sprintf("session:%s", s.Sub)

	// Redisにセッション情報を格納する
	// NOTE: firebaseの認証の有効期限が1時間のためそちらに合わせる、ただし本番運用の場合は1時間は短すぎるので24時間くらいが妥当だと思う。
	if err := client.Set(ctx, key, data, time.Hour).Err(); err != nil {
		return fmt.Errorf("failed to save session data to redis: %w", err)
	}

	// todo: 検証終了後削除
	slog.InfoContext(ctx, "success save to redis.", slog.String("key", key))

	return nil
}
