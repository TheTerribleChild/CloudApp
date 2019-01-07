package contextbuilder

import (
	"context"
	"time"

	contextutil "theterriblechild/CloudApp/tools/utils/context"
)

func BuildStorageServerContext(token string) (ctx context.Context, CancelFunc func()) {
	builder := contextutil.ContextBuilder{}
	return builder.SetTimeout(2 * time.Hour).SetAuth(token).Build()
}
