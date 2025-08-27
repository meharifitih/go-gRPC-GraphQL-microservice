package main

import "context"

type queryResolver struct {
	server *Server
}

func (q *queryResolver) Accounts(ctx context.Context, pagination *PaginationInput,
	id *string) ([]*Account, error)
func (q *queryResolver) Products(ctx context.Context, pagination *PaginationInput,
	query *string, id *string) ([]*Product, error)
