package main

import (
	"context"
	"fmt"
	"log"

	"github.com/psenna/grpc-go-course/greet/greetpb"
	"google.golang.org/grpc"
)

func main() {
	fmt.Println("client")
	conn, err := grpc.Dial("localhost:5000", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("%v", err)
	}

	defer conn.Close()

	c := greetpb.NewGreatServiceClient(conn)

	req := &greetpb.GreetRequest{
		Greeting: &greetpb.Greeting{
			Fn: "Pedro",
			Ln: "senna",
		},
	}

	c.Grre(context.Background(), in*greetpb.GreetRequest)
	fmt.Printf("Created client: %f", c)
}
