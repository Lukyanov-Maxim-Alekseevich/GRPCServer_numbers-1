package adder

import (
	"context"
	"fmt"
	"test/pkg/api/pkg/api"
	"test/repository"
	
)

type GRPCServer struct {
	api.UnimplementedNumberServer
}

func (s *GRPCServer)YesNumber(ctx context.Context, req *api.YesRequest) (*api.YesResponse,error){

	answer, count:=repository.FindInSQLStorage(req.FindNumber)
	if answer=="yes"{
		fmt.Printf("вы выиграли, найденных сходств %d \n",count)
	}
	if answer=="no"{
		fmt.Printf("вы проиграли, найденных сходств %d \n",count)
	}
	req.FindNumber=count
	return &api.YesResponse{AnswerFindNumber: req.GetFindNumber()},nil
}

func(s *GRPCServer)NoNumber(ctx context.Context, req *api.NoRequest)(*api.NoResponse,error){

	answer, count:=repository.FindInSQLStorage(req.NoFindNumber)
	if answer=="no"{
		fmt.Printf("вы выиграли, найденных сходств %d \n",count)
	}
	if answer=="yes"{
		fmt.Printf("вы проиграли, найденных сходств %d \n",count)
	}
	req.NoFindNumber=count
	return &api.NoResponse{AnswerNoFindNumber: req.GetNoFindNumber()},nil
}