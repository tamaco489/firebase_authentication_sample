package firebase

import (
	"context"
	"log/slog"
	"sync"

	"google.golang.org/api/option"

	firebase_sdk "firebase.google.com/go/v4"
	firebase_auth "firebase.google.com/go/v4/auth"
)

type IFirebase interface {
	VerifyIDToken(ctx context.Context, idToken string) (*firebase_auth.Token, error)
}

var _ IFirebase = (*firebaseClient)(nil)

type firebaseClient struct {
	auth *firebase_auth.Client
}

// シングルトンインスタンスの設定
var (
	instance *firebaseClient
	once     sync.Once
)

func NewFirebaseClient(ctx context.Context, gsa []byte) (*firebaseClient, error) {

	// シングルトンインスタンスが初期化に失敗した場合、slogでエラーを出力し、呼び出し元に返す。
	once.Do(func() {
		// Firebase SDKの初期化
		opt := option.WithCredentialsJSON(gsa)
		app, err := firebase_sdk.NewApp(ctx, nil, opt)
		if err != nil {
			slog.ErrorContext(ctx, "failed to initialize firebase sdk", slog.String("error", err.Error()))
			return
		}

		// Firebase Authenticationクライアントの初期化
		authClient, err := app.Auth(ctx)
		if err != nil {
			slog.ErrorContext(ctx, "failed to initialize firebase sdk authentication client", slog.String("error", err.Error()))
			return
		}

		// インスタンスの作成
		instance = &firebaseClient{
			auth: authClient,
		}
	})

	return instance, nil
}

// 提供された ID トークンが正しい形式で、期限切れではなく、適切に署名されていれば、メソッドはデコードされた ID トークンを返します。
//
// デコードされたトークンからユーザーまたはデバイスの uid を取得できます。
//
// DOC: https://firebase.google.com/docs/auth/admin/verify-id-tokens?hl=ja&_gl=1*cin7r1*_up*MQ..*_ga*MTA5MjI0NTMzNi4xNzQyNjQ5MzQ1*_ga_CW55HF8NVT*MTc0MjY0OTM0NC4xLjAuMTc0MjY0OTM0NC4wLjAuMA..#verify_id_tokens_using_the_firebase_admin_sdk
func (fc *firebaseClient) VerifyIDToken(ctx context.Context, idToken string) (*firebase_auth.Token, error) {
	return fc.auth.VerifyIDToken(ctx, idToken)
}
