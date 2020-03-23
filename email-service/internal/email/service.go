package email

import (
	"context"
	"encoding/csv"
	"errors"
	"io"
	"log"
	"os"
	"strconv"
)

type Service interface {
	GetEmail(ctx context.Context, userID int64) (string, error)
	ReadUsersFile(usersFileName string) error
}

type emailService struct {
	users map[int]userRecord
}

type userRecord struct {
	ID    int
	Email string
}

func New() Service {
	return &emailService{
		users: make(map[int]userRecord),
	}
}

func (s *emailService) ReadUsersFile(usersFileName string) error {
	usersFile, err := os.Open(usersFileName)
	if err != nil {
		return err
	}

	r := csv.NewReader(usersFile)

	for {
		record, err := r.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		userID, err := strconv.Atoi(record[0])
		s.users[userID] = userRecord{
			ID:        userID,
			Email: record[1],
		}
	}
	return nil
}

func (s *emailService) GetEmail(_ context.Context, userID int64) (string, error) {
	record, found := s.users[int(userID)]
	if !found {
		return "", errors.New("user not found")
	}
	return record.Email, nil
}
