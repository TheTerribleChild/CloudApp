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
}

const(
	Toe = "toe"
	Auth = "authorization"
	UserName = "username"
	UserID = "userid"
	AgentID = "agentid"
)

func GetContextBuilder() *ContextBuilder {
	return &ContextBuilder{header:make(map[string]string)}
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
	instance.AddHeader(Toe, toe)
	return instance
}

func (instance *ContextBuilder) SetAuth(auth string) *ContextBuilder {
	instance.AddHeader(Auth, auth)
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

func GetAuth(ctx context.Context) (string, error) {
	strs, err := GetHeaderContent(ctx, Auth)
	if len(strs) == 1{
		return strs[0], nil
	}
	return "", err
}

func GetToe(ctx context.Context) (str string, ok bool) {
	str, ok = ctx.Value(Toe).(string)
	return
}


func SetToe(ctx context.Context, toe string) (context.Context) {
	if ctx == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, Toe, toe)
	return ctx
}

func GetUserName(ctx context.Context) (str string, ok bool) {
	str, ok = ctx.Value(UserName).(string)
	return
}

func SetUserName(ctx context.Context, userName string) (context.Context) {
	if ctx == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, UserName, userName)
	return ctx
}

func GetUserID(ctx context.Context) (str string, ok bool) {
	str, ok = ctx.Value(UserID).(string)
	return
}

func SetUserID(ctx context.Context, userID string) (context.Context) {
	if ctx == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, UserID, userID)
	return ctx
}

func GetAgentID(ctx context.Context) (str string, ok bool) {
	str, ok = ctx.Value(AgentID).(string)
	return
}

func SetAgentID(ctx context.Context, agentID string) (context.Context) {
	if ctx == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, AgentID, agentID)
	return ctx
}

func SetAuth(ctx context.Context, auth string) (context.Context) {
	if ctx == nil {
		return ctx
	}
	ctx = context.WithValue(ctx, "authorization", auth)
	return ctx
}