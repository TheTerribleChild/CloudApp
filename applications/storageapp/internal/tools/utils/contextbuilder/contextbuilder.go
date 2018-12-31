package contextbuilder

import (
	"context"
	"time"

	contextutil "github.com/TheTerribleChild/CloudApp/tools/utils/context"
)

func BuildStorageServerContext(token string) (ctx context.Context, CancelFunc func()) {
	builder := contextutil.ContextBuilder{}
	return builder.SetTimeout(2*time.Hour).AddHeader("authorization", token).Build()
}
