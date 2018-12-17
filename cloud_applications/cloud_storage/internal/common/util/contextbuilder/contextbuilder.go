package contextbuilder

import(
	"time"
	"context"
	contextutil "github.com/TheTerribleChild/cloud_appplication_portal/commons/utils/contextutil"
)

func BuildStorageServerContext(token string) (ctx context.Context, CancelFunc func()) {
	builder := contextutil.ContextBuilder{}
	return builder.SetTimeout(2*time.Hour).AddHeader("authorization", token).Build()
}