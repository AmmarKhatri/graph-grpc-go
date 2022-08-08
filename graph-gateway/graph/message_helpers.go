package graph

import (
	"context"
	"graph-gateway/grpc/message"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

//all the helper functions for the grpc client of message service are here

func SaveMessage(msg *message.Message) (*message.MessageResponse, error) {
	input := message.MessageRequest{
		Msg: msg,
	}
	conn, err := grpc.Dial("message-service:50001", grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		return nil, err
	}
	defer conn.Close()
	c := message.NewMessageServiceClient(conn)
	res, err := c.SaveMessage(context.Background(), &input)
	if err != nil {
		return nil, err
	}
	return res, nil
}
