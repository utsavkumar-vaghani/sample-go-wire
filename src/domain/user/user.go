package user

import "context"

type User struct {
	Name string `json:"name"  bson:"name"`
}

type UserRepository interface {
	Add(context.Context, User) error
	Find(ctx context.Context, name string) (*User, error)
}

type UserService interface {
	Add(context.Context, User) error
}
