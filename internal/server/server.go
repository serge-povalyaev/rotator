package server

import (
	"context"
	"net"

	"bannerRotator/internal/config"
	"bannerRotator/internal/logger"
	"bannerRotator/internal/service"
	"google.golang.org/grpc"
)

type Server struct {
	rotatorService *service.RotatorService
	grpc           *grpc.Server
	logger         *logger.Logger
}

func NewGRPCServer(rotatorService *service.RotatorService, logger *logger.Logger) *Server {
	grpcServer := &Server{
		rotatorService: rotatorService,
		grpc:           grpc.NewServer(),
		logger:         logger,
	}
	RegisterEventServiceServer(grpcServer.grpc, grpcServer)

	return grpcServer
}

func (s *Server) Start(config config.GRPC) error {
	lsn, err := net.Listen("tcp", net.JoinHostPort(config.Host, config.Port))
	if err != nil {
		return err
	}

	if err = s.grpc.Serve(lsn); err != nil {
		return err
	}

	return nil
}

func (s *Server) Stop() {
	s.grpc.GracefulStop()
}

func (s Server) AddBanner(ctx context.Context, request *AddBannerRequest) (*Error, error) {
	err := s.rotatorService.AddBanner(int(request.BannerId), int(request.SlotId))
	if err != nil {
		return &Error{Code: 1, Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) RemoveBanner(ctx context.Context, request *RemoveBannerRequest) (*Error, error) {
	err := s.rotatorService.RemoveBanner(int(request.BannerId), int(request.SlotId))
	if err != nil {
		return &Error{Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) ClickBanner(ctx context.Context, request *ClickBannerRequest) (*Error, error) {
	err := s.rotatorService.ClickBanner(int(request.BannerId), int(request.SlotId), int(request.SocialGroupId))
	if err != nil {
		return &Error{Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) GetBanner(ctx context.Context, request *GetBannerRequest) (*GetBannerResponse, error) {
	response := &GetBannerResponse{}
	bannerID, err := s.rotatorService.Get(int(request.SlotId), int(request.SocialGroupId))
	if err != nil {
		response.Error = &Error{Message: err.Error()}
		return response, err
	}

	response.BannerId = int64(*bannerID)

	return response, nil
}

func (s Server) mustEmbedUnimplementedEventServiceServer() {}
