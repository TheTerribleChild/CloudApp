package contextutil

import(
	"time"
	"context"
	"google.golang.org/grpc/metadata"
	"errors"
)

type ContextBuilder struct{
	timeout time.Duration;
	header map[string] string;
	toe string;
}

func (instance *ContextBuilder) Build() (ctx context.Context, CancelFunc func()){
	ctx = context.Background()
	if instance.timeout != 0 {
		ctx, CancelFunc = context.WithTimeout(ctx, instance.timeout)
	}
	if instance.header != nil {
		md := metadata.New(instance.header)
		ctx = metadata.NewOutgoingContext(ctx, md)
	}
	return ctx, CancelFunc
}

func (instance *ContextBuilder) AddHeader(key string, value string) *ContextBuilder {
	if instance.header == nil {
		instance.header = make(map[string] string)
	}
	instance.header[key] = value
	return instance
}

func (instance *ContextBuilder) SetTimeout(timeout time.Duration) *ContextBuilder {
	instance.timeout = timeout
	return instance
}

func (instance *ContextBuilder) SetToe(toe string) *ContextBuilder {
	instance.AddHeader("toe", toe)
	return instance
}

func GetHeaderContent(ctx context.Context, header string) ([]string, error) {
	if ctx == nil {
		return nil, errors.New("No context provided")
	}
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, errors.New("Unable to retrieve metadata from context")
	}
	contents := md.Get(header)
	return contents, nil
}

func GetToe(ctx context.Context) (string, error) {
	strs, err := GetHeaderContent(ctx, "toe")
	if len(strs) == 1{
		return strs[0], nil
	}
	return "", err
}