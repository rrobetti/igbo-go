package client

import (
	"context"
	"errors"
	"fmt"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	api "igbo-go/grpc"
)

type IgboDbClient struct {
	client api.IgboDBClient
}

func NewIgboDbClient(URL string) IgboDbClient {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := api.NewIgboDBClient(conn)
	return IgboDbClient{client: client}
}

func (c *IgboDbClient) Create(ctx context.Context, objects *api.Objects) (*api.OperationResults, error) {
	resp, err := c.client.Create(ctx, objects)
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}

func (c *IgboDbClient) Update(ctx context.Context, objects *api.Objects) (*api.OperationResults, error) {
	resp, err := c.client.Update(ctx, objects)
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}

func (c *IgboDbClient) Delete(ctx context.Context, ids *api.Ids) (*api.OperationResults, error) {
	resp, err := c.client.Delete(ctx, ids)
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *IgboDbClient) Retrieve(ctx context.Context, ids *api.Ids) (*api.Objects, error) {
	resp, err := c.client.Retrieve(ctx, ids)
	if err != nil {
		st, _ := status.FromError(err)
		if st.Code() == codes.NotFound {
			return nil, ErrIDNotFound
		} else {
			return nil, fmt.Errorf("Unexpected Insert failure: %w", err)
		}
	}
	return resp, nil
}
