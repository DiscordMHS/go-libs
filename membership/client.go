package membership

import (
	"context"
	"fmt"

	membershipspb "github.com/DiscordMHS/protocols/gen/go/memberships/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// Client представляет клиент для работы с membership сервисом.
type Client struct {
	conn   *grpc.ClientConn
	client membershipspb.MembershipServiceClient
}

// NewClient создает новый клиент для membership сервиса.
// address - адрес gRPC сервера (например, "localhost:50051").
func NewClient(address string) (*Client, error) {
	conn, err := grpc.NewClient(address, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		return nil, fmt.Errorf("grpc.NewClient: %w", err)
	}

	return &Client{
		conn:   conn,
		client: membershipspb.NewMembershipServiceClient(conn),
	}, nil
}

// NewClientWithConn создает новый клиент с использованием существующего gRPC соединения.
func NewClientWithConn(conn *grpc.ClientConn) *Client {
	return &Client{
		conn:   conn,
		client: membershipspb.NewMembershipServiceClient(conn),
	}
}

// Close закрывает соединение с сервисом.
func (c *Client) Close() error {
	if c.conn != nil {
		return c.conn.Close()
	}
	return nil
}

// CheckPermission проверяет, имеет ли пользователь указанное разрешение для ресурса.
// userID - идентификатор пользователя.
// resource - путь к ресурсу (например, "guild/{guild_id}/channel/{channel_id}").
// permission - название разрешения (например, "send_messages", "view_channel").
// Возвращает true, если разрешение есть, и false в противном случае.
func (c *Client) CheckPermission(
	ctx context.Context,
	userID uint64,
	resource string,
	permission string,
) (bool, error) {
	req := &membershipspb.CheckPermissionRequest{
		UserId:     userID,
		Resource:   resource,
		Permission: permission,
	}

	resp, err := c.client.CheckPermission(ctx, req)
	if err != nil {
		return false, fmt.Errorf("c.client.CheckPermission: %w", err)
	}

	return resp.Allowed, nil
}
