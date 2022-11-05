package main

import (
	"context"
	pb "github.com/wzslr321/server/trackers/proto"
	"google.golang.org/grpc"
	"log"
	"net"
)

type server struct {
	pb.UnimplementedTrackerServer
}

func (s *server) GetTracker(ctx context.Context, in *pb.GetTrackersRequest) (*pb.GetTrackersResponse, error) {
	return &pb.GetTrackersResponse{Trackers: "El Trackerinio"}, nil
}

func main() {
	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterTrackerServer(s, &server{})
	log.Printf("Server listening on port %v", listen.Addr())
	if err := s.Serve(listen); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
