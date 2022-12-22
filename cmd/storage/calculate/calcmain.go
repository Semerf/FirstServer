package main

import (
	"fmt"
	"log"
	"net"

	calc "github.com/semerf/FirstServer/internal/calcpb"
	calcpb "github.com/semerf/FirstServer/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":9997")
	if err != nil {
		log.Fatal(err)
	}

	grpcs := grpc.NewServer()
	srv := calc.GRPCServer{}
	calcpb.RegisterCalculatorServer(grpcs, &srv)

	if err := grpcs.Serve(l); err != nil {
		log.Fatal(err)
	}
	fmt.Println("Hello, world!")
}
