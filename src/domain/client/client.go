package client

import "context"

type Client struct {
	Name string `json:"name"  bson:"name"`
}

type ClientRepository interface {
	Add(context.Context, Client) error
	Find(ctx context.Context, name string) (*Client, error)
}

type ClientService interface {
	Add(context.Context, Client) error
}
