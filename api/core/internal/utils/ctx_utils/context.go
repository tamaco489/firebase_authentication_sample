package ctx_utils

import "github.com/gin-gonic/gin"

type authUIDKey string

const (
	firebaseUIDKey authUIDKey = "firebase_uid"
	auth0UIDKey    authUIDKey = "auth0_uid"
	githubUIDKey   authUIDKey = "github_uid"
	coreUIDKey     authUIDKey = "core_uid"
)

func (ak authUIDKey) String() string {
	return string(ak)
}

func SetFirebaseUID(ctx *gin.Context, firebaseUID string) {
	ctx.Set(firebaseUIDKey.String(), firebaseUID)
}

func GetFirebaseUID(ctx *gin.Context) (string, bool) {
	sub, ok := ctx.Value(firebaseUIDKey.String()).(string)
	return sub, ok
}

func SetCoreUID(ctx *gin.Context, uid string) {
	ctx.Set(coreUIDKey.String(), uid)
}

func GetCoreUID(ctx *gin.Context) (string, bool) {
	uid, ok := ctx.Value(coreUIDKey.String()).(string)
	return uid, ok
}

type authProviderTypeKey string

const (
	FirebaseProviderKey authProviderTypeKey = "firebase"
)

func (ap authProviderTypeKey) String() string {
	return string(ap)
}

func SetFirebaseProviderType(ctx *gin.Context, providerType string) {
	ctx.Set(FirebaseProviderKey.String(), providerType)
}

func GetFirebaseProviderType(ctx *gin.Context) (string, bool) {
	provider, ok := ctx.Value(FirebaseProviderKey.String()).(string)
	return provider, ok
}
