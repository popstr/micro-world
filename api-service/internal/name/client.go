package name

import (
	"context"
	proto "github.com/popstr/micro-world/name-service/proto"
)

type nameClient struct {
	client proto.NamesService
}

type Client interface {
	GetNames(ctx context.Context, userIDs []int64) ([]Name, error)
}

func NewNamesClient(c proto.NamesService) *nameClient {
	return &nameClient{
		client: c,
	}
}

func (c *nameClient) GetNames(ctx context.Context, userIDs []int64) ([]Name, error) {
	r, err := c.client.GetNames(ctx, &proto.Request{
		UserIds: userIDs,
	})
	if err != nil {
		return nil, err
	}
	return toModel(r.Users), nil
}

func toModel(names []*proto.User) (result []Name) {
	for _, n := range names {
		result = append(result, Name{
			ID:        int(n.Id),
			FirstName: n.FirstName,
			LastName:  n.LastName,
		})
	}
	return
}
