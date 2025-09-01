package order

import (
	"context"
	"log"
	"time"

	pb "github.com/meharifitih/go-grpc-graphql-microservice/order/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type client struct {
	conn    *grpc.ClientConn
	service pb.OrderServiceClient
}

func NewClient(url string) (*client, error) {
	conn, err := grpc.NewClient(url, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, err
	}

	c := pb.NewOrderServiceClient(conn)

	return &client{
		conn:    conn,
		service: c,
	}, nil
}

func (c *client) Close() {
	c.conn.Close()
}

func (c *client) PostOrder(ctx context.Context, accountId string, products []OrderedProduct) (*Order, error) {
	protoProducts := []*pb.PostOrderRequest_OrderProduct{}

	for _, p := range products {
		protoProducts = append(protoProducts, &pb.PostOrderRequest_OrderProduct{
			ProductId: p.ID,
			Quantity:  uint32(p.Quantity),
		})
	}

	r, err := c.service.PostOrder(ctx, &pb.PostOrderRequest{
		AccountId: accountId,
		Products:  protoProducts,
	})
	if err != nil {
		return nil, err
	}

	newOrder := r.Order
	newOrderCreatedAt := time.Time{}
	newOrderCreatedAt.UnmarshalBinary(newOrder.CreatedAt)

	return &Order{
		ID:         newOrder.Id,
		CreatedAt:  newOrderCreatedAt,
		TotalPrice: newOrder.TotalPrice,
		AccountID:  newOrder.AccountId,
		Products:   products,
	}, nil
}

func (c *client) GetOrdersForAccount(ctx context.Context, accountID string) ([]Order, error) {
	res, err := c.service.GetOrdersForAccount(ctx, &pb.GetOrdersForAccountRequest{AccountId: accountID})
	if err != nil {
		log.Println(err)
		return nil, err
	}

	orders := []Order{}

	for _, orderProto := range res.Orders {
		newOrder := Order{
			ID:         orderProto.Id,
			TotalPrice: orderProto.TotalPrice,
			AccountID:  orderProto.AccountId,
		}
		newOrder.CreatedAt = time.Time{}
		newOrder.CreatedAt.UnmarshalBinary(orderProto.CreatedAt)

		products := []OrderedProduct{}
		for _, p := range orderProto.Products {
			products = append(products, OrderedProduct{
				ID:          p.Id,
				Name:        p.Name,
				Description: p.Description,
				Price:       p.Price,
				Quantity:    int64(p.Quantity),
			})
		}
		newOrder.Products = products

		orders = append(orders, newOrder)

	}

	return orders, nil
}
