package grpcutil

import "github.com/grpc-ecosystem/grpc-gateway/runtime"

func CustomMatcher(key string) (string, bool) {
    switch key {
    case "X-Cloudapp-Authorization":
        return "authorization", true
    case "X-Cloudapp-Toe":
        return "toe", true
    default:
        return runtime.DefaultHeaderMatcher(key)
    }
}

