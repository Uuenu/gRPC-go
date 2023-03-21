package service

import (
	"context"
	"log"

	"github.com/google/uuid"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"main.go/pb"
)

// LaptopServer is the server that provide laptop service
type LaptopServer struct {
}

func NewLaptopService() *LaptopServer {
	return &LaptopServer{}
}

// CreateLaptop is a unary RPC to create a new laptop
func (server *LaptopServer) CreateLaptop(
	ctx context.Context,
	in *pb.CreateLaptopRequest,
	opts ...grpc.CallOption,
) (*pb.CreateLaptopResponse, error) {
	laptop := in.GetLaptop()
	log.Printf("recive a create-laptop request with id: %s", laptop.Id)

	if len(laptop.Id) > 0 {
		// check if its a valid UUID
		_, err := uuid.Parse(laptop.Id)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "laptop ID is not a valid UUID: %w", err)
		}
	} else {
		id, err := uuid.NewRandom()
		if err != nil {
			return nil, status.Errorf(codes.Internal, "cannot generate a new laptop ID: %w", err)
		}
		laptop.Id = id.String()
	}

	// save laptop to db
	return nil, nil
}
