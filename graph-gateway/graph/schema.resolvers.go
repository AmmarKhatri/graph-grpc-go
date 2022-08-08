package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"graph-gateway/graph/generated"
	"graph-gateway/graph/model"
	"graph-gateway/grpc/message"
	"log"
)

// CreateMessage is the resolver for the createMessage field.
func (r *mutationResolver) CreateMessage(ctx context.Context, input model.NewMessage) (*model.Message, error) {
	msg := message.Message{
		Data: input.Data,
	}
	res, err := SaveMessage(&msg)
	if err != nil {
		return nil, err
	}
	output := &model.Message{
		ID:        res.Msg.Id,
		Data:      res.Msg.Data,
		CreatedAt: res.Msg.CreatedAt,
		UpdatedAt: res.Msg.UpdatedAt,
	}
	log.Println("Output:", output)
	return output, nil
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

type mutationResolver struct{ *Resolver }
