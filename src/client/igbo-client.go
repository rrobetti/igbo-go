package client

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/status"
	api "igbo-go/grpc"
)

type IgboDbClient struct {
	client api.IgboDBClient
	stream api.IgboDB_OperationsStreamClient
}

func NewIgboDbClient(URL string) IgboDbClient {
	conn, err := grpc.Dial(URL, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	client := api.NewIgboDBClient(conn)
	stream, err := client.OperationsStream(context.Background())
	done := make(chan bool)

	// if stream is finished it closes done channel
	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			log.Printf("Response received for req id %v Req type %v, Result: %v", resp.RequestId.Id, resp.RequestId.Type, resp.Payload)
		}
	}()

	return IgboDbClient{client, stream}
}

func (c *IgboDbClient) OperationsStream(request api.OperationRequest) error {
	err := c.stream.Send(&request)
	if err != nil {
		return fmt.Errorf("Sending create request failure: %w", err)
	}
	return nil
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

func (c *IgboDbClient) Delete(ctx context.Context, keys *api.ObjectKeys) (*api.OperationResults, error) {
	resp, err := c.client.Delete(ctx, keys)
	if err != nil {
		return nil, fmt.Errorf("Insert failure: %w", err)
	}
	return resp, nil
}

var ErrIDNotFound = errors.New("Id not found")

func (c *IgboDbClient) Retrieve(ctx context.Context, keys *api.ObjectKeys) (*api.Objects, error) {
	resp, err := c.client.Retrieve(ctx, keys)
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
