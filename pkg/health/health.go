package health

import (
	"context"

	grpc2 "google.golang.org/grpc"
	"google.golang.org/grpc/health/grpc_health_v1"

	grpc "github.com/spiral/php-grpc"
)

const ID = "health"

type Service struct {
	grpc_health_v1.UnimplementedHealthServer
}

func (s *Service) Init(g *grpc.Service) (bool, error) {
	return true, g.AddService(func(server *grpc2.Server) {
		grpc_health_v1.RegisterHealthServer(server, s)
	})
}

func (*Service) Check(context.Context, *grpc_health_v1.HealthCheckRequest) (*grpc_health_v1.HealthCheckResponse, error) {
	return &grpc_health_v1.HealthCheckResponse{Status: grpc_health_v1.HealthCheckResponse_SERVING}, nil
}
