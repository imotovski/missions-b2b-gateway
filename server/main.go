package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/imotovski/missions-b2b-gateway/gen/missions"
	"google.golang.org/grpc"
)

var (
	port = flag.Int("port", 9090, "The server port")
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedMissionServiceServer
}

func (s *server) ListMissions(
	ctx context.Context,
	req *pb.ListMissionsRequest,
) (*pb.ListMissionsResponse, error) {
	log.Printf("Received: %v", req)
	return &pb.ListMissionsResponse{Missions: []*pb.ListMissionsResponse_Mission{
		{
			Id:                "asd",
			InternalName:      "asd",
			ScenarioType:      0,
			Status:            0,
			OptInStrategyType: 0,
			ExposedToCount:    0,
			ConditionType:     0,
			OptedInCount:      0,
			TimesCompleted:    0,
			ActiveFrom:        nil,
			ActiveTo:          nil,
			CreatedAt:         nil,
		},
	}}, nil
}

func main() {
	flag.Parse()
	lis, err := net.Listen("tcp", fmt.Sprintf(":%d", *port))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterMissionServiceServer(s, &server{})
	log.Printf("server listening at %v", lis.Addr())
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
