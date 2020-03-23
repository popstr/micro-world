package email

import (
	"context"
	proto "github.com/popstr/micro-world/email-service/proto"
)

type Handler struct {
	emailService Service
}

func NewHandler(emailService Service) *Handler {
	return &Handler{
		emailService: emailService,
	}
}

func (g *Handler) GetEmail(ctx context.Context, req *proto.Request, rsp *proto.Response) error {
	for _, userID := range req.UserIds {
		email, err := g.emailService.GetEmail(ctx, userID)
		if err != nil {
			return err
		}
		rsp.Users = append(rsp.Users, &proto.User{
			Id: userID,
			Email: email,
		})
	}
	return nil
}
