package user

import (
	"github.com/labstack/echo"
	"github.com/postr/micro-world/api-service/internal/email"
	"github.com/postr/micro-world/api-service/internal/name"
	"net/http"
	"strconv"
)

type UserResponse struct {
	Users map[int]*UserData `json:"users"`
}

type UserData struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

func HandleGetUser(namesClient name.Client, emailClient email.Client) func(echo.Context) error {
	return func(c echo.Context) error {
		userIDStr := c.Param("id")
		userID, err := strconv.Atoi(userIDStr)
		users := []int64{int64(userID)}
		ctx := c.Request().Context()

		names, err := namesClient.GetNames(ctx, users)
		if err != nil {
			return err
		}

		emails, err := emailClient.GetEmails(ctx, users)
		if err != nil {
			return err
		}

		nameResponse := UserResponse{
			Users: make(map[int]*UserData),
		}
		for _, n := range names {
			nameResponse.Users[n.ID] = &UserData{
				FirstName: n.FirstName,
				LastName:  n.LastName,
			}
		}
		for _, e := range emails {
			data, found := nameResponse.Users[e.ID]
			if !found {
				nameResponse.Users[e.ID] = &UserData{
					Email: e.Email,
				}
			} else {
				data.Email = e.Email
			}
		}
		err = c.JSON(http.StatusOK, nameResponse)
		return err
	}
}
