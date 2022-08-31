package server

import (
	"context"
	"net"

	"bannerRotator/internal/config"
	"google.golang.org/grpc"
)

type rotator interface {
	AddBanner(bannerID, slotID int) error
	RemoveBanner(bannerID, slotID int) error
	ClickBanner(bannerID, slotID, socialGroupID int) error
	Get(slotID, socialGroupID int) (*int, error)
}

type serverLogger interface {
	Error(message string)
	Info(message string)
	Warning(message string)
	Debug(message string)
	Fatal(message string)
}

type Server struct {
	rotator
	serverLogger
	grpc *grpc.Server
}

func NewGRPCServer(rotatorService rotator, serverLogger serverLogger) *Server {
	grpcServer := &Server{
		rotator:      rotatorService,
		grpc:         grpc.NewServer(),
		serverLogger: serverLogger,
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
	err := s.rotator.AddBanner(int(request.BannerId), int(request.SlotId))
	if err != nil {
		return &Error{Code: 1, Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) RemoveBanner(ctx context.Context, request *RemoveBannerRequest) (*Error, error) {
	err := s.rotator.RemoveBanner(int(request.BannerId), int(request.SlotId))
	if err != nil {
		return &Error{Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) ClickBanner(ctx context.Context, request *ClickBannerRequest) (*Error, error) {
	err := s.rotator.ClickBanner(int(request.BannerId), int(request.SlotId), int(request.SocialGroupId))
	if err != nil {
		return &Error{Message: err.Error()}, err
	}

	return &Error{}, nil
}

func (s Server) GetBanner(ctx context.Context, request *GetBannerRequest) (*GetBannerResponse, error) {
	response := &GetBannerResponse{}
	bannerID, err := s.rotator.Get(int(request.SlotId), int(request.SocialGroupId))
	if err != nil {
		response.Error = &Error{Message: err.Error()}
		return response, err
	}

	response.BannerId = int64(*bannerID)

	return response, nil
}

func (s Server) mustEmbedUnimplementedEventServiceServer() {}
