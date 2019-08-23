package main

import (
	"context"
	"fmt"
	"log"
	"net"

	"github.com/psenna/grpc-go-course/greet/greetpb"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Grre(ctx context.Context, req *greetpb.GreetRequest) (*greetpb.GreetResponse, error) {
	fmt.Printf("foi")
	firstName := req.GetGreeting().GetFn()

	result := "Hello " + firstName

	res := &greetpb.GreetResponse{
		Result: result,
	}

	return res, nil
}

func main() {
	fmt.Println("Hello word")

	lis, err := net.Listen("tcp", ":5000")

	if err != nil {
		log.Fatalf("Failed to listem: %v", err)
	}

	s := grpc.NewServer()

	greetpb.RegisterGreatServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}

}
