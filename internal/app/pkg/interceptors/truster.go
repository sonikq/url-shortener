package interceptors

import (
	"context"
	"fmt"
	"google.golang.org/genproto/googleapis/rpc/code"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
	"net"
	"strings"
)

// UnaryServerInterceptorOpts middleware for trusted IP
func UnaryServerInterceptorOpts(trustedSubnet string) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (any, error) {
		md, _ := metadata.FromIncomingContext(ctx)
		method := strings.Split(info.FullMethod, "/")[2]

		if method == "GetStats" {
			trasted := false
			for k, v := range md {
				if strings.Contains(":authority x-real-ip x-forwarded-for", k) {
					ip, _, err := net.SplitHostPort(v[0])
					if err != nil {
						fmt.Println("invalid addr: ", err)
						continue
					}
					if strings.Contains(trustedSubnet, ip) {
						trasted = true
						break
					}
				}
			}

			if !trasted {
				return nil, status.Errorf(codes.Code(code.Code_CANCELLED), "Acccess only trasted iP")
			}
		}

		return handler(ctx, req)
	}
}
