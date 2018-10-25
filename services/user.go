package grpctest

import (
	"context"
	"encoding/json"
	"fmt"
	"grpc-test/messages/pb"
	models "grpc-test/models"
)

// UserService ...
type UserService models.UserService

var user = make(map[string]interface{})

func (s *UserService) Get(ctx context.Context, req *pb.GetRequest) (res *pb.GetResponse, err error) {
	details := user[req.UserID]
	serializedDetails, _ := json.Marshal(details)

	res = &pb.GetResponse{
		Details: serializedDetails,
	}

	return res, nil
}

func (s *UserService) Set(ctx context.Context, req *pb.User_SetRequest) (res *pb.User_SetResponse, err error) {
	details := make(map[string]interface{})

	err = json.Unmarshal(req.Details, &details)
	if err != nil {
		fmt.Println("error: ", err)
	}

	user[req.UserID] = details
	serializedDetails, _ := json.Marshal(details)

	res = &pb.User_SetResponse{
		Status:  true,
		Details: serializedDetails,
	}

	return res, nil
}

func (s *UserService) SendJson(ctx context.Context, req *pb.SendJsonRequest) (res *pb.SendJsonResponse, err error) {
	fmt.Println("\n===========\nrequest received by sendJson is: ", req)
	res = &pb.SendJsonResponse{}
	return res, nil
}

func (s *UserService) SendJsonString(ctx context.Context, req *pb.SendJsonStringRequest) (res *pb.SendJsonStringResponse, err error) {
	fmt.Println("\n===========\nrequest received by sendJsonString is: ", req)
	res = &pb.SendJsonStringResponse{}
	return res, nil
}
