package main

import (
	"context"
	"fmt"
	"log"
	"message-service/data"
	"message-service/message"
	"net"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

type MessageServer struct {
	message.MessageServiceServer
	Models data.Models
}

var mongoURL = "mongodb://mongod:27017"
var collection *mongo.Collection
var port = "50001"

func main() {
	// if we crash the go code, we get the file and line number
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	//initializing our db
	clientOptions := options.Client().ApplyURI(mongoURL)
	clientOptions.SetAuth(options.Credential{
		Username: "admin",
		Password: "secretpass",
	})

	// connect
	c, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	client := c
	collection = client.Database("message").Collection("message")
	// initialize the server
	lis, err := net.Listen("tcp", fmt.Sprintf(":%s", port))
	if err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	s := grpc.NewServer()

	message.RegisterMessageServiceServer(s, &MessageServer{
		Models: data.Models{
			Msg: data.Msg{},
		},
	})

	log.Printf("gRPC Server started on port %s", port)

	// go func() {
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to listen for gRPC: %v", err)
	}
	// }()

	// //Wait for control C to exit
	// ch := make(chan os.Signal, 1)
	// signal.Notify(ch, os.Interrupt)

	// // block until a signal is received
	// <-ch

	// fmt.Println("Stopping the server")
	// s.Stop()
	// fmt.Println("Stopping the listener")
	// lis.Close()
	// fmt.Println("Closing MongoDB connection")
	// client.Disconnect(context.TODO())
	// fmt.Println("End of Program")
}

func (*MessageServer) SaveMessage(ctx context.Context, req *message.MessageRequest) (*message.MessageResponse, error) {
	msg := req.GetMsg()
	data := data.Msg{
		Data:      msg.GetData(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
	res, err := collection.InsertOne(context.TODO(), data)
	if err != nil {
		log.Println("Error inserting into msg:", err)
		return nil, err
	}
	oid, ok := res.InsertedID.(primitive.ObjectID)
	if !ok {
		log.Println("Error casting to ObjectID")
		return nil, err
	}
	output := message.MessageResponse{
		Msg: &message.Message{
			Id:        oid.Hex(),
			Data:      msg.GetData(),
			CreatedAt: data.CreatedAt.Local().String(),
			UpdatedAt: data.CreatedAt.Local().String(),
		},
	}
	log.Println("Created At:", data.CreatedAt.GoString())
	log.Println("Updated At:", data.UpdatedAt.GoString())
	return &output, nil
}
