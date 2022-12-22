package calc

import (
	"context"
	"fmt"

	"github.com/semerf/FirstServer/internal/calculate"
	"github.com/semerf/FirstServer/internal/database"
	calcpb "github.com/semerf/FirstServer/proto"
)

type GRPCServer struct {
	calcpb.CalculatorServer
}

func (s *GRPCServer) Calc(ctx context.Context, req *calcpb.RequestQuery) (*calcpb.Response, error) {
	tasks := []database.Task{}
	var task database.Task
	for _, v := range req.Requests {
		task.Task_id = int(v.TaskId)
		task.Task_name = v.TaskName
		task.Duration = int(v.Duration)
		task.Resource = int(v.Resource)
		task.Prev_work = v.PrevWork
		task.Order_id = int(v.OrderId)
		tasks = append(tasks, task)
	}
	result := calculate.Calculator(tasks)
	fmt.Println(result)
	return &calcpb.Response{Message: int32(result)}, nil
}
