package email

import (
	"context"
	proto "github.com/popstr/micro-world/email-service/proto"
)

type emailClient struct {
	client proto.EmailService
}

type UserEmail struct {
	ID    int
	Email string
}

type Client interface {
	GetEmails(ctx context.Context, userIDs []int64) ([]UserEmail, error)
}

func NewEmailClient(c proto.EmailService) *emailClient {
	return &emailClient{
		client: c,
	}
}

func (c *emailClient) GetEmails(ctx context.Context, userIDs []int64) ([]UserEmail, error) {
	r, err := c.client.GetEmail(ctx, &proto.Request{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}
	return toModel(r.Users), nil
}

func toModel(names []*proto.User) (result []UserEmail) {
	for _, n := range names {
		result = append(result, UserEmail{
			ID:    int(n.Id),
			Email: n.Email,
		})
	}
	return
}
