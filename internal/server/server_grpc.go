package server

import (
	"context"
	"log"

	calcpb "github.com/semerf/FirstServer/proto"
)

func callCalc(client calcpb.CalculatorClient, tasks *calcpb.RequestQuery) int32 {
	stream, err := client.Calc(context.Background(), tasks)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(stream.Message)
	return stream.Message
}
