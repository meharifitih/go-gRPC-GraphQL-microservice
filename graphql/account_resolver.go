package main

import (
	"context"
	"log"
	"time"
)

type accountResolver struct {
	server *Server
}

func (a *accountResolver) Orders(ctx context.Context, obj *Account) ([]*Order, error) {
	ctx, cancel := context.WithTimeout(ctx, 3*time.Second)
	defer cancel()

	ordersList, err := a.server.orderClient.GetOrdersForAccount(ctx, obj.ID)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var orders []*Order

	for _, o := range ordersList {
		var product []*OrderProduct
		for _, p := range o.Products {
			product = append(product, &OrderProduct{
				ID:          p.ID,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    int(p.Quantity),
			})
		}

		orders = append(orders, &Order{
			ID:         o.ID,
			CreatedAt:  o.CreatedAt,
			TotalPrice: o.TotalPrice,
			Products:   product,
		})
	}

	return orders, nil
}
