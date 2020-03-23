package name

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
	GetNames(ctx context.Context, userID int64) (string, string, error)
	ReadUsersFile(usersFileName string) error
}

type nameService struct {
	users map[int]userRecord
}

type userRecord struct {
	ID        int
	FirstName string
	LastName  string
}

func New() Service {
	return &nameService{
		users: make(map[int]userRecord),
	}
}

func (s *nameService) ReadUsersFile(usersFileName string) error {
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
			FirstName: record[1],
			LastName:  record[2],
		}
	}
	return nil
}

func (s *nameService) GetNames(_ context.Context, userID int64) (string, string, error) {
	record, found := s.users[int(userID)]
	if !found {
		return "", "", errors.New("user not found")
	}
	return record.FirstName, record.LastName, nil
}
