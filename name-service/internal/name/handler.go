package name

import (
	"context"
	proto "github.com/popstr/micro-world/name-service/proto"
)

type Handler struct {
	nameService Service
}

func NewHandler(nameService Service) *Handler {
	return &Handler{
		nameService: nameService,
	}
}

func (g *Handler) GetNames(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	for _, userID := range req.UserIds {
		firstName, lastName, err := g.nameService.GetNames(ctx, userID)
		if err != nil {
			return err
		}
		rsp.Users = append(rsp.Users, &proto.User{
			Id:        userID,
			FirstName: firstName,
			LastName:  lastName,
		})
	}
	return nil
}
